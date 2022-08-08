package users

import (
	context "context"
	"errors"
	"fmt"

	users "github.com/alehechka/buf-connect-playground/proto/gen/users/v1"
	"github.com/alehechka/buf-connect-playground/services/users/collection"
	"github.com/alehechka/buf-connect-playground/utils/database"
	connect_go "github.com/bufbuild/connect-go"
)

func (s *server) ListUsers(ctx context.Context, req *connect_go.Request[users.ListUsersRequest], stream *connect_go.ServerStream[users.GetUserResponse]) error {
	numUsers := req.Msg.GetNumUsers()
	fmt.Printf("Got request for %d users\n", numUsers)

	userChan := make(chan *users.User)
	errChan := make(chan error, 1)

	go collection.ListItems(ctx, numUsers, userChan, errChan)

	go func() {
		for user := range userChan {
			if err := stream.Send(&users.GetUserResponse{User: user}); err != nil {
				errChan <- err
			}
		}
	}()

	err := <-errChan
	if err != nil && errors.Is(err, database.EOD) {
		return nil
	}
	return err
}
