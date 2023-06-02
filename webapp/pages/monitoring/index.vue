<template>
    <div class="container-fluid vh-100 vw-100 g-0 d-flex flex-column">
        <div class="row g-0 flex-grow-1">
            <Transition name="slide-x">
                <div class="col-3 shadow-lg" style="z-index: 1000; position: relative; overflow: auto; max-height: 100%;" v-if="panelState">
                    <div style="width: 100%; position: absolute; top: 0;" class="d-flex flex-column align-items-end p-1">
                        <button class="btn btn-lg btn-link text-dark" @click="closePanel()">
                            <i class="bi bi-x-circle-fill border-white"></i>
                        </button>
                    </div>
                    <div class="p-3" v-if="panelState.mode == 'broadcast'">
                        <h1 class="h1 mb-2">Communication</h1>

                        <div class="mb-3">
                            <div class="form-group mb-3">
                                <textarea class="form-control" rows="10">

                                </textarea>
                            </div>
                            <div class="form-group">
                                <button class="btn btn-primary"><i class="bi bi-broadcast text-white"></i> Diffuser</button>
                            </div>
                        </div>

                        <h2 class="h2">Répondeurs en cours</h2>
                        <ul class="list-group">
                            <li class="list-group-item" v-for="broadcast in broadcasts?.filter((b) => b.closed_at == null)">
                                <div class="d-flex w-100 justify-content-between">
                                    <span>{{ broadcast.content }}</span>
                                    <small v-if="broadcast.closed_at == null">
                                        <i class="bi bi-x-circle-fill"></i>
                                    </small>
                                 </div>
                            </li>
                        </ul>
                    </div>
                    <div class="p-3" v-if="panelState.mode == 'stats'">
                        <h1 class="h1 mb-2">Statistiques</h1>
                        <Responsive class="w-100">
                            <template #main="{ width }">
                                <Chart 
                                    :config="{ controlHover: false }"
                                    :size="{ width, height: 200 }"
                                    direction="circular" :data="reports_per_type">
                                    <template #layers>
                                        <Pie :data-keys="['type', 'count']" :pie-style="{ innerRadius: 50, padAngle: 0.05 }" ></Pie>
                                    </template>
                                    <template #widgets>
                                    <Tooltip
                                        :config="{
                                        type: { },
                                        count : {label: 'value'}
                                        }"
                                        hideLine
                                    />
                                    </template>
                                </Chart>
                            </template>
                        </Responsive>
                    </div>
                    <div class="p-3" v-if="panelState.mode == 'config'">
                        <h1 class="h1 mb-2">Configuration</h1>
                        <h2 class="h2">En tant que</h2>
                        <select class="form-select" v-model="current_organisation">
                            <option v-for="organisation in organisations" :value="organisation">
                                {{ organisation.name }}
                            </option>
                        </select>
                        <template v-if="current_organisation">
                            <h2 class="h2">Secteur(s)</h2>
                            <select class="form-select" multiple v-model="selected_sector_monitorings">
                                <option v-for="sector in sector_monitorings.features" :value="sector">{{ sector.properties.label }}</option>
                            </select>
                        </template>
                    </div>
                </div>
            </Transition>
            <div class="col" style="z-index: 0; position: relative">
                <div class="h-100">
                    <l-map ref="map" :useGlobalLeaflet="false" v-model:bounds="bounds" v-model:center="center" v-model:zoom="zoom" @ready="onReady" >    
                        <l-geo-json 
                            :geojson="pollution_tiles" 
                            :options-style="applyHeatmap"
                            ></l-geo-json>
                        <l-geo-json 
                            :geojson="selected_sector_monitorings" 
                            :options-style="() => ({fillOpacity: 0.0, opacity: 1.0})"></l-geo-json>

                        <l-tile-layer v-for="layer in layers" 
                            :attribution="layer.attribution" 
                            :url="layer.url"
                        ></l-tile-layer>
                    </l-map>
                </div>
                <div class="p-3 w-100 d-flex flex-column align-items-end" style="position: absolute; top: 0; z-index: 2000">
                    <div class="btn-group">
                        <button class="btn-primary btn" @click="panelState = {mode: 'config'}">
                            <i class="bi bi-gear-fill"></i>
                        </button>
                        <button class="btn btn-success" @click="panelState = {mode: 'stats'}">
                            <i class="bi bi-bar-chart-fill"></i>
                        </button>
                        <button class="btn btn-warning" @click="panelState = {mode: 'broadcast'}" v-if="current_organisation?.permissions?.includes('broadcast')">
                            <i class="bi bi-broadcast text-white"></i>
                        </button>
                        <button class="btn btn-danger" @click="panelState = {mode: 'history'}">
                            <i class="bi bi-clock"></i>
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import "leaflet/dist/leaflet.css"
import { Chart, Pie, Responsive } from 'vue3-charts'
import { LMap, LGeoJson, LTileLayer, LCircle} from "@vue-leaflet/vue-leaflet";
import { Feature, FeatureCollection, bboxPolygon, featureCollection } from "@turf/turf";
import { LatLngBounds, Polygon } from "leaflet";
import { latLngBoundsToFeature } from "~/geo/utils";
import {Broadcast} from '~/repository/communication';

const {$api} = useNuxtApp()

import {Organisation} from '~/repository/organisation'
import {PollutionMatrix, PollutionTile, PollutionData, SectorMonitoring} from '~/repository/monitoring'
import {tiles} from '~/geo/utils'
import color from "~/color";

const layers = [{
    url: 'https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png',
    attribution: '&copy; <a target="_blank" href="http://osm.org/copyright">OpenStreetMap</a> contributors'
}]

const panelState = ref<any>(null);
const map = ref(null);
const center = ref([48.77976043401817, 2.488472329373121]);
const zoom = ref(13);
const bounds = ref<LatLngBounds | null>(null);

const box = computed(() => {
    if(bounds.value) {
        return latLngBoundsToFeature(bounds.value!).geometry
    } else {
        return null
    }
})

const closePanel = () => { panelState.value = null }

const {data: organisations} = useAsyncData(() => $api.organisation.GetMine())
const current_organisation = ref<Organisation | null>(null);
const selected_sector_monitorings = ref<Feature<Polygon, SectorMonitoring>[]>([]);
const {data: org_monitoring} = useAsyncData('fetchMonitoringPrimeter', async () => {
    if(!current_organisation.value) {
        return null;
    }
    return await $api.monitoring.getOrganisationMonitoring(current_organisation.value.id)
}, {
    watch: [current_organisation]
})
const sector_monitorings = computed(() => {
    if(!org_monitoring.value) return featureCollection([])
    return org_monitoring.value.sectors
})

const {data: broadcasts}  = useAsyncData<Array<Broadcast>>(async () => {
    if(!current_organisation.value) return []
    return $api.communication.GetBroadcastsByOrganisation(current_organisation.value.id)
}, {
    watch: [current_organisation]
})

const {data: pollution, refresh: refreshPollution, error, pending: loadingPollution} = useAsyncData<PollutionMatrix>('fetchPollutionTiles', async () => {
    if(!current_organisation.value || !box.value) return []

    const sector_ids = [...selected_sector_monitorings.value || []].map((f) => f.properties.id)

    if(sector_ids.length == 0) return []

    return await $api.monitoring.getCurrentPollution(
        current_organisation.value!.id, 
        sector_ids, 
        box.value!, 
        (zoom.value + 4)
    )
}, {
    watch: [current_organisation, selected_sector_monitorings, box, zoom]
})

const reports_per_type = computed(() => {
    var types = {}
    for(const tile of pollution.value) {
        for(const per_type of tile.types) {
            if(types[per_type.type] === undefined) {
                types[per_type.type] = 0
            }
            types[per_type.type] += per_type.count
        }
    }

    var stats = []
    for(const [k,v] of Object.entries(types)) {
        stats.push({type: k, count: v})
    }
    return stats
})

/**
 * Convert a pollution matrix into a feature collection to be displayed on the leaflet map.
 */
const pollution_tiles = computed(() => {
    return featureCollection(
        pollution.value?.map((tile: PollutionTile) => {
            const weight = tile.types.map((polData: PollutionData) => polData.count).reduce((acc, w) => acc + w, 0)
            if(weight == 0) return null;
            return bboxPolygon(tiles.toBox(tile.coordinates), {
                properties: {
                    weight
                }
            })
        }).filter((tile) => tile !== null) || []
    )
})

/**
 * Compute the color depending of the weight of the pollution tile
 * @param feature 
 */
const applyHeatmap = function(feature: Feature) {
    const w = feature.properties?.weight | 0.0
    const lambda = Math.max(0.0, Math.min(w / 100.0, 1.0));
    const hue_c = color.walk({h: 0.0, s: 1.0, l: 0.5}, {h: 0.3, s: 0.0, l: 0.5}, 1.0 - lambda)
    const c = color.hue.toRGB(hue_c);
    const hC = color.rgb.toHex(c)
    return {
        opacity: 0.0,
        fillOpacity: w == 0 ? 0.0 : 0.6,
        fillColor: hC
    }
}

watchEffect(() => {
    current_organisation.value;
    selected_sector_monitorings.value = []
})

const interval_id = ref(null);
const onReady = () => { bounds.value = map.value?.leafletObject?.getBounds() }
onMounted(() => {
    interval_id.value = setInterval(() => {
        if(loadingPollution.value)
            return
        refreshPollution()
    }, 5000)
})
onBeforeUnmount(() => clearInterval(interval_id.value))
</script>