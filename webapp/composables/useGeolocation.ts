import { Feature, point} from "@turf/turf";

export default function() {
    const position = ref<Feature | null>(null);

    const getPosition = function(pos: GeolocationPosition) {
        position.value = point([pos.coords.latitude, pos.coords.longitude])
    }
    
    if (process.client && navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(getPosition);
    }

    return {position}
}