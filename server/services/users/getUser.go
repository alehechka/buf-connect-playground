package users

import (
	context "context"
	"errors"
	"fmt"
	"log"
	"net/http"

	users "github.com/alehechka/buf-connect-playground/proto/gen/users/v1"
	"github.com/alehechka/buf-connect-playground/utils/grpc"
	connect_go "github.com/bufbuild/connect-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func (s *server) GetUser(ctx context.Context, req *connect_go.Request[users.GetUserRequest]) (*connect_go.Response[users.GetUserResponse], error) {
	ctx, span := otel.Tracer("grpc-server").Start(ctx, "GetUser")
	defer span.End()

	userID := req.Msg.GetUserId()
	span.SetAttributes(attribute.String("userID", userID))

	if len(userID) == 0 {
		return nil, connect_go.NewError(connect_go.CodeInvalidArgument, errors.New("no userID provided"))
	}

	log.Printf("Recieved request for user with ID: %s", userID)

	fmt.Printf("Cookies: \n\t%s\n\t%s\n", grpc.GetCookie(req, "SessionID"), grpc.GetCookie(req, "ContextID"))

	res := connect_go.NewResponse(&users.GetUserResponse{
		User: generateUser(userID),
	})

	grpc.AddCookie(res, &http.Cookie{Name: "SessionID", Value: "bababooie", HttpOnly: true})
	grpc.AddCookie(res, &http.Cookie{Name: "ContextID", Value: "booboobaie", HttpOnly: true})

	return res, nil
}
