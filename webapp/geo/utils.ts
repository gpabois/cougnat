import { BBox, Feature, Point, Polygon, bbox, bboxPolygon, featureCollection, point } from "@turf/turf";
import {LatLngBounds, polygon} from 'leaflet';

export function latLngBoundsToFeature(refBbox: LatLngBounds | Ref<LatLngBounds>): Feature<Polygon> {
    const bbox = unref(refBbox);
    const ne = bbox.getNorthEast();
    const sw = bbox.getSouthWest();
    return bboxPolygon([ne.lng, ne.lat, sw.lng, sw.lat])
}

export namespace tiles {
    /**
     * Zoom of 12 is enough to aggregate reports into tiles with enough precision.
     */

    export interface TileIndex {
        x: number,
        y: number,
        zoom: number,
    }

    export interface TileIndexBounds {
        ne: TileIndex,
        sw: TileIndex,
        zoom: number
    }

    export function nextNE(tileIndex: TileIndex): TileIndex {
        return {
            x: tileIndex.x - 1,
            y: tileIndex.y - 1,
            zoom: tileIndex.zoom
        }
    }

    export function nextSW(tileIndex: TileIndex): TileIndex {
        return {
            x: tileIndex.x + 1,
            y: tileIndex.y + 1,
            zoom: tileIndex.zoom
        }
    }

    export function * iterBounds(bounds: TileIndexBounds): Generator<TileIndex> {
        for(let x = bounds.ne.x; x <= bounds.sw.x; x++) {
            for(let y = bounds.ne.y; y >= bounds.sw.y; y--) {
                const tile = {
                    x, y, zoom: bounds.zoom
                }
                yield tile
            }
        }
    }

    /**
     * Return tile coordinates based on a point expressed in the WSG84 Coordinates System 
     * 
     * See : https://wiki.openstreetmap.org/wiki/Slippy_map_tilenames#Lon..2Flat._to_tile_numbers_2
     * @param point In WSG 84
     */
    export function fromPoint(point: Point, zoom: number): TileIndex {
        const n = Math.pow(2, zoom);
        const x = n * ((point.coordinates[0] + 180.0) / 360.0);
        const y = (Math.floor((1-Math.log(Math.tan(point.coordinates[1]*Math.PI/180) + 1/Math.cos(point.coordinates[1]*Math.PI/180))/Math.PI)/2 * n ));

        return {
            x: Math.round(x), y: Math.round(y), zoom
        }
    }

    // Returns a LatLng vector
    export function toVec(tile: TileIndex) : Vector {
        const n = Math.pow(2, tile.zoom);
        const long_deg = tile.x / n * 360.0 - 180.0;
        const lat_rad = Math.atan(Math.sinh(Math.PI * (1 - 2 * tile.y / n)));
        const lat_deg = lat_rad * 180.0 / Math.PI;
        return [lat_deg, long_deg]
    }

    export function toPoint(tile: TileIndex): Point {
        const v = toVec(tile)
        return point([v[1], v[0]]).geometry
    }

    export function fromBox(box: BBox, zoom: number): TileIndexBounds {
        const ne: Point = point([box[0], box[1]]).geometry;
        const sw: Point = point([box[2], box[3]]).geometry;
        return {
            ne: fromPoint(ne, zoom),
            sw: nextNE(fromPoint(sw, zoom)),
            zoom
        }
    }

    export function toBox(tile: TileIndex): BBox {
        const ne = point(toPoint(tile).coordinates);
        const sw = point(toPoint(nextSW(tile)).coordinates);
        return bbox(featureCollection([ne, sw]))    
    }
}
