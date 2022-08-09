package main

import (
	"context"
	"fmt"

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
	user, err := client.GetUser(context.Background(), &usersv1.GetUserRequest{UserId: "62f2c57bd5061f85ee13f9b1"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(user)

}
