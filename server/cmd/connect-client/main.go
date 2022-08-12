package main

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"unsafe"

	"buf-connect-playground/cmd"
	"buf-connect-playground/utils"
	otel_shared "buf-connect-playground/utils/otel"

	usersv1 "buf-connect-playground/proto/gen/users/v1"

	"buf-connect-playground/proto/gen/users/v1/usersv1connect"

	"github.com/bufbuild/connect-go"
	connect_go "github.com/bufbuild/connect-go"
	"go.opentelemetry.io/otel"
)

func main() {
	shutdownTracer, err := otel_shared.InitializeOpenTelTracer()
	utils.Check(err)
	defer shutdownTracer()

	client := usersv1connect.NewUsersServiceClient(http.DefaultClient, cmd.ConnectServerHost, connect.WithInterceptors(&otelInterceptor{}))
	res, err := client.GetUser(context.Background(), connect_go.NewRequest(&usersv1.GetUserRequest{UserId: cmd.UserID}))

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Msg.GetUser())
}

type otelInterceptor struct{}

func (o *otelInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		ctx, span := otel.Tracer("buf-connect-playground/cmd/connect-client/otel").Start(ctx, req.Spec().Procedure)
		defer span.End()

		fmt.Printf("%#v\n", ctx.Value("val"))

		traceparent := fmt.Sprintf("00-%s-%s-01", span.SpanContext().TraceID().String(), span.SpanContext().SpanID().String())
		req.Header().Add("Traceparent", traceparent)
		return next(ctx, req)
	}
}

func (o *otelInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(ctx context.Context, spec connect.Spec) connect.StreamingClientConn {
		return &otelClientConn{
			StreamingClientConn: next(ctx, spec),
		}
	}
}

func (o *otelInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(ctx context.Context, conn connect.StreamingHandlerConn) error {
		return next(ctx, &otelHandlerConn{
			StreamingHandlerConn: conn,
		})
	}
}

type otelHandlerConn struct {
	connect.StreamingHandlerConn

	inspectedResponse     bool
	inspectResponseHeader func(connect.Spec, http.Header)
}

func (oh *otelHandlerConn) Send(msg any) error {
	return oh.StreamingHandlerConn.Send(msg)
}

type otelClientConn struct {
	connect.StreamingClientConn

	inspectedRequest      bool
	inspectRequestHeader  func(connect.Spec, http.Header)
	inspectedResponse     bool
	inspectResponseHeader func(connect.Spec, http.Header)
}

func (oc *otelClientConn) Send(msg any) error {
	return oc.StreamingClientConn.Send(msg)
}

func (oc *otelClientConn) Receive(msg any) error {
	return oc.StreamingClientConn.Receive(msg)
}

func printContextInternals(ctx interface{}, inner bool) {
	contextValues := reflect.ValueOf(ctx).Elem()
	contextKeys := reflect.TypeOf(ctx).Elem()

	if !inner {
		fmt.Printf("\nFields for %s.%s\n", contextKeys.PkgPath(), contextKeys.Name())
	}

	if contextKeys.Kind() == reflect.Struct {
		for i := 0; i < contextValues.NumField(); i++ {
			reflectValue := contextValues.Field(i)
			reflectValue = reflect.NewAt(reflectValue.Type(), unsafe.Pointer(reflectValue.UnsafeAddr())).Elem()

			reflectField := contextKeys.Field(i)

			if reflectField.Name == "Context" {
				printContextInternals(reflectValue.Interface(), true)
			} else {
				fmt.Printf("field name: %+v\n", reflectField.Name)
				fmt.Printf("value: %+v\n", reflectValue.Interface())
			}
		}
	} else {
		fmt.Printf("context is empty (int)\n")
	}
}
