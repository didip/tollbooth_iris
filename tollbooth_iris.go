// Package tollbooth_iris provides rate-limiting logic to Iris request handler.
package tollbooth_iris

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/kataras/iris"
)

// LimitHandler is a middleware that performs
// rate-limiting given a "limiter" configuration.
func LimitHandler(lmt *limiter.Limiter) iris.Handler {
	return func(ctx iris.Context) {
		httpError := tollbooth.LimitByRequest(lmt, ctx.ResponseWriter(), ctx.Request())
		if httpError != nil {
			ctx.StatusCode(httpError.StatusCode)
			ctx.WriteString(httpError.Message)
			ctx.StopExecution()
			return
		}

		ctx.Next()
	}
}
