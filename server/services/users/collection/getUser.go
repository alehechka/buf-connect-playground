package collection

import (
	"context"

	users "github.com/alehechka/buf-connect-playground/proto/gen/users/v1"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetUser attempts to retrieve a single user denoted the provided id.
func GetUser(ctx context.Context, id string) (*users.User, error) {
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
