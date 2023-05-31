import { FeatureCollection, circle, point, featureCollection } from "@turf/turf";

interface Organisation {
    id: string,
    name: string
}

interface IOrganisationRepository {
    GetMine(): Promise<Array<Organisation>>
}

const ORGANISATION_FIXTURES: Array<Organisation> = [
    {
        id: "acme",
        name: 'Acme'
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