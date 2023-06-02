import { FeatureCollection, circle, point, featureCollection, Feature, GeometryCollection, intersect, Polygon, bbox, bboxPolygon, Geometry, polygon, union } from "@turf/turf";
import haversineDistance from "haversine-distance";
import { max } from "moment";
import { tiles } from "~/geo/utils";
import itertools from "~/itertools";
import vector from "~/vector";

interface PollutionTile {
    coordinates: tiles.TileIndex
    types: PollutionData[]
}
interface PollutionData {type: string, count: number}
interface PollutionMatrix extends Array<PollutionTile> {}

interface SectorMonitoring {
    id: string,
    label: string
}

interface OrganisationMonitoring {
    organisation_id: string,
    sectors: FeatureCollection<Polygon, SectorMonitoring>
}

const ORG_MONITORING_FIXTURES: Array<OrganisationMonitoring> = [
    {
        organisation_id: "acme",
        sectors: {
            "type": "FeatureCollection",
            "features": [
              {
                "type": "Feature",
                "properties": {
                    id: "zone_bitume",
                    label: "Zone de dispersion des odeurs de bitume"
                },
                "geometry": {
                  "coordinates": [
                    [
                      [
                        2.488579976014819,
                        48.77984317666676
                      ],
                      [
                        2.5200164952908324,
                        48.78739853245142
                      ],
                      [
                        2.4713156580518785,
                        48.7988567309111
                      ],
                      [
                        2.4740285637413706,
                        48.788862478466086
                      ],
                      [
                        2.488579976014819,
                        48.77984317666676
                      ]
                    ]
                  ],
                  "type": "Polygon"
                }
              }
            ]
          }
    }, {
        organisation_id: "mairie-saint-maur",
        sectors: featureCollection([{"type":"Feature","geometry":{"type":"Polygon","coordinates":[[[2.52228,48.80407],[2.52282,48.80217],[2.52348,48.7958],[2.52245,48.79228],[2.51991,48.78833],[2.51686,48.78647],[2.51358,48.78568],[2.51251,48.78492],[2.5102,48.78419],[2.50139,48.78472],[2.4984,48.78468],[2.49241,48.78411],[2.48717,48.78421],[2.48004,48.78569],[2.47486,48.78764],[2.47215,48.79088],[2.47188,48.79366],[2.47034,48.7992],[2.46835,48.80438],[2.46402,48.80736],[2.46063,48.80936],[2.46458,48.81179],[2.46546,48.81107],[2.46995,48.81254],[2.47277,48.81456],[2.47392,48.81592],[2.48019,48.81352],[2.48677,48.81253],[2.49324,48.81108],[2.49801,48.81044],[2.5055,48.81048],[2.51652,48.80893],[2.51916,48.80758],[2.52228,48.80407]]]},
        "properties":{"id":"commune","label":"Commune"}
        }])
    }, {
        organisation_id: "DRIEAT-IF-UD94",
        sectors: featureCollection([{
            "type":"Feature",
            "geometry":{
                "type":"Polygon",
                "coordinates":[[[2.3319,48.81701],[2.33371,48.81677],[2.34717,48.8161],[2.35239,48.81856],[2.35615,48.81598],[2.36293,48.81608],[2.37361,48.81934],[2.38076,48.8217],[2.38153,48.82242],[2.38901,48.82512],[2.39396,48.82743],[2.40247,48.82954],[2.40493,48.82844],[2.4077,48.82635],[2.41114,48.82489],[2.41652,48.82474],[2.41996,48.82408],[2.42624,48.82413],[2.42923,48.82362],[2.43491,48.81967],[2.43745,48.81911],[2.43746,48.81818],[2.44186,48.81795],[2.4498,48.81788],[2.45333,48.81716],[2.45723,48.81698],[2.459,48.81724],[2.4627,48.81906],[2.46286,48.82018],[2.46504,48.82409],[2.46618,48.82733],[2.46457,48.82766],[2.46496,48.83044],[2.46569,48.83181],[2.46896,48.83392],[2.46971,48.83658],[2.46522,48.84111],[2.46226,48.84254],[2.45789,48.84349],[2.44641,48.84493],[2.44653,48.84575],[2.44075,48.84596],[2.44074,48.84441],[2.43796,48.84467],[2.43718,48.84089],[2.43354,48.84105],[2.42453,48.84189],[2.42184,48.84444],[2.41987,48.84345],[2.41953,48.84145],[2.42093,48.83933],[2.42214,48.83579],[2.41734,48.83419],[2.4135,48.83372],[2.41224,48.83454],[2.41357,48.83826],[2.41574,48.84531],[2.41634,48.84924],[2.41931,48.84933],[2.42918,48.84874],[2.42953,48.85098],[2.43443,48.85085],[2.43518,48.85328],[2.44214,48.85291],[2.44438,48.85178],[2.4471,48.85116],[2.44874,48.85344],[2.45386,48.85546],[2.45604,48.85582],[2.4673,48.85666],[2.46676,48.86003],[2.46825,48.8607],[2.47357,48.86025],[2.47596,48.86042],[2.48078,48.86015],[2.48153,48.86141],[2.48875,48.86055],[2.49101,48.86081],[2.49162,48.85968],[2.49395,48.86058],[2.49612,48.86076],[2.49647,48.85919],[2.49496,48.85883],[2.49652,48.85609],[2.49897,48.85675],[2.49972,48.85568],[2.50394,48.85733],[2.50544,48.85719],[2.5068,48.85508],[2.50938,48.8531],[2.51197,48.85242],[2.51367,48.85017],[2.51565,48.85138],[2.51898,48.84818],[2.52484,48.84907],[2.5287,48.84376],[2.53443,48.84528],[2.53563,48.8429],[2.53745,48.84117],[2.54026,48.83934],[2.53665,48.83851],[2.5382,48.8368],[2.54418,48.83499],[2.5474,48.83611],[2.55284,48.83252],[2.55501,48.83134],[2.55665,48.83134],[2.55758,48.83009],[2.56037,48.82882],[2.56132,48.82648],[2.56812,48.82426],[2.57221,48.82366],[2.57098,48.82268],[2.5685,48.81828],[2.57006,48.81504],[2.57454,48.81369],[2.57413,48.81291],[2.57801,48.81211],[2.58243,48.81024],[2.59024,48.80901],[2.59228,48.80744],[2.59648,48.80613],[2.59422,48.80168],[2.59211,48.7996],[2.59294,48.79761],[2.59605,48.79695],[2.59977,48.79498],[2.59928,48.79335],[2.59423,48.78832],[2.59495,48.78504],[2.59196,48.78536],[2.59141,48.78361],[2.58848,48.78299],[2.58895,48.78074],[2.58845,48.77827],[2.58561,48.77862],[2.58552,48.77692],[2.5868,48.77355],[2.58931,48.77152],[2.59067,48.77183],[2.59278,48.77119],[2.59495,48.77214],[2.59796,48.77244],[2.59802,48.77321],[2.60136,48.77397],[2.607,48.7744],[2.60787,48.77283],[2.61113,48.77035],[2.61308,48.76579],[2.61484,48.76395],[2.61557,48.76237],[2.61482,48.76112],[2.61133,48.7612],[2.6116,48.76042],[2.60274,48.75897],[2.59754,48.76057],[2.59937,48.75549],[2.60234,48.75373],[2.60035,48.75095],[2.5975,48.74916],[2.59619,48.74732],[2.59177,48.7476],[2.59047,48.74657],[2.59012,48.74512],[2.58753,48.744],[2.58541,48.74171],[2.59036,48.73995],[2.59179,48.73842],[2.59151,48.73743],[2.59531,48.73652],[2.59415,48.73567],[2.59468,48.73181],[2.59032,48.73114],[2.58915,48.72917],[2.58786,48.72913],[2.58494,48.72711],[2.5809,48.72194],[2.57921,48.72275],[2.57748,48.71939],[2.57666,48.71574],[2.57528,48.71283],[2.5721,48.71346],[2.57107,48.71131],[2.56929,48.7109],[2.56814,48.70897],[2.5688,48.70722],[2.57056,48.70502],[2.57,48.70348],[2.57313,48.7019],[2.57522,48.70036],[2.58186,48.69744],[2.5735,48.69583],[2.57064,48.69494],[2.57166,48.69345],[2.57166,48.69201],[2.56616,48.69181],[2.56561,48.69113],[2.56139,48.69078],[2.56111,48.68966],[2.55387,48.68832],[2.55036,48.68854],[2.55079,48.69053],[2.54781,48.69269],[2.54565,48.69319],[2.54654,48.69517],[2.54422,48.69472],[2.5428,48.69688],[2.54431,48.69845],[2.5416,48.70012],[2.53622,48.69716],[2.53638,48.69987],[2.5359,48.70081],[2.53108,48.69979],[2.52858,48.70388],[2.52947,48.70792],[2.52642,48.70806],[2.52643,48.70927],[2.52255,48.70982],[2.52136,48.71229],[2.51915,48.71469],[2.51731,48.71591],[2.51755,48.71712],[2.51575,48.72894],[2.51026,48.73457],[2.50754,48.7331],[2.50755,48.73481],[2.5062,48.73571],[2.50354,48.73521],[2.48551,48.72939],[2.47877,48.72764],[2.47246,48.72756],[2.46793,48.72658],[2.46739,48.72871],[2.46597,48.72951],[2.46487,48.7272],[2.46281,48.72632],[2.46055,48.72618],[2.46001,48.72384],[2.45627,48.72398],[2.45423,48.72468],[2.45333,48.72253],[2.45416,48.72045],[2.44843,48.7218],[2.44752,48.71949],[2.44793,48.71795],[2.4506,48.71579],[2.4488,48.71483],[2.44624,48.71692],[2.44337,48.72079],[2.44237,48.72155],[2.44354,48.72446],[2.44116,48.7241],[2.44013,48.72543],[2.43868,48.72485],[2.43377,48.72382],[2.42554,48.7227],[2.42072,48.72105],[2.41426,48.71782],[2.41327,48.71869],[2.41582,48.72086],[2.41238,48.72193],[2.41013,48.72574],[2.40224,48.72655],[2.40155,48.72462],[2.39966,48.72466],[2.39747,48.7212],[2.39171,48.72263],[2.39,48.7208],[2.38675,48.72098],[2.38505,48.71914],[2.37824,48.7208],[2.37655,48.71941],[2.37071,48.72018],[2.37,48.73708],[2.36935,48.74606],[2.36291,48.74609],[2.36301,48.74411],[2.35934,48.7424],[2.35644,48.73816],[2.35395,48.73866],[2.34775,48.74148],[2.34549,48.74168],[2.34461,48.74064],[2.34169,48.74046],[2.33715,48.74226],[2.33771,48.74324],[2.33375,48.7468],[2.33096,48.74792],[2.32688,48.75116],[2.32479,48.75055],[2.32266,48.74812],[2.32072,48.74876],[2.31414,48.75135],[2.31235,48.75135],[2.31012,48.75219],[2.31164,48.75302],[2.31146,48.75433],[2.3087,48.75542],[2.31336,48.76015],[2.31311,48.76158],[2.31437,48.76206],[2.31567,48.7668],[2.32044,48.77135],[2.32374,48.77658],[2.32572,48.77707],[2.32662,48.77873],[2.32458,48.77907],[2.32581,48.78191],[2.32445,48.78225],[2.32512,48.78631],[2.3187,48.788],[2.32596,48.80666],[2.324,48.80537],[2.32039,48.80615],[2.31834,48.80603],[2.31786,48.80841],[2.31905,48.80995],[2.32551,48.80953],[2.3319,48.81701]]]
            },
            "properties":{"id":"94","label":"Département"}
        }])
    }
]

interface IMonitoringRepository {
    /**
     * Returns the monitoring properties at the organisation level, it includes all the sectors
     * @param organisation_id
     */
    getOrganisationMonitoring(organisation_id: string): Promise<OrganisationMonitoring>
    /**
     * Returns the current pollution for all the sectors based as parameters
     * @param organisation_id
     * @param sectors 
     * @param bounds 
     * @param zoom 
     */
    getCurrentPollution(organisation_id: string, sector_ids: string[], bounds: Polygon, zoom: number): Promise<PollutionMatrix>

    /**
     * Returns the pollution reported within a timeframe
     * @param organisation_id 
     * @param timeframe 
     * @param sector_ids 
     * @param bounds 
     * @param zoom 
     */
    getPollution(organisation_id: string, timeframe: [Date, Date], sector_ids: string[], bounds: Polygon, zoom: number) : Promise<PollutionMatrix>
}

class MockMonitoringRepository implements IMonitoringRepository 
{
    async generateRandomPollution(organisation_id: string, sectors: string[], bounds: Polygon, zoom: number) : Promise<PollutionMatrix> {
        zoom = Math.min(18, zoom)
        const perimeter = await this.getOrganisationMonitoring(organisation_id);
        const perimeterPolygons = perimeter.sectors.features.map((feature) => feature)
        const roi = intersect(perimeterPolygons[0], bounds)
                
        if(roi == null)
            return []

        const box = bbox(roi)
        const tileBox = tiles.fromBox(box, zoom);
        const pollution: PollutionMatrix = [];

        // Simulate wind SW/NE
        const wind_p1 =  [48.78024732197537, 2.4891518487230773];
        const wind_p2 = [48.78841736865386, 2.498748485396096]
        const wind_right = vector.diff(wind_p2, wind_p1)
        const wind_n = vector.normalise(wind_right) // LatLng
        const wind_speed = 10 // m/s
        const s = (Date.now() / 1000) % 10 // Create a pattern of 10 s 
        const max_scatter = wind_speed * s;

        for(const tileIndex of tiles.iterBounds(tileBox)) 
        {
            const tileCoords    = tiles.toVec(tileIndex)
            const tileBBox      = tiles.toBox(tileIndex)
            const tilePolygon   = bboxPolygon(tileBBox)
            const wind_tile_v   = vector.diff(tileCoords, wind_p1)

            const proj_wind_tile_v = itertools.toArray(
                vector.pipeline.add(
                    wind_p1,
                    vector.pipeline.scalar_mult(wind_n, vector.scalar_product(wind_tile_v, wind_n))
                )
            )
            const d = haversineDistance({latitude: proj_wind_tile_v[0], longitude: proj_wind_tile_v[1]}, {latitude: tileCoords[0], longitude: tileCoords[1]})
            const d2 = haversineDistance({
                latitude: wind_p1[0],
                longitude: wind_p1[1]
            }, {latitude: tileCoords[0], longitude: tileCoords[1]})

            const angle = Math.acos(vector.scalar_product(
                wind_tile_v, 
                wind_n
            ) / (vector.norm(wind_tile_v) * vector.norm(wind_n))) * 180.0 / Math.PI
            
            
            const attenuation = Math.exp(d2 / 200.0) / 3000.0
            const scattering = 1.0 / attenuation * Math.exp(-Math.abs(angle) / 10.0);
            const count = Math.random() * 100 * scattering;

            pollution.push({
                coordinates: tileIndex,
                types: [{
                    type: 'bitume',
                    count
                }]
            })
        }

        return pollution
    }

    getOrganisationMonitoring(organisation_id: string): Promise<OrganisationMonitoring> {
        const perm = ORG_MONITORING_FIXTURES.find((perim) => perim.organisation_id == organisation_id)
        if (perm) {
            return Promise.resolve(perm)
        } else {
            return Promise.reject("not found")
        }
    }

    getCurrentPollution(organisation_id: string, sectors: string[], bounds: Polygon, zoom: number): Promise<PollutionMatrix> {
        return this.generateRandomPollution(organisation_id, sectors, bounds, zoom)
    }

    getPollution(organisation_id: string, timeframe: [Date, Date], sector_ids: string[], bounds: Polygon, zoom: number) : Promise<PollutionMatrix> {
        return this.generateRandomPollution(organisation_id, sectors, bounds, zoom)
    }
}

function MonitoringRepositoryFactory(container: any): IMonitoringRepository {
    return new MockMonitoringRepository()
}

export {
    SectorMonitoring, OrganisationMonitoring,
    IMonitoringRepository,
    PollutionMatrix, PollutionTile, PollutionData,
    MonitoringRepositoryFactory
}



