package users

import (
	"buf-connect-playground/proto/gen/users/v1/usersv1connect"
)

type server struct {
	usersv1connect.UnimplementedUsersServiceHandler
}

func NewServer() usersv1connect.UsersServiceHandler {
	return &server{}
}
