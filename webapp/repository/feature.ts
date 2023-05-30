import {booleanWithin, Feature, FeatureCollection, point} from '@turf/turf'

interface IFeatureRepository {
    GetWithin(area: Feature): Promise<FeatureCollection>
}

// Order is longitude/latitude
const FEATURE_FIXTURES: Feature[] = [
    point([2.4904595576244026, 48.77932087129807], {
        type: "facility", 
        name: "Acme", 
        thumbnail: "/factory.jpeg",
        rating: 4,
        description: "Une usine qui produit des nuisances olfactives.",
        address: "27 route De L'Isle Saint-Julien, 94380 Bonneuil-sur-Marne"
    })
];

class MockFeatureRepository {
    async GetWithin(polygon: Feature): Promise<FeatureCollection> {
        return Promise.resolve({
            type: "FeatureCollection",
            features: FEATURE_FIXTURES.filter((feature) => booleanWithin(feature, polygon))
        })
    }
}

function FeatureRepositoryFactory(container: any): IFeatureRepository {
    return new MockFeatureRepository()
}

export {
    IFeatureRepository,
    FeatureRepositoryFactory
}

