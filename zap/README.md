# go-problem/zap

[![Go Reference](https://img.shields.io/badge/go.dev-reference-007d9c?style=for-the-badge&logo=go&logoColor=white)](https://pkg.go.dev/github.com/neocotic/go-problem/zap)
[![Build Status](https://img.shields.io/github/actions/workflow/status/neocotic/go-problem/ci.yml?style=for-the-badge)](https://github.com/neocotic/go-problem/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/neocotic/go-problem?style=for-the-badge)](https://github.com/neocotic/go-problem)
[![License](https://img.shields.io/github/license/neocotic/go-problem?style=for-the-badge)](https://github.com/neocotic/go-problem/blob/main/LICENSE.md)

Supports seamless logging of problems via [zap](https://github.com/uber-go/zap).

## Installation

Install using [go get](https://go.dev/ref/mod#go-get):

``` sh
go get github.com/neocotic/go-problem github.com/neocotic/go-problem/zap go.uber.org/zap
```

Then import the package into your own code:

``` go
import (
    "github.com/neocotic/go-problem"
    "github.com/neocotic/go-problem/zap"
    "go.uber.org/zap"
)
```

## Documentation

Documentation is available on [pkg.go.dev](https://pkg.go.dev/github.com/neocotic/go-problem/zap#section-documentation).
It contains an overview and reference.

### Example

Use the global `zap.Logger` to log problems:

``` go
problem.DefaultGenerator.Logger = problemzap.GlobalLogger()
```

Use a specific `zap.Logger` to log problems:

``` go
problem.DefaultGenerator.Logger = problemzap.LoggerFrom(zap.Must(zap.NewDevelopment()))
```

Populate fields on the `zap.Logger` from a `context.Context`:

``` go
problem.DefaultGenerator.Logger = problemzap.LoggerFromContext(zap.Must(zap.NewDevelopment()), func(ctx context.Context, logger *zap.Logger) *zap.Logger {
    return logger.With(zap.Any("correlationId", ctx.Value("correlationId")))
})
```

The above examples use `problem.DefaultGenerator` for brevity, but you can also assign the `problem.Logger` to custom
`problem.Generator` if you prefer.

Finally, if you never plan of logging directly via the `problem` package but are still using `zap`, you can just use
`problemzap.Field`, `problemzap.FieldUsing`, or `problemzap.NamedField` to get a structured `zap.Field` to represent a
problem:

``` go
prob := problem.Build().
    Title(http.StatusText(http.StatusNotFound)).
    Status(http.StatusNotFound).
	Code(problem.MustBuildCode(404, "USER")).
	Detail("User not found").
	Instance("https://api.example.void/users/123").
	Problem()
// ...
zap.L().Error("Failed to get user", problemzap.NamedField("problem", prob))
```

However, it's still recommended setting up
`problem.DefaultGenerator.Logger` to integrate with `zap` properly to avoid missing logs for problems in your
application.

## Issues

If you have any problems or would like to see changes currently in development you can do so
[here](https://github.com/neocotic/go-problem/issues).

## Contributors

If you want to contribute, you're a legend! Information on how you can do so can be found in
[CONTRIBUTING.md](https://github.com/neocotic/go-problem/blob/main/CONTRIBUTING.md). We want your suggestions and pull
requests!

A list of contributors can be found in [AUTHORS.md](https://github.com/neocotic/go-problem/blob/main/AUTHORS.md).

## License

Copyright © 2025 neocotic

See [LICENSE.md](https://github.com/neocotic/go-problem/raw/main/LICENSE.md) for more information on our MIT license.
