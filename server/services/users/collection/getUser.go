package collection

import (
	"context"

	users "buf-connect-playground/proto/gen/users/v1"

	"buf-connect-playground/utils/otel"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetUser attempts to retrieve a single user denoted the provided id.
func GetUser(ctx context.Context, id string) (*users.User, error) {
	ctx, span := otel.StartSpan(ctx)
	defer span.End()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res := userCollection().FindOne(ctx, user{ID: oid})
	if res.Err() != nil {
		return nil, res.Err()
	}

	internalUser := user{}
	err = res.Decode(&internalUser)
	return internalUser.User(), err
}
