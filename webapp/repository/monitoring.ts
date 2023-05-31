import { FeatureCollection, circle, point, featureCollection, Feature, GeometryCollection, intersect, Polygon, bbox, bboxPolygon, Geometry } from "@turf/turf";
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
                300
            )
        ])
    }
]

interface IMonitoringRepository {
    GetPerimeter(organisation_id: string): Promise<MonitoringPerimeter>
    GetCurrentPollution(organisation_id: string, bounds: Polygon, zoom: number): Promise<PollutionTileCollection>
}

class MockMonitoringRepository {
    GetPerimeter(organisation_id: string): Promise<MonitoringPerimeter> {
        const perm = MONITORING_PERIMETER_FIXTURES.find((perim) => perim.organisation_id == organisation_id)
        if (perm) {
            return Promise.resolve(perm)
        } else {
            return Promise.reject("not found")
        }
    }

    async GetCurrentPollution(organisation_id: string, bounds: Polygon, zoom: number): Promise<PollutionTileCollection> {
        const perimeter = await this.GetPerimeter(organisation_id);
        const box = bbox(intersect(bounds, perimeter.areas.features[0].geometry as Polygon))
        const tileBox = tiles.fromBox(box, zoom);
        const pollutionTiles: Array<Feature<Geometry, PollutionTileProperties>> = [];

        for(const tileIndex of tiles.iterBounds(tileBox)) {
            pollutionTiles.push(bboxPolygon<PollutionTileProperties>(tiles.toBox(tileIndex), {
                properties: {
                    weight: Math.random() * 300
                }
            }))
        }

        return featureCollection(pollutionTiles)
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



