package option

type Option[T any] struct {
	value *T
}

func (opt Option[T]) IsSome() bool {
	return opt.value != nil
}

func (opt Option[T]) IsNone() bool {
	return opt.value == nil
}

func (opt Option[T]) Expect() T {
	if !opt.IsSome() {
		panic("empty value")
	}

	return *opt.value
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

func None[T any]() Option[T] {
	return Option[T]{value: nil}
}
