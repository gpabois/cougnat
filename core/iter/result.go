package iter

import "github.com/gpabois/cougnat/core/result"

type IterResult[T any] Iterator[result.Result[T]]

// Take an iterator over a result and reduce it
func Result_FromIter[T any, S ~[]T](iter Iterator[result.Result[T]]) result.Result[S] {
	var array S

	for c := iter.Next(); c.IsSome(); c = iter.Next() {
		val := c.Expect()

		if val.HasFailed() {
			return result.Failed[S](val.UnwrapError())
		}

		array = append(array, val.Expect())
	}

	return result.Success(array)
}