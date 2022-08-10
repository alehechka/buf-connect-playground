package main

import (
	"context"
	"fmt"

	"github.com/alehechka/buf-connect-playground/cmd"
	usersv1 "github.com/alehechka/buf-connect-playground/proto/gen/users/v1"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:3000", grpc.WithInsecure())
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
