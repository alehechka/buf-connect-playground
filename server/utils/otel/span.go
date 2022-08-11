package otel

import (
	"context"
	"runtime"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func StartSpan(ctx context.Context) (context.Context, trace.Span) {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)

	var callerName string
	if ok && details != nil {
		callerName = details.Name()
	}

	return otel.Tracer("grpc-method").Start(ctx, callerName)
}
