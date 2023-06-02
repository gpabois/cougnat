export function all(values: Iterable<boolean>): boolean {
    for(const v of values) {
        if(!v) 
            return false
    }
    return true;
}

export function any(values: Iterable<boolean>): boolean {
    for(const v of values) {
        if(v) return true
    }

    return false;
}

export function * zip<T>(...vs: Iterable<T>[]): Generator<T[]> {
    const its = vs.map((i) => i[Symbol.iterator]())

    const next = function<T>(its: Iterator<T>[]) : Array<IteratorResult<T>> {
        return its.map((it) => it.next())
    }

    const values = function<T>(ns: IteratorResult<T>[]): Array<T> {
        return ns.map((n) => n.value)
    }

    const at_least_one_is_exhausted = function<T>(ns: IteratorResult<T>[]): boolean {
        const result = any(ns.map((n) => n.done === true))
        return result
    }
    
    for(let ns = next(its); !at_least_one_is_exhausted(ns); ns = next(its)) yield values(ns)
}

export function * map<T, R>(iter: Iterable<T>, mapper: (el: T) => R): Generator<R> {
    for(const el of iter) yield mapper(el)
}

export function reduce<T>(iter: Iterable<T>, reducer: (acc: T, el: T) => T, init: T): T {
    let acc = init;
    for(const el of iter)
        acc = reducer(acc, el)
    return acc
}

export function toArray<T>(iter: Iterable<T>): Array<T> {
    var arr: Array<T> = [];
    for(const el of iter) arr.push(el)
    return arr;
}

export namespace async {
    export async function toArray<T, TReturn = any, TNext = unknown>(asyncIter: AsyncIterable<T>): Promise<Array<T>> {
        var arr = [];
        for await (const el of asyncIter) {
            arr.push(el)
        }
        return arr;
    }
}

export default {
    zip, map, reduce, toArray,
    async
}