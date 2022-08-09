package users

import (
	context "context"
	"fmt"

	users "github.com/alehechka/buf-connect-playground/proto/gen/users/v1"
	"github.com/alehechka/buf-connect-playground/services/users/collection"
	connect_go "github.com/bufbuild/connect-go"
)

func (s *server) GetUser(ctx context.Context, req *connect_go.Request[users.GetUserRequest]) (*connect_go.Response[users.GetUserResponse], error) {
	userID := req.Msg.GetUserId()
	fmt.Printf("Got request to for user with ID: %s\n", userID)
	fmt.Println("Headers: ", req.Header())

	user, err := collection.GetUser(ctx, userID)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeNotFound, err)
	}

	res := connect_go.NewResponse(&users.GetUserResponse{
		User: user,
	})

	return res, nil
}
