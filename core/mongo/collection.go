package mongo

import (
	"context"

	"github.com/gpabois/cougnat/core/iter"
	"github.com/gpabois/cougnat/core/option"
	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/core/serde"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mng "go.mongodb.org/mongo-driver/mongo"
)

type Collection[T any] struct {
	inner *mng.Collection
}

func (coll Collection[T]) InsertOne(ctx context.Context, doc T) result.Result[primitive.ObjectID] {
	rawRes := serde.MarshalBson(doc)

	if rawRes.HasFailed() {
		return result.Failed[primitive.ObjectID](rawRes.UnwrapError())
	}

	insRes, err := coll.inner.InsertOne(ctx, rawRes)
	if err != nil {
		return result.Failed[primitive.ObjectID](err)
	}

	return result.Success(insRes.InsertedID.(primitive.ObjectID))
}

func (coll Collection[T]) Find(ctx context.Context, filter any) result.Result[iter.IterResult[T]] {
	cursor, err := coll.inner.Find(ctx, filter)
	if err == nil {
		return result.Failed[iter.IterResult[T]](err)
	}

	it := IterCursor[T](ctx, cursor)
	return result.Success[iter.IterResult[T]](it)
}

func (coll Collection[T]) FindOne(ctx context.Context, filter any) result.Result[option.Option[T]] {
	resFind := coll.Find(ctx, filter)

	if resFind.HasFailed() {
		return result.Failed[option.Option[T]](resFind.UnwrapError())
	}

	it := resFind.Expect()
	return option.Swap[T](iter.First[result.Result[T]](it))
}
