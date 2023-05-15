package tuple

type Pair[T1 any, T2 any] struct {
	El0 T1
	El1 T2
}

func NewPair[T1 any, T2 any](el0 T1, el1 T2) Pair[T1, T2] {
	return Pair[T1, T2]{el0, el1}
}

type Triplet[T1 any, T2 any, T3 any] struct {
	El0 T1
	El1 T2
	El2 T3
}

func NewTriplet[T1 any, T2 any, T3 any](el0 T1, el1 T2, el2 T3) Triplet[T1, T2, T3] {
	return Triplet[T1, T2, T3]{el0, el1, el2}
}
