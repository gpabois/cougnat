package middlewares

import (
	"context"
	"net/http"

	"github.com/gpabois/cougnat/core/iter"
	"github.com/gpabois/cougnat/core/result"
)

type Middleware interface {
	Handle(ctx context.Context, r *http.Request) result.Result[context.Context]
}

type Middlewares []Middleware

func (middlewares *Middlewares) Iter() iter.Iterator[Middleware] {
	return iter.IterSlice(middlewares)
}

func (middlewares *Middlewares) Handle(ctx context.Context, r *http.Request) result.Result[context.Context] {
	for _, middleware := range *middlewares {
		res := middleware.Handle(ctx, r)
		if res.HasFailed() {
			return result.Result[context.Context]{}.Failed(res.UnwrapError())
		}
		ctx = res.Expect()
	}

	return result.Success(ctx)
}
