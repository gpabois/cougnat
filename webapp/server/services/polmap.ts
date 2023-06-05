import {MonitoringRepositoryFactory, PollutionData} from '../..//repository/monitoring';

import { tiles } from "../../geo/utils";
import { PNGStream, createCanvas } from "canvas";

import vector from "../../vector";
import color from "../../color";

const POL_TILE_ZOOM = 18;
const TILE_DIM: [number, number] = [256, 256]

export interface CurrentPollutionArgs {organisation_id: string, x: number, y: number, zoom: number}

function apply_heatmap(w: number) : string {
    const lambda = Math.max(0.0, Math.min(w / 100.0, 1.0));
    const hue_c = color.walk({h: 0.0, s: 1.0, l: 0.5}, {h: 0.3, s: 0.0, l: 0.5}, 1.0 - lambda)
    const c = color.hue.toRGB(hue_c);
    const hC = color.rgb.toHex(c)
    return hC
}

async function create_pollution_tile({x, y, zoom, organisation_id}: CurrentPollutionArgs) : Promise<PNGStream> {
    // The service to fetch the pollution matrix from the server.
    const monitoring = MonitoringRepositoryFactory(null);
    
    // Get the tile box.
    const tile_box = tiles.toBoxVec({x, y, zoom});
    
    // Create a scale for dims.
    const scale = vector.div(TILE_DIM, [tile_box[2], tile_box[3]])
    
    // Polygon to get the raw pollution matrix (polmat)
    const tile_polygon = tiles.toPolygon({x, y, zoom});
    
    // Fetch the polmat
    const pol_matrix = await monitoring.getCurrentPollution(organisation_id, [], tile_polygon, zoom + 2);

    // Create the canvas to draw on it.
    const tile_canvas = createCanvas(...TILE_DIM);

    // Loop over all PPX
    for(const pol_tile of pol_matrix) {
        // Weight
        const weight = pol_tile.types.map((polData: PollutionData) => polData.count).reduce((acc, w) => acc + w, 0)
        
        // PPx = Pollution's Pixel (Smallest unit of pollution)
        const ppx_box = tiles.toBoxVec(pol_tile.coordinates)
        const [t_x, t_y] = vector.mul(vector.diff(ppx_box, tile_box), scale);
        const [t_w, t_h] = vector.mul([ppx_box[2], ppx_box[3]], scale);

        const img_box: [number, number, number, number] = [t_x, t_y, t_w, t_h];

        // Draw the PPx
        if(weight > 0) {
            tile_canvas.getContext('2d').fillStyle = apply_heatmap(weight);
            tile_canvas.getContext('2d').fillRect(...img_box);
        }
    }

    // Return stream
    return tile_canvas.createPNGStream()
}

export {
    create_pollution_tile
}