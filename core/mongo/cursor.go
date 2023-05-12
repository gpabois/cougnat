package mongo

import (
	"context"

	"github.com/gpabois/cougnat/core/option"
	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/core/serde"
	"go.mongodb.org/mongo-driver/mongo"
)

// Cursor over model-based elements
type Cursor[T any] struct {
	ctx   context.Context
	inner *mongo.Cursor
}

func (it Cursor[T]) Next() option.Option[result.Result[T]] {
	if it.inner.Next(it.ctx) {
		rawEls, err := it.inner.Current.Elements()
		if err != nil {
			return option.Some(result.Failed[T](err))
		}

		rawEl := rawEls[0]
		return option.Some(serde.UnMarshalBson[T](rawEl))
	} else {
		return option.None[result.Result[T]]()
	}
}

// Create a cursor iterator
func IterCursor[T any](ctx context.Context, cursor *mongo.Cursor) Cursor[T] {
	return Cursor[T]{ctx, cursor}
}
