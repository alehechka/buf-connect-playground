package middleware

import (
	"net/http"

	"buf-connect-playground/utils/otel"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func AttachOpenTelemetry(next http.Handler, operation string) http.Handler {
	return otelhttp.NewHandler(next, operation, otelhttp.WithTracerProvider(otel.OpenTelTracer))
}
