import { FeatureCollection, circle, point, featureCollection } from "@turf/turf";

interface Organisation {
    id: string,
    name: string,
    areas: FeatureCollection
}

interface IOrganisationRepository {
    GetMine(): Promise<Array<Organisation>>
}

const ORGANISATION_FIXTURES: Array<Organisation> = [
    {
        id: "1",
        name: 'Acme',
        areas: featureCollection([
            circle(
                point([2.4904595576244026, 48.77932087129807]),
                300
            )
        ])
    }
];
