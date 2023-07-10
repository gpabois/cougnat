package pg

import (
	"context"
	"errors"
	"reflect"

	"github.com/gpabois/cougnat/core/transaction"
	"github.com/gpabois/gostd/geojson"
	"github.com/gpabois/gostd/iter"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
	"github.com/gpabois/gostd/serde"
	"github.com/jackc/pgx/v5"
)

// Transform the query arguments (option, geojson, ...)
func transformQueryArgs(args []any) result.Result[[]any] {
	return iter.Result_FromIter[[]any](iter.Map(iter.IterSlice(&args), func(arg any) result.Result[any] {
		if option.Reflect_IsOptionType(reflect.TypeOf(arg)) {
			innerOpt := option.Reflect_Get(reflect.ValueOf(arg))
			if innerOpt.IsSome() {
				return result.Success[any](innerOpt.Expect())
			} else {
				return result.Success[any](nil)
			}
		}

		switch a := arg.(type) {
		// Encode GeoJSON
		case geojson.Feature, geojson.Geometry:
			serRes := serde.Serialize(a, "application/json")
			if serRes.HasFailed() {
				return result.Failed[any](serRes.UnwrapError())
			}
			return result.Success[any](serRes)
		}

		return result.Success(arg)
	}))
}

type WithQueryFunc[R any] func(rows pgx.Rows) result.Result[R]

func WithQuery[R any](fn WithQueryFunc[R]) func(tx pgx.Tx, ctx context.Context, query string, args ...any) result.Result[R] {
	return func(tx pgx.Tx, ctx context.Context, query string, args ...any) result.Result[R] {
		res := Query(tx, ctx, query, args...)
		if res.HasFailed() {
			return result.Result[R]{}.Failed(res.UnwrapError())
		}

		rows := res.Expect()
		defer rows.Close()
		return fn(rows)
	}
}

type WithExpectedOneFunc[R any] func(rows pgx.Rows) result.Result[R]

func WithExpectedOne[R any](fn WithExpectedOneFunc[R]) WithQueryFunc[R] {
	return func(rows pgx.Rows) result.Result[R] {
		if rows.Next() {
			res := fn(rows)
			if res.HasFailed() {
				return result.Result[R]{}.Failed(res.UnwrapError())
			}
			return result.Success(res.Expect())
		}
		return result.Failed[R](errors.New("expected at least one row"))
	}
}

type WithOneFunc[R any] func(rows pgx.Rows) result.Result[R]

func WithOne[R any](fn WithOneFunc[R]) WithQueryFunc[option.Option[R]] {
	return func(rows pgx.Rows) result.Result[option.Option[R]] {
		if rows.Next() {
			res := fn(rows)
			if res.HasFailed() {
				return result.Result[option.Option[R]]{}.Failed(res.UnwrapError())
			}
			return result.Success(option.Some(res.Expect()))
		}
		return result.Success(option.None[R]())
	}
}

func Query(q pgx.Tx, ctx context.Context, sql string, args ...any) result.Result[pgx.Rows] {
	// Transform the args (Option -> nil/value)
	tfArgs := transformQueryArgs(args)
	if tfArgs.HasFailed() {
		return result.Failed[pgx.Rows](tfArgs.UnwrapError())
	}

	args = tfArgs.Expect()
	rows, err := q.Query(ctx, sql, args...)
	if err != nil {
		return result.Failed[pgx.Rows](err)
	}

	return result.Success(rows)
}

// Create a new transaction from the pg connection.
func NewTransaction(conn pgx.Conn, ctx context.Context) result.Result[transaction.PgTransaction] {
	pgTx, err := conn.Begin(context.Background())

	if err != nil {
		return result.Failed[transaction.PgTransaction](err)
	}

	return result.Success(transaction.PgTransaction{Inner: pgTx})
}

// Begin a transaction
func Begin(conn pgx.Conn, ctx context.Context, parent option.Option[transaction.Transaction], txName option.Option[string]) result.Result[transaction.PgTransaction] {
	if parent.IsSome() {
		return result.Map(parent.Expect().GetOrCreate(
			txName.UnwrapOr(func() string { return "pg" }),
			func() result.Result[transaction.ITransaction] {
				return result.Map(NewTransaction(conn, ctx), func(pgTx transaction.PgTransaction) transaction.ITransaction {
					pgTx.Managed = true
					return pgTx
				})
			},
		), func(t transaction.ITransaction) transaction.PgTransaction {
			return t.(transaction.PgTransaction)
		})
	}

	return NewTransaction(conn, ctx)
}
