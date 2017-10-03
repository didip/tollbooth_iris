// Package tollbooth_iris provides rate-limiting logic to Iris (current version 8.x) request handler.
package tollbooth_iris

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/kataras/iris/context"
)

// LimitHandler is a middleware that performs
// rate-limiting given a "limiter" configuration.
//
// Read more at: https://github.com/didip/tollbooth
// And https://github.com/didip/tollbooth_iris
func LimitHandler(l *limiter.Limiter) context.Handler {
	return func(ctx context.Context) {
		httpError := tollbooth.LimitByRequest(l, ctx.ResponseWriter(), ctx.Request())
		if httpError != nil {
			ctx.StatusCode(httpError.StatusCode)
			ctx.WriteString(httpError.Message)
			ctx.StopExecution()
			return
		}

		ctx.Next()
	}
}
