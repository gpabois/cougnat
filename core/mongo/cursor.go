package mongo

import (
	"context"

	"github.com/gpabois/cougnat/core/iter"
	"github.com/gpabois/cougnat/core/option"
	"go.mongodb.org/mongo-driver/mongo"
)

// Wrap a mongo cursor into an iterator
type CursorIterator struct {
	cursor mongo.Cursor
}

// Create a cursor iterator
func IterCursor(cursor mongo.Cursor) CursorIterator {
	return CursorIterator{cursor}
}

func (it CursorIterator) Next() option.Option[mongo.Cursor] {
	if !it.cursor.Next(context.TODO()) {
		return option.None[mongo.Cursor]()
	}

	return option.Some(it.cursor)
}

// Decode the cursor elements
func Decode[T any](it CursorIterator) iter.Iterator[T] {
	return iter.Map[mongo.Cursor](it, func(cursor mongo.Cursor) T {
		var val T
		cursor.Decode(&val)
		return val
	})
}
