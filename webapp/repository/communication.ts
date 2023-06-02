import SnowflakeId from 'snowflake-id'

const snowflake = new SnowflakeId({
    mid : 42,
    offset : (2019-1970)*31536000*1000
});

interface Broadcast {
    id: string,
    organisation_id: string,
    content: string,
    created_at: Date,
    closed_at: Date | null
}

interface ICommunicationRepository {
    GetBroadcastsByOrganisation(organisation_id: string): Promise<Array<Broadcast>>
}

const BROADCAST_FIXTURES: Array<Broadcast> = [
    {
        id:  snowflake.generate(),
        organisation_id: 'acme',
        content: "Une activité potentiellement émettrice de nuisances est en cours pendant deux heures.",
        created_at: new Date(),
        closed_at: null
    }
];

class MockCommunicationRepository {
    async GetBroadcastsByOrganisation(organisation_id: string): Promise<Array<Broadcast>> {
        return Promise.resolve(BROADCAST_FIXTURES.filter((b) => b.organisation_id == organisation_id))
    }
}

function CommunicationRepositoryFactory(container: any): ICommunicationRepository {
    return new MockCommunicationRepository()
}

export {
    Broadcast,
    ICommunicationRepository,
    CommunicationRepositoryFactory
}