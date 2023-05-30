import { Feature, bboxPolygon } from "@turf/turf";
import {LatLngBounds} from 'leaflet';

export function latLngBoundsToFeature(refBbox: LatLngBounds | Ref<LatLngBounds>): Feature {
    const bbox = unref(refBbox);
    const ne = bbox.getNorthEast();
    const sw = bbox.getSouthWest();
    return bboxPolygon([ne.lng, ne.lat, sw.lng, sw.lat])
}