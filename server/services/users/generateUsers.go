package users

import (
	context "context"
	"errors"
	"fmt"

	users "buf-connect-playground/proto/gen/users/v1"

	"buf-connect-playground/services/users/collection"

	connect_go "github.com/bufbuild/connect-go"
)

func (s *server) GenerateUsers(ctx context.Context, req *connect_go.Request[users.GenerateUsersRequest]) (*connect_go.Response[users.GenerateUsersResponse], error) {
	numUsers := int(req.Msg.GetNumUsers())
	fmt.Printf("Got request to create %d users\n", numUsers)

	if numUsers == 0 {
		return nil, connect_go.NewError(connect_go.CodeInvalidArgument, errors.New("must send numUsers greater than zero"))
	}

	genUsers := make([]*users.User, 0)
	for i := 0; i < numUsers; i++ {
		genUsers = append(genUsers, collection.GenerateNewUser())
	}

	numCreated, err := collection.BatchCreateUsers(ctx, genUsers)

	return connect_go.NewResponse(&users.GenerateUsersResponse{NumUsers: int32(numCreated)}), err
}
