package result

import opt "github.com/gpabois/cougnat/core/option"

type Result[T any] struct {
	inner opt.Option[T]
	err   error
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
		return res.inner.Expect(), nil
	}
}

func (res Result[T]) Expect() T {
	if res.HasFailed() {
		panic(res.err)
	}

	return res.inner.Expect()
}

func (res Result[T]) UnwrapError() error {
	return res.err
}

func Success[T any](value T) Result[T] {
	return Result[T]{
		inner: opt.Some(value),
		err:   nil,
	}
}

func Failed[T any](err error) Result[T] {
	return Result[T]{
		inner: opt.None[T](),
		err:   err,
	}
}
