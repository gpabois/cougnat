<template>
    <div class="container-fluid p-4 shadow">
        <h1 class="h1 mb-4">Signaler une nuisance</h1>
        <div class="row align-items-center mb-3">
            <div class="col-2">Nature</div>
            <div class=" col">
                <div class="input-group">
                <input type="checkbox" class="btn-check" id="btn-check-smell" value="smell" v-model="selectedNatures">
                <label class="btn btn-outline-primary" for="btn-check-smell">Odeur</label><br>

                <input type="checkbox" class="btn-check" id="btn-check-noise" value="noise" v-model="selectedNatures">
                <label class="btn btn-outline-primary" for="btn-check-noise">Bruit</label><br>
                </div>
            </div>
        </div>
        <div class="row align-items-center mb-3">
            <div class="col-2">Localisation</div>
            <div class="col">
                <i class="bi bi-check-square-fill text-primary" v-if="report.location"></i>
                <i class="bi bi-x-square-fill text-danger" v-else></i>
            </div>
        </div>
        <div class="row align-items-center mb-3" v-if="types?.length">
            <div class="col-2">Nuisance</div>
            <div class="col">
                <select class="form-select" v-model="report.type">
                    <option v-for="t in types" :value="t">
                        {{ t.label }}
                    </option>
                </select>
            </div>
        </div>
        <div class="row align-items-center mb-3">
            <button class="btn btn-primary col" :disabled="!isValid" @click="create">Signaler</button>
        </div>
    </div>
</template>
<script setup lang="ts">
import { Feature } from '@turf/helpers';
import { Report, ReportType } from '~/repository/report';

const {$api} = useNuxtApp();
const emits = defineEmits(['created'])
const {position} = useGeolocation();
const report = reactive<{
    type: ReportType | null,
    location: Feature | null
}>({
    type: null,
    location: null
})

const selectedNatures = ref([]);
const {data: allTypes} = useAsyncData('fetchReportTypes', () => $api.reports.GetTypes()); 

const types = computed(function() {
    return allTypes.value?.filter((type) => selectedNatures.value.includes(type.nature as never))
});

const isValid = computed(function() {
    return report.type && report.location
})

watchEffect(function() {
    report.location = unref(position)
})

const create = async function() {
    if (!isValid.value) {
        throw "invalid report"
    }

    const report_id = await $api.reports.Create(report as Report);
    emits("created", report_id)
}

</script>