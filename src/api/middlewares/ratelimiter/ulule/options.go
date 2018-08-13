package echo

import (
	"net/http"

	"github.com/labstack/echo"
)

// Option is used to define Middleware configuration.
type Option interface {
	apply(*Middleware)
}

type option func(*Middleware)

func (o option) apply(middleware *Middleware) {
	o(middleware)
}

// ErrorHandler is an handler used to inform when an error has occurred.
type ErrorHandler func(c *echo.Context, err error)

// WithErrorHandler will configure the Middleware to use the given ErrorHandler.
func WithErrorHandler(handler ErrorHandler) Option {
	return option(func(middleware *Middleware) {
		middleware.OnError = handler
	})
}

// DefaultErrorHandler is the default ErrorHandler used by a new Middleware.
func DefaultErrorHandler(c *echo.Context, err error) {
	panic(err)
}

// LimitReachedHandler is an handler used to inform when the limit has exceeded.
type LimitReachedHandler func(c *echo.Context)

// WithLimitReachedHandler will configure the Middleware to use the given LimitReachedHandler.
func WithLimitReachedHandler(handler LimitReachedHandler) Option {
	return option(func(middleware *Middleware) {
		middleware.OnLimitReached = handler
	})
}

// DefaultLimitReachedHandler is the default LimitReachedHandler used by a new Middleware.
func DefaultLimitReachedHandler(c *echo.Context) {
	c.String(http.StatusTooManyRequests, "Limit exceeded")
}

// KeyGetter will define the rate limiter key given the echo Context
type KeyGetter func(c *echo.Context) string

// WithKeyGetter will configure the Middleware to use the given KeyGetter
func WithKeyGetter(KeyGetter KeyGetter) Option {
	return option(func(middleware *Middleware) {
		middleware.KeyGetter = KeyGetter
	})
}

// DefaultKeyGetter is the default KeyGetter used by a new Middleware
// It returns the Client IP address
func DefaultKeyGetter(c *echo.Context) string {
	return c.ClientIP()
}
