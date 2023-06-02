import { FeatureCollection, circle, point, featureCollection } from "@turf/turf";

interface Organisation {
    id: string,
    name: string,
    permissions?: string[]
}

interface IOrganisationRepository {
    GetMine(): Promise<Array<Organisation>>
}

const ORGANISATION_FIXTURES: Array<Organisation> = [
    {
        id: "acme",
        name: 'Acme',
        permissions: ["broadcast"]
    }, {
        id: "mairie-saint-maur",
        name: 'Mairie Saint-Maur-des-Fossées'
    }, {
        id: "DRIEAT-IF-UD94",
        name: 'DRIEAT-IF/UD94'
    }
];

class MockOrganisationRepository {
    async GetMine(): Promise<Array<Organisation>> {
        return Promise.resolve(ORGANISATION_FIXTURES)
    }
}

function OrganisationRepositoryFactory(container: any): IOrganisationRepository {
    return new MockOrganisationRepository()
}

export {
    Organisation,
    IOrganisationRepository,
    OrganisationRepositoryFactory
}