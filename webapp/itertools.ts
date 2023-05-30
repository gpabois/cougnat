async function toArray<T, TReturn = any, TNext = unknown>(asyncIter: AsyncIterable<T>): Promise<Array<T>> {
    var arr = [];
    for await (const el of asyncIter) {
        arr.push(el)
    }
    return arr;
}

export default {toArray}