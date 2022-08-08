package users

import (
	context "context"
	"fmt"
	"strconv"

	users "github.com/alehechka/buf-connect-playground/proto/gen/users/v1"
	"github.com/brianvoe/gofakeit/v6"
	connect_go "github.com/bufbuild/connect-go"
)

func (s *server) GetUsers(ctx context.Context, req *connect_go.Request[users.GetUsersRequest], stream *connect_go.ServerStream[users.GetUserResponse]) error {
	numUsers := int(req.Msg.GetNumUsers())

	fmt.Printf("Got request for %d users\n", numUsers)

	for i := 0; i < numUsers; i++ {
		if err := stream.Send(&users.GetUserResponse{User: generateUser(strconv.Itoa(gofakeit.Number(1, 1000)))}); err != nil {
			return err
		}
	}

	return nil
}
