package ops

import "golang.org/x/exp/constraints"

type Number interface {
	int | float32 | float64
}

type Interval[T constraints.Ordered] struct {
	begin T
	end   T
}

func (itv Interval[T]) Max() T {
	return Max(itv.begin, itv.end)
}

func (itv Interval[T]) Min() T {
	return Min(itv.begin, itv.end)
}

func Add[T Number](values ...T) T {
	var acc T
	for _, value := range values {
		acc = acc + value
	}
	return acc
}

func Max[T constraints.Ordered](acc T, values ...T) T {
	for _, value := range values {
		if value > acc {
			acc = value
		}
	}
	return acc
}

func Max2[T constraints.Ordered](a, b T) T {
	return Max(a, b)
}

func Min[T constraints.Ordered](acc T, values ...T) T {
	for _, value := range values {
		if value < acc {
			acc = value
		}
	}
	return acc
}

func Min2[T constraints.Ordered](a, b T) T {
	return Min(a, b)
}

func Add2[T Number](a, b T) T {
	return Add(a, b)
}

func IsTrue(b bool) bool {
	return b
}

func IsFalse(b bool) bool {
	return !b
}
