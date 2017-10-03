# tollbooth_iris

[Iris](https://github.com/kataras/iris) middleware for rate limiting HTTP requests.

## Prerequisites

1. Install the [Iris web framework](https://github.com/kataras/iris); open a terminal and execute `go get -u github.com/kataras/iris`
2. Install the tollbooth package; open a terminal and execute `go get -u github.com/didip/tollbooth`
3. Install the tollbooth middleware for Iris; open a terminal and execute `go get -u github.com/didip/tollbooth_iris`

## Five Minutes Tutorial

```go
// main.go

package main

import (
    "time"

    "github.com/kataras/iris"

    "github.com/didip/tollbooth"
    "github.com/didip/tollbooth_iris"
)

func main() {
    app := iris.New()

    // Create a limiter struct.
    limiter := tollbooth.NewLimiter(1, time.Second, nil)

    app.Get("/", tollboothic.LimitHandler(limiter), func(ctx iris.Context) {
        ctx.HTML("<b>Hello, world!</b>")
    })

    app.Run(iris.Addr(":8080"))
}
```

```bash
$ go run main.go
Now listening on: http://localhost:8080
Application started. Press CTRL+C to shut down
```

> Keep note that you'll need to navigate and use the [tollbooth_iris v6](v6) if you're using the older Iris version 6.