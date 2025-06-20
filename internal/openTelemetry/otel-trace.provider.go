package opentelemetry

import (
	"context"
	"gabtec/go-echo-obs-app/version"
	"log/slog"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func createExporter(ctx context.Context, url string) sdktrace.SpanExporter {
	exp, err := otlptracehttp.New(ctx,
		otlptracehttp.WithInsecure(),    // demo – plaintext
		otlptracehttp.WithEndpoint(url), // points to collector service name in compose
	)
	if err != nil {
		slog.Error(err.Error())
	}

	return exp
}

func createResource(serviceName string) *resource.Resource {
	// create one OR merge many
	res, error := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(version.Version()),
		),
	)
	if error != nil {
		slog.Warn(error.Error())
	}

	return res
}

func NewTraceProvider(ctx context.Context, url, serviceName string) *sdktrace.TracerProvider {
	exporter := createExporter(ctx, url)
	resource := createResource(serviceName)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource),
	)

	return tp
}
