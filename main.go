package main

import (
	"context"
	"gabtec/go-echo-obs-app/internal/handlers"
	mw "gabtec/go-echo-obs-app/internal/middleware"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Common middleware
	e.Use(middleware.RequestID())
	// e.Use(middleware.Logger())
	e.Use(mw.NewCustomLogger())
	e.Use(middleware.Recover())

	// API Endpoints
	e.GET("/", handlers.IndexHandler)
	e.GET("/random", handlers.RandomHandler)
	e.GET("/log/:type", handlers.LogHandler)

	// e.Logger.Fatal(e.Start(":1323"))
	startServerWithGracefulShutdown(e, ":1323")
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
