<template>
    <div class="container-fluid vh-100 vw-100 g-0 d-flex flex-column">
        <div class="row g-0 flex-grow-1">
            <Transition name="slide-x">
                <div class="col-3 shadow-lg" style="z-index: 1000; position: relative; overflow: scroll; max-height: 100%;" v-if="panelState">
                    <div style="width: 100%; position: absolute; top: 0;" class="d-flex flex-column align-items-end p-1">
                        <button class="btn btn-lg btn-link text-dark" @click="closePanel()">
                            <i class="bi bi-x-circle-fill border-white"></i>
                        </button>
                    </div>
                </div>
            </Transition>
            <div class="col" style="z-index: 0; position: relative">
                <div class="h-100">
                    <l-map ref="map" :useGlobalLeaflet="false" v-model:bounds="bounds" v-model:center="center" v-model:zoom="zoom" >
                        <l-tile-layer v-for="layer in layers" 
                            :attribution="layer.attribution" 
                            :url="layer.url"
                        ></l-tile-layer>
                    </l-map>
                </div>
                <div class="p-3 w-100 d-flex flex-column align-items-end" style="position: absolute; top: 0; z-index: 2000">
                    <div class="btn-group">
                        <select class="form-select" v-model="current_organisation">
                            <option v-for="organisation in organisations">
                                {{ organisation.name }}
                            </option>
                        </select>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import { LMap, LGeoJson, LTileLayer, LCircle} from "@vue-leaflet/vue-leaflet";
import { FeatureCollection } from "@turf/turf";
import { LatLngBounds, Polygon } from "leaflet";
import "leaflet/dist/leaflet.css"
import { latLngBoundsToFeature } from "~/geo/utils";
const {$api} = useNuxtApp()

import {Organisation} from '~/repository/organisation'

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

const closePanel = function() {
    panelState.value = null 
}


const current_organisation = ref<Organisation | null>(null);
const monitoring_perimeter = ref<FeatureCollection | null>(null);
const {data: organisations} = useAsyncData(() => $api.organisation.GetMine())
const {data: pollutionTiles, refresh: refreshPollutionTiles} = useAsyncData(
    'fetchPollutionTiles', async () => {
    
    if(!current_organisation.value) {
        return []
    }

    if(!box.value) {
        return []
    }

    return await $api.monitoring.GetCurrentPollution(current_organisation.value!.id, box.value!, 13)
}, {
    watch: [current_organisation, box, zoom]
})

onBeforeMount(() => setInterval(() => refreshPollutionTiles(), 1000));

</script>