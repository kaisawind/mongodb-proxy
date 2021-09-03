# negroni-logrus

[![PkgGoDev](https://pkg.go.dev/badge/github.com/meatballhat/negroni-logrus)](https://pkg.go.dev/github.com/meatballhat/negroni-logrus)
[![Main Workflow Status](https://github.com/meatballhat/negroni-logrus/workflows/main/badge.svg)](https://github.com/meatballhat/negroni-logrus/workflows/main)

logrus middleware for negroni

## Usage

Take a peek at the [basic example](./examples/basic/example.go) or the [custom
example](./examples/custom/example.go), the latter of which includes setting a
custom log level on the logger with `NewCustomMiddleware`.

If you want to reuse an already initialized `logrus.Logger`, you should be using
`NewMiddlewareFromLogger` (see [existinglogrus](./examples/existinglogrus/example.go)).
