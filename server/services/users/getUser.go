package users

import (
	context "context"
	"fmt"

	users "buf-connect-playground/proto/gen/users/v1"

	"buf-connect-playground/services/users/collection"

	"buf-connect-playground/utils/otel"

	connect_go "github.com/bufbuild/connect-go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func (s *server) GetUser(ctx context.Context, req *connect_go.Request[users.GetUserRequest]) (*connect_go.Response[users.GetUserResponse], error) {
	fmt.Printf("%#v\n", req.Header())
	spanCtx := trace.SpanContextFromContext(ctx)
	fmt.Println(spanCtx.TraceID())

	ctx, span := otel.StartSpan(ctx)
	defer span.End()

	userID := req.Msg.GetUserId()
	fmt.Printf("Got request to for user with ID: %s\n", userID)
	span.SetAttributes(attribute.String("attributes.userId", userID))

	user, err := collection.GetUser(ctx, userID)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeNotFound, err)
	}

	res := connect_go.NewResponse(&users.GetUserResponse{
		User: user,
	})

	return res, nil
}
