package otel

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var GinTracer = otel.Tracer("gin-server")
var GrpcTracer = otel.Tracer("grpc-server")

var OpenTelTracer *sdktrace.TracerProvider

func InitializeOpenTelTracer() (disconnect func() error, err error) {
	ctx := context.Background()
	exporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		return nil, err
	}

	OpenTelTracer = sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(OpenTelTracer)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return func() error {
		return OpenTelTracer.Shutdown(ctx)
	}, nil
}
