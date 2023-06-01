import { FeatureCollection, circle, point, featureCollection, Feature, GeometryCollection, intersect, Polygon, bbox, bboxPolygon, Geometry, polygon, union } from "@turf/turf";
import { tiles } from "~/geo/utils";

interface PollutionTileProperties {
    weight: number
}


interface PollutionTileCollection extends FeatureCollection<Geometry, PollutionTileProperties> {

}

interface MonitoringPerimeter {
    organisation_id: string,
    areas: FeatureCollection
}

const MONITORING_PERIMETER_FIXTURES: Array<MonitoringPerimeter> = [
    {
        organisation_id: "acme",
        areas: featureCollection([
            circle(
                point([2.4904595576244026, 48.77932087129807]),
                5
            )
        ])
    }
]

interface IMonitoringRepository {
    GetPerimeter(organisation_id: string): Promise<MonitoringPerimeter>
    GetCurrentPollution(organisation_id: string, bounds: Polygon, zoom: number): Promise<FeatureCollection>
}

class MockMonitoringRepository implements IMonitoringRepository {
    GetPerimeter(organisation_id: string): Promise<MonitoringPerimeter> {
        const perm = MONITORING_PERIMETER_FIXTURES.find((perim) => perim.organisation_id == organisation_id)
        if (perm) {
            return Promise.resolve(perm)
        } else {
            return Promise.reject("not found")
        }
    }

    async GetCurrentPollution(organisation_id: string, bounds: Polygon, zoom: number): Promise<FeatureCollection> {
        zoom = Math.max(zoom, 16)
        const perimeter = await this.GetPerimeter(organisation_id);
        const perimeterPolygons = perimeter.areas.features.map((feature) => feature)
        const roi = intersect(perimeterPolygons[0], bounds)
                
        if(roi == null)
            return featureCollection([])

        const box = bbox(roi)


        const tileBox = tiles.fromBox(box, zoom);
        const pollution_tiles: Array<Feature<Geometry, PollutionTileProperties>> = [];
        
        for(const tileIndex of tiles.iterBounds(tileBox)) {
            const tileBBox = tiles.toBox(tileIndex)
            const tilePolygon = bboxPolygon<PollutionTileProperties>(tileBBox, {
                properties: {
                    weight: Math.random() * 300
                }
            })

            const croppedPolygon = intersect(perimeterPolygons[0], tilePolygon)
            if(!croppedPolygon) continue;
            pollution_tiles.push(croppedPolygon)
        }

        return featureCollection(pollution_tiles)
    }
}

function MonitoringRepositoryFactory(container: any): IMonitoringRepository {
    return new MockMonitoringRepository()
}

export {
    MonitoringPerimeter,
    PollutionTileCollection,
    IMonitoringRepository,
    MonitoringRepositoryFactory
}



