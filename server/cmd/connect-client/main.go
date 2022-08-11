package main

import (
	"context"
	"fmt"
	"net/http"

	"buf-connect-playground/cmd"

	usersv1 "buf-connect-playground/proto/gen/users/v1"

	"buf-connect-playground/proto/gen/users/v1/usersv1connect"

	connect_go "github.com/bufbuild/connect-go"
)

func main() {
	client := usersv1connect.NewUsersServiceClient(http.DefaultClient, cmd.ConnectServerHost)
	res, err := client.GetUser(context.Background(), connect_go.NewRequest(&usersv1.GetUserRequest{UserId: cmd.UserID}))

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Msg.GetUser())
}
