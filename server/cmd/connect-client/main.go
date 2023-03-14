package main

import (
	"context"
	"fmt"
	"net/http"

	"buf-connect-playground/cmd"
	"buf-connect-playground/utils"
	otel_shared "buf-connect-playground/utils/otel"

	usersv1 "buf-connect-playground/proto/gen/users/v1"

	"buf-connect-playground/proto/gen/users/v1/usersv1connect"

	connect_go "github.com/bufbuild/connect-go"
	otelconnect "github.com/bufbuild/connect-opentelemetry-go"
)

func main() {
	shutdownTracer, err := otel_shared.InitializeOpenTelTracer()
	utils.Check(err)
	defer shutdownTracer()

	client := usersv1connect.NewUsersServiceClient(http.DefaultClient, cmd.ConnectServerHost, connect_go.WithInterceptors(otelconnect.NewInterceptor()))
	res, err := client.GetUser(context.Background(), connect_go.NewRequest(&usersv1.GetUserRequest{UserId: cmd.UserID}))

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Msg.GetUser())
}
