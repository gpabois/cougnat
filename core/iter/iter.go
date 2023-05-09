package iter

import opt "github.com/gpabois/cougnat/core/option"

type Iterable[T any] interface {
	Iter() Iterator[T]
}

type Iterator[T any] interface {
	Next() opt.Option[T]
}

func ForEach[T any](iter Iterator[T], fn func(T)) {
	for el := iter.Next(); el.IsSome(); el = iter.Next() {
		fn(el.Expect())
	}
}

func Filter[T any](iter Iterator[T], filter func(T) bool) Iterator[T] {
	return FilteredIterator[T]{
		inner:  iter,
		filter: filter,
	}
}

func Find[T any](iter Iterator[T], filter func(T) bool) opt.Option[T] {
	it := Filter(iter, filter)
	return it.Next()
}

func Map[T any, R any](inner Iterator[T], mapper func(T) R) Iterator[R] {
	return MappedIterator[T, R]{
		mapper,
		inner,
	}
}

func Collect[T any](iter Iterator[T]) []T {
	var array []T
	ForEach(iter, func(el T) {
		array = append(array, el)
	})
	return array
}

type MappedIterator[T any, R any] struct {
	mapper func(T) R
	inner  Iterator[T]
}

func (iter MappedIterator[T, R]) Next() opt.Option[R] {
	el := iter.inner.Next()
	if el.IsNone() {
		return opt.None[R]()
	}
	val := el.Expect()
	return opt.Some(iter.mapper(val))
}

type FilteredIterator[T any] struct {
	filter func(T) bool
	inner  Iterator[T]
}

func (iter FilteredIterator[T]) Next() opt.Option[T] {
	for el := iter.Next(); el.IsSome(); el = iter.Next() {
		value := el.Expect()
		if iter.filter(value) {
			return el
		}
	}
	return opt.None[T]()
}

type ArrayIterator[T any] struct {
	array  *[]T
	cursor int
}

func IterArray[T any](array *[]T) Iterator[T] {
	return ArrayIterator[T]{
		array:  array,
		cursor: 0,
	}
}

func (iter ArrayIterator[T]) Next() opt.Option[T] {
	if iter.cursor >= len(*iter.array) {
		return opt.None[T]()
	}
	return opt.Some((*iter.array)[iter.cursor])
}
