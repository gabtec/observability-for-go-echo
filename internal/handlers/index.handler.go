package handlers

import (
	u "gabtec/go-echo-obs-app/internal/utils"
	"gabtec/go-echo-obs-app/version"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func IndexHandler(c echo.Context) error {

	// add a custom span
	customTrace, ok := c.Get("tracer").(trace.Tracer)
	if ok && customTrace != nil {
		_, span := customTrace.Start(c.Request().Context(), "manual-span-v3")
		defer span.End()
		span.SetName("manual-span-named") // will override "manual-span-v3"
		span.SetAttributes(attribute.String("key1", "value2"))
		span.SetAttributes(attribute.String("env", "dev"))
	}
	// --end

	endpoints := map[string]string{
		"/":          "this page",
		"/log/:type": "returns a success (type=ok), or error (type=error), response",
		"/random":    "returns a random response",
		"/demo":      "same as /random (for back compatibility)",
	}

	resp := map[string]interface{}{
		"endpointsList": endpoints,
		"version":       version.Version(),
		"observability": "implemented",
	}

	return u.JSONOK(c, resp)
}
