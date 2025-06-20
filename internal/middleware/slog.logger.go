package middleware

import (
	"context"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// -----------------------------------------------------------------------------
// CustomLogger is an Echo middleware that writes each request as structured JSON
// using the Go 1.21 slog package.  Import the package and call:
//
//	e.Use(mw.CustomLogger)
//
// -----------------------------------------------------------------------------
var customLogger echo.MiddlewareFunc

func init() {
	// 1. Create a JSON slog logger that writes to stdout.
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// 2. Configure Echo’s RequestLogger.
	cfg := middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		LogMethod:   true,
		HandleError: true, // keep error propagation
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("method", v.Method),
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("method", v.Method),
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}

	// 3. Build the Echo middleware and expose it as a package‑level variable.
	customLogger = middleware.RequestLoggerWithConfig(cfg)
}

func NewCustomLogger() echo.MiddlewareFunc {
	return customLogger
}
