# tollbooth_iris

[Iris.v6](https://github.com/kataras/iris/tree/v6) middleware for rate limiting HTTP requests.

## Five Minutes Tutorial

```go
package main

import (
    "time"

    "github.com/didip/tollbooth"
    "github.com/didip/tollbooth_iris/v6"

    "gopkg.in/kataras/iris.v6"
    "gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
    app := iris.New()
    app.Adapt( httprouter.New() )

    // Create a limiter struct.
    limiter := tollbooth.NewLimiter(1, time.Second, nil)

    app.Get("/", tollbooth_iris.LimitHandler(limiter), func(ctx *iris.Context) {
        ctx.WriteString("Hello, world!")
    })

    app.Listen(":8080")
}
```
