import { FeatureCollection, circle } from "@turf/turf";

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
        nature: 'smell',
        label: 'Odeur de bitume'
    }, {
        name: 'concasseur',
        nature: 'noise',
        label: 'Bruit de concasseur'
    }
];
