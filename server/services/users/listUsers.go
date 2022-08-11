package users

import (
	context "context"
	"errors"
	"fmt"

	users "buf-connect-playground/proto/gen/users/v1"

	"buf-connect-playground/services/users/collection"

	"buf-connect-playground/utils/database"

	connect_go "github.com/bufbuild/connect-go"
)

func (s *server) ListUsers(ctx context.Context, req *connect_go.Request[users.ListUsersRequest], stream *connect_go.ServerStream[users.ListUsersResponse]) error {
	numUsers := req.Msg.GetNumUsers()
	page := req.Msg.GetPage()
	fmt.Printf("Got request for %d users on page %d\n", numUsers, page)

	userChan := make(chan *users.User)
	errChan := make(chan error, 1)

	go collection.ListItems(ctx, numUsers, page, userChan, errChan)

	for user := range userChan {
		if err := stream.Send(&users.ListUsersResponse{User: user}); err != nil {
			errChan <- err
		}
	}

	err := <-errChan
	if err != nil && errors.Is(err, database.EOD) {
		return nil
	}
	return err
}
