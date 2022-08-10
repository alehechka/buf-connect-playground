package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/alehechka/buf-connect-playground/cmd"
	usersv1 "github.com/alehechka/buf-connect-playground/proto/gen/users/v1"
	"github.com/alehechka/buf-connect-playground/proto/gen/users/v1/usersv1connect"
	connect_go "github.com/bufbuild/connect-go"
)

func main() {
	client := usersv1connect.NewUsersServiceClient(http.DefaultClient, "http://localhost:3000/api")
	res, err := client.GetUser(context.Background(), connect_go.NewRequest(&usersv1.GetUserRequest{UserId: cmd.UserID}))

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Msg.GetUser())
}
