package iter

import (
	"reflect"

	opt "github.com/gpabois/cougnat/core/option"
)

type MapIterator[K comparable, V any, M ~map[K]V] struct {
	keyIt Iterator[K]
	inner *M
}

func (it *MapIterator[K, V, M]) Next() opt.Option[KV[K, V]] {
	key := it.keyIt.Next()
	if key.IsNone() {
		return opt.None[KV[K, V]]()
	}

	if val, ok := (*it.inner)[key.Expect()]; ok {
		return opt.Some(KV[K, V]{Key: key.Expect(), Value: val})
	}

	return opt.None[KV[K, V]]()
}

func IterMap[K comparable, V any, M ~map[K]V](m *M) Iterator[KV[K, V]] {
	keys := reflect.ValueOf(m).MapKeys()
	return &MapIterator[K, V, M]{
		keyIt: Map(IterSlice(&keys), func(value reflect.Value) K { return value.Interface().(K) }),
		inner: m,
	}
}
