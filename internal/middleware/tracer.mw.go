package middleware

import (
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/trace"
)

// we can put tracer inside echo.Context or we can use a singleton
// here we go with context approach
func NewTracerMiddleware(tracer trace.Tracer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("tracer", tracer)
			return next(c)
		}
	}
}
