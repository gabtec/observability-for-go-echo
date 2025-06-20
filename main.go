package main

import (
	"context"
	"gabtec/go-echo-obs-app/internal/handlers"
	mw "gabtec/go-echo-obs-app/internal/middleware"
	opentelemetry "gabtec/go-echo-obs-app/internal/openTelemetry"
	u "gabtec/go-echo-obs-app/internal/utils"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
)

func main() {
	// load env - do not fail on error because containers auto inject var without dotenv file
	godotenv.Load()

	addr := u.GetStringEnv("SERVER_ADDR", ":1323")
	serviceName := u.GetStringEnv("OTEL_SERVICE_NAME", "my-echo-log-app")
	otelURL := u.GetStringEnv("OTEL_ENDPOINT", "localhost:4316")

	// observability: traces
	ctx := context.Background()
	tp := opentelemetry.NewTraceProvider(ctx, otelURL, serviceName)
	defer func() { _ = tp.Shutdown(ctx) }()

	otel.SetTracerProvider(tp)

	// App
	e := echo.New()

	// Common middleware
	e.Use(middleware.RequestID())
	// e.Use(middleware.Logger())
	e.Use(mw.NewCustomLogger())
	e.Use(mw.NewPrometheusPerRequestMeter())
	// // Finally, set the tracer that can be used to create custom SPAN's, inside each route handler
	// handlers.SetCustomTracer(tp.Tracer(serviceName))
	e.Use(mw.NewTracerMiddleware(tp.Tracer(serviceName)))
	// // Or
	// // just use a middleware (2**)
	e.Use(otelecho.Middleware(serviceName)) // (2**)
	e.Use(middleware.Recover())

	// API Endpoints
	e.GET("/", handlers.IndexHandler)
	e.GET("/random", handlers.RandomHandler)
	e.GET("/log/:type", handlers.LogHandler)

	// observability: metrics
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	// e.Logger.Fatal(e.Start(":1323"))
	startServerWithGracefulShutdown(e, addr)
}

func startServerWithGracefulShutdown(e *echo.Echo, addr string) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(addr); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
