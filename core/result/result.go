package result

import (
	"reflect"
)

type IResult interface {
	TypeOf() reflect.Type
	UnwrapError() error
	HasFailed() bool
	IsSuccess() bool
}

type Result[T any] struct {
	inner T
	err   error
}

func Flatten[T any](value Result[Result[T]]) Result[T] {
	if value.HasFailed() {
		return Failed[T](value.UnwrapError())
	} else {
		return value.Expect()
	}
}

func FlatMap[T any, U any](val Result[T], mapper func(val T) Result[U]) Result[U] {
	return Flatten(Map(val, mapper))
}

func ChainMap[T any, U any](mapper func(val T) Result[U], val Result[T]) Result[U] {
	return FlatMap(val, mapper)
}

func (res Result[T]) Then(then func(val T)) Result[T] {
	if res.IsSuccess() {
		then(res.Expect())
	}

	return res
}

func Map[T any, U any](val Result[T], mapper func(val T) U) Result[U] {
	if val.HasFailed() {
		return Failed[U](val.err)
	} else {
		return Success(mapper(val.Expect()))
	}
}

func (res Result[T]) HasFailed() bool {
	return res.err != nil
}

func (res Result[T]) IsSuccess() bool {
	return res.err == nil
}

func (res Result[T]) IntoAnyTuple() (any, error) {
	if res.HasFailed() {
		return nil, res.err
	} else {
		return res, nil
	}
}

func (res Result[T]) Expect() T {
	if res.HasFailed() {
		panic(res.err)
	}

	return res.inner
}

func (res Result[T]) UnwrapError() error {
	return res.err
}

func Success[T any](value T) Result[T] {
	return Result[T]{
		inner: value,
		err:   nil,
	}
}

func Failed[T any](err error) Result[T] {
	return Result[T]{
		err: err,
	}
}
