package iter

import (
	"github.com/gpabois/cougnat/core/option"
	opt "github.com/gpabois/cougnat/core/option"
	"github.com/gpabois/cougnat/core/result"
)

type Iterable[T any] interface {
	Iter() Iterator[T]
}

type Iterator[T any] interface {
	Next() opt.Option[T]
}

type FromIterator[T any] interface {
	FromIter(iter Iterator[T])
}

func ForEach[T any](iter Iterator[T], fn func(T)) {
	for el := iter.Next(); el.IsSome(); el = iter.Next() {
		fn(el.Expect())
	}
}

func First[T any](iter Iterator[T]) option.Option[T] {
	for el := iter.Next(); el.IsSome(); {
		return option.Some(el.Expect())
	}
	return option.None[T]()
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

func Reduce[R any, T any](iter Iterator[T], reducer func(R, T) R, init R) R {
	agg := init
	for c := iter.Next(); c.IsSome(); c = iter.Next() {
		v := c.Expect()
		agg = reducer(agg, v)
	}
	return agg
}

func Any[T any](iter Iterator[T], predicate func(T) bool) bool {
	return Reduce(iter, func(agg bool, el T) bool {
		return agg || predicate(el)
	}, false)
}
func All[T any](iter Iterator[T], predicate func(T) bool) bool {
	return Reduce(iter, func(agg bool, el T) bool {
		return agg && predicate(el)
	}, true)
}

func Map[T any, R any](inner Iterator[T], mapper func(T) R) Iterator[R] {
	return MappedIterator[T, R]{
		mapper,
		inner,
	}
}

func Collect[C, T any](fromIter func(iter Iterator[T]) C, iter Iterator[T]) C {
	return fromIter(iter)
}

func CollectToArray[T any](iter Iterator[T]) []T {
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

type IterResult[T any] Iterator[result.Result[T]]

// Take an iterator over a result and reduce it
func Result_FromIter[T any](iter Iterator[result.Result[T]]) result.Result[[]T] {
	var array []T

	for c := iter.Next(); c.IsSome(); c = iter.Next() {
		val := c.Expect()

		if val.HasFailed() {
			return result.Failed[[]T](val.UnwrapError())
		}

		array = append(array, val.Expect())
	}

	return result.Success(array)
}
