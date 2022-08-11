package main

import (
	"context"
	"fmt"

	"buf-connect-playground/cmd"
	"buf-connect-playground/utils"
	"buf-connect-playground/utils/otel"

	usersv1 "buf-connect-playground/proto/gen/users/v1"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func main() {
	shutdownTracer, err := otel.InitializeOpenTelTracer()
	utils.Check(err)
	defer shutdownTracer()

	conn, err := grpc.Dial(cmd.GrpcServerHost, ClientDialOptions()...)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client := usersv1.NewUsersServiceClient(conn)
	user, err := client.GetUser(context.Background(), &usersv1.GetUserRequest{UserId: cmd.UserID})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(user)

}

func ClientDialOptions() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor(otelgrpc.WithTracerProvider(otel.OpenTelTracer))),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor(otelgrpc.WithTracerProvider(otel.OpenTelTracer))),
	}
}
