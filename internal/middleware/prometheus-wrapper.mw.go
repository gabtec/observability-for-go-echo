package middleware

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response duration for handler.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method", "status"},
	)
)

func NewPrometheusPerRequestMeter() echo.MiddlewareFunc {
	prometheus.MustRegister(httpDuration)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			status := c.Response().Status
			duration := time.Since(start).Seconds()

			httpDuration.WithLabelValues(
				c.Path(), // echo route pattern, not raw path
				c.Request().Method,
				http.StatusText(status),
			).Observe(duration)

			return err
		}
	}
}
