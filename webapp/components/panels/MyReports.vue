<template>
<div class="p-3">
    <h1 class="h1 mb-3"><i class="bi bi-flag"></i> Mes signalements</h1>
    <ul class="list-group">
        <li v-for="report in reports" class="list-group-item">
            <div class="d-flex w-100 justify-content-between">
                <h5 class="mb-1">{{ report.type.label }}</h5>
                <small>{{ fromNow(report.reported_at!) }}</small>
            </div>
        </li>
    </ul>
</div>
</template>
<script setup lang="ts">
import moment from 'moment'
import 'moment/locale/fr';
const {$api} = useNuxtApp();
const {data: reports} = useAsyncData('fetchMyReports', () => $api.reports.GetMine()); 
const fromNow = function(date: Date): string {
    moment.locale('fr');
    return moment(date).fromNow()
}
</script>