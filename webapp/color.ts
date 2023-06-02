import vector, { Vector } from "./vector";

interface Hue {
    h: number,
    s: number,
    l: number
}

interface RGB {
    r: number,
    g: number,
    b: number
}

export namespace rgb {
    function _byteToHex(c: number): string {
        const h = c.toString(16);
        return h.length == 1 ? "0" + h : h;
    }

    export function toHex(color: RGB): string {
        return `#${_byteToHex(color.r)}${_byteToHex(color.g)}${_byteToHex(color.b)}`
    }
}

export namespace hue {
    export function toVec(hue: Hue): Vector {
        return [hue.h, hue.s, hue.l]
    }
    
    export function fromVec(v: [number, number, number]): Hue {
        return {
            h: v[0],
            s: v[1],
            l: v[2]
        }
    }

    export function toRGB(hue: Hue): RGB {
        var rgb: RGB = {r: 0, g: 0, b: 0};

        if(hue.s == 0) {
            rgb.r = rgb.g = rgb.b = hue.l;
        } else {
            var hue2rgb = function hue2rgb(p: number, q: number, t: number){
                if(t < 0) t += 1;
                if(t > 1) t -= 1;
                if(t < 1/6) return p + (q - p) * 6 * t;
                if(t < 1/2) return q;
                if(t < 2/3) return p + (q - p) * (2/3 - t) * 6;
                return p;
            }
            var q = hue.l < 0.5 ? hue.l * (1 + hue.s) : hue.l + hue.s - hue.l * hue.s;
            var p = 2 * hue.l - q;
            rgb.r = hue2rgb(p, q, hue.h + 1/3);
            rgb.g = hue2rgb(p, q, hue.h);
            rgb.b = hue2rgb(p, q, hue.h - 1/3);
        }

        rgb = {
            r: Math.round(rgb.r * 255),
            g: Math.round(rgb.g * 255),
            b: Math.round(rgb.b * 255)
        }

        return rgb;
    }
}


function walk(min: Hue, max: Hue, distance: number): Hue {
    return hue.fromVec(
            vector.walk(
                hue.toVec(min),
                hue.toVec(max),
                distance
            ) as [number, number, number]
    )
}

export default {
    walk,
    hue, 
    rgb
}