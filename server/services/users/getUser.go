package users

import (
	context "context"
	"errors"
	"fmt"
	"log"

	users "github.com/alehechka/buf-connect-playground/proto/gen/users/v1"
	connect_go "github.com/bufbuild/connect-go"
)

func (s *server) GetUser(ctx context.Context, req *connect_go.Request[users.GetUserRequest]) (*connect_go.Response[users.GetUserResponse], error) {
	userID := req.Msg.GetUserId()

	if len(userID) == 0 {
		return nil, connect_go.NewError(connect_go.CodeInvalidArgument, errors.New("no userID provided"))
	}

	log.Printf("Recieved request for user with ID: %s", userID)

	fmt.Println("Request Headers: ", req.Header())
	fmt.Printf("Request Spec: %#v\n", req.Spec())

	res := connect_go.NewResponse(&users.GetUserResponse{
		User: generateUser(userID),
	})

	fmt.Println("Response Headers: ", res.Header())
	fmt.Println("Response Trailer: ", res.Trailer())

	return res, nil
}
