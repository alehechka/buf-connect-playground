package users

import (
	context "context"
	"fmt"

	users "github.com/alehechka/buf-connect-playground/proto/gen/users/v1"
	"github.com/alehechka/buf-connect-playground/services/users/collection"
	connect_go "github.com/bufbuild/connect-go"
)

func (s *server) GenerateUsers(ctx context.Context, req *connect_go.Request[users.GenerateUsersRequest]) (*connect_go.Response[users.GenerateUsersResponse], error) {
	numUsers := int(req.Msg.GetNumUsers())
	fmt.Printf("Got request to create %d users\n", numUsers)

	genUsers := make([]*users.User, 0)
	for i := 0; i < numUsers; i++ {
		genUsers = append(genUsers, collection.GenerateNewUser())
	}

	numCreated, err := collection.BatchCreateUsers(ctx, genUsers)

	return connect_go.NewResponse(&users.GenerateUsersResponse{NumUsers: int32(numCreated)}), err
}
