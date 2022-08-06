package users

import (
	context "context"
	"errors"
	"log"

	_type "github.com/alehechka/buf-connect-playground/proto/gen/google/type"
	v1 "github.com/alehechka/buf-connect-playground/proto/gen/users/v1"
	connect_go "github.com/bufbuild/connect-go"
)

func (s *server) GetUser(ctx context.Context, req *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.GetUserResponse], error) {
	userID := req.Msg.GetUserId()

	if len(userID) == 0 {
		return nil, connect_go.NewError(connect_go.CodeInvalidArgument, errors.New("no userID provided"))
	}

	log.Printf("Recieved request for user with ID: %s", userID)

	return connect_go.NewResponse(&v1.GetUserResponse{
		User: &v1.User{
			UserId:    userID,
			FirstName: "Adam",
			LastName:  "Lehechka",
			Gender:    v1.Gender_GENDER_MALE,
			Birthday: &_type.Date{
				Month: 3,
				Day:   7,
				Year:  1998,
			},
		},
	}), nil
}
