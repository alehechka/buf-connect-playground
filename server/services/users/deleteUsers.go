package users

import (
	context "context"
	"fmt"

	users "buf-connect-playground/proto/gen/users/v1"

	"buf-connect-playground/services/users/collection"

	connect_go "github.com/bufbuild/connect-go"
)

func (s *server) DeleteAllUsers(ctx context.Context, req *connect_go.Request[users.DeleteAllUsersRequest]) (*connect_go.Response[users.DeleteAllUsersResponse], error) {
	fmt.Println("Got request to delete all users")

	deleted, err := collection.DeleteAllUsers(ctx)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}

	return connect_go.NewResponse(&users.DeleteAllUsersResponse{NumUsers: int32(deleted)}), nil
}
