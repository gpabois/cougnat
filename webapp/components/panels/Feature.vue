<template>
    <div class="thumbnail" v-if="feature?.properties?.thumbnail" style="max-height: 250px; overflow: hidden">
        <img class="img-fluid" :src="feature?.properties?.thumbnail"/>
    </div>
    <div class="p-3">
        <h1 class="h1" v-if="feature?.properties?.name">
            {{ feature?.properties?.name }}
            <small>
                <i class="bi bi-box-arrow-up-left "></i>
            </small>
        </h1>
        <p class="fw-light lead" v-if="feature?.properties?.description">
            {{ feature?.properties?.description }}
        </p>
        <client-only>
            <star-rating 
                v-if="feature?.properties?.rating" 
                :star-size="15"
                :rating="feature?.properties?.rating">
            </star-rating>
        </client-only>
        <hr/>
        <ul class="list-group list-group-flush">
            <li v-if="feature?.properties?.address" class="list-group-item align-items-center d-flex">
                <i class="bi bi-geo-alt mx-2"></i>
                {{ feature?.properties?.address }}
            </li>
        </ul>
    </div>
</template>
<script setup lang="ts">
import { Feature } from '@turf/turf';
const {feature} = defineProps<{feature: Feature}>();
</script>