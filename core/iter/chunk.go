package iter

import "github.com/gpabois/cougnat/core/option"

type Chunk[T any] []T

type ChunkIterator[T any] struct {
	chunkSize int
	inner     Iterator[Enumeration[T]]
}

func ChunkEvery[T any](it Iterator[T], chunkSize int) Iterator[Chunk[T]] {
	return &ChunkIterator[T]{
		chunkSize: chunkSize,
		inner:     Enumerate(it),
	}
}

func (chunkIter *ChunkIterator[T]) Next() option.Option[Chunk[T]] {
	var acc Chunk[T]

	for c := chunkIter.inner.Next(); c.IsSome(); c = chunkIter.inner.Next() {
		acc = append(acc, c.Expect().Second)
		if c.Expect().First%chunkIter.chunkSize == 0 {
			return option.Some(acc)
		}
	}

	return option.Some(acc)
}
