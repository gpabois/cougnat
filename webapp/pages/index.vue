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
                    <PanelsFeature :feature="panelState?.feature!" v-if="panelState?.mode == 'feature'"></PanelsFeature>
                    <PanelsMyReports v-else-if="panelState?.mode == 'my-reports'"></PanelsMyReports>
                </div>
            </Transition>
            <div class="col" style="z-index: 0; position: relative">
                <div class="cg-map h-100">
                    <l-map ref="map" :useGlobalLeaflet="false" v-model:bounds="bounds" v-model:center="center" v-model:zoom="zoom" >
                        <l-circle :radius="100" :latLng="position?.geometry?.coordinates" v-if="position">
                        </l-circle>
                        <l-geo-json 
                            name="features"
                            :geojson="features"
                            :options="featureOptions"
                        ></l-geo-json>
                        <l-tile-layer v-for="layer in layers" 
                            :attribution="layer.attribution" 
                            :url="layer.url"
                        ></l-tile-layer>
                    </l-map>
                </div>
                <div class="p-3 w-100 d-flex flex-column align-items-end" style="position: absolute; top: 0; z-index: 2000">
                    <div class="btn-group">
                        <button type="button" class="btn btn-success" @click="centerOnPosition()" :disabled="!position">
                            <i class="bi bi-geo"></i>
                        </button>
                        <button type="button" class="btn btn-primary" @click="displayReportForm = !displayReportForm">
                            <i class="bi bi-megaphone-fill" v-if="!displayReportForm"></i>
                            <i class="bi bi-x-circle-fill" v-else></i>
                        </button>
                        <NuxtLink class="btn btn-warning" :to="{name: 'monitoring'}"><i class="bi bi-binoculars text-white"></i></NuxtLink>
                        <button type="button" class="btn btn-light" @click="selectMyReports()">Archive</button>
                    </div>
                </div>
            </div>
        </div>
        <Transition name="slide-y">
            <div class="row bg-white" v-if="displayReportForm" style="z-index: 3000">
                <div class="col">
                    <FormsNewReport @created="onReport"></FormsNewReport>
                </div>
            </div>
        </Transition>
    </div>
   
</template>
<style>
.slide-x-leave-active,
.slide-x-enter-active {
  transition: 0.3s;
}
.slide-x-enter {
  transform: translate(100%, 0);
}
.slide-x-leave-to {
  transform: translate(-100%, 0);
}

.slide-y-leave-active,
.slide-y-enter-active {
  transition: 0.3s;
}
.slide-y-enter {
  transform: translate(0, -100%);
}
.slide-y-leave-to {
  transform: translate(0, 100%);
}

</style>
<script setup lang="ts">
import "leaflet/dist/leaflet.css"
import {latLngBoundsToFeature} from '../geo/utils';
import { Feature, FeatureCollection, featureCollection, point, bboxPolygon, Point} from "@turf/turf";
import { LMap, LGeoJson, LTileLayer, LCircle} from "@vue-leaflet/vue-leaflet";
import {LatLng, Layer} from 'leaflet';

const {$api} = useNuxtApp()
const {position} = useGeolocation();

const layers = [{
    url: 'https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png',
    attribution: '&copy; <a target="_blank" href="http://osm.org/copyright">OpenStreetMap</a> contributors'
}]

const displayReportForm = ref(false);
const panelState = ref<any>(null);
const map = ref(null);
const center = ref([48.77976043401817, 2.488472329373121]);
const zoom = ref(13);
const bounds = ref(null);
const features = ref<FeatureCollection>(featureCollection([]))

const fetchFeatures = async function() {
    if(bounds.value)
        features.value = await $api.features.GetWithin(latLngBoundsToFeature(bounds))  
}

const centerOnPosition = function() {
    const m = map.value?.leafletObject;
    m.flyTo(position.value?.geometry.coordinates, 15);
}

const selectFeature = function(feature: Feature) {
    panelState.value = {mode: 'feature', feature}
}

const selectMyReports = function() {
    panelState.value = {mode: 'my-reports'}
}

const closePanel = function() {
    panelState.value = null;
}

// Events recevivers
const onFeatureClicked = selectFeature;
const onEachFeature = function(feature: Feature, layer: Layer) {
    layer.on('click', function(e) {
        onFeatureClicked(feature)
    })
    switch(feature.properties?.type) {
        case "facility":
    }
}

const onReport = function(_e: any) {
    displayReportForm.value = false
}

const featureOptions = ref({
    onEachFeature,
})

watchEffect(fetchFeatures)
</script>