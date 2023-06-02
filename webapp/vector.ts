import itertools from '~/itertools'

export interface Vector extends Array<number> {}

export namespace pipeline {
    export function add(...vs: Iterable<number>[]) : Generator<number> {
        return itertools.map(
            itertools.zip(...vs),
            (els) => els.reduce((acc, v) => acc + v, 0)
        )
    }

    export function scalar_mult(v: Iterable<number>, scalar: number): Generator<number> {
        return itertools.map(v, (e1) => e1 * scalar)
    }   
}

// v2 - v1
function diff(v2: Vector, v1: Vector): Vector {
    return itertools.toArray(pipeline.add(
        v2,
        pipeline.scalar_mult(v1, -1.0)
    ))
}

function scalar_product(v2: Vector, v1: Vector): number {
    return itertools.reduce(
        itertools.map(
            itertools.zip(v2, v1),
            ([e2, e1]) => e2 * e1
        ),
        (acc: number, e: number) => acc + e,
        0
    )
}

function norm(v: Vector): number {
    return Math.sqrt(v.map((e) => Math.pow(e, 2)).reduce((acc, e) => acc + e, 0))
}

function normalise(v: Vector): Vector {
    const n = norm(v)
    return itertools.toArray(pipeline.scalar_mult(v, 1.0/n))
}


function walk(begin: Vector, end: Vector, lambda: number): Vector {
    return itertools.toArray(pipeline.add(
        pipeline.scalar_mult(end, lambda),
        pipeline.scalar_mult(begin, 1.0 - lambda)
    ))
}

export default {
    diff,
    walk,
    norm,
    normalise,
    scalar_product,
    pipeline
}