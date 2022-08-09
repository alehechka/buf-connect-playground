package collection

import (
	"context"

	users "github.com/alehechka/buf-connect-playground/proto/gen/users/v1"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(ctx context.Context, user *users.User) (id primitive.ObjectID, err error) {
	createUser := newUser(user)

	res, err := userCollection().InsertOne(ctx, createUser)
	if err != nil {
		return
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func BatchCreateUsers(ctx context.Context, users []*users.User) (inserted int, err error) {

	operations := make([]mongo.WriteModel, 0)
	for _, user := range users {
		operations = append(operations, &mongo.InsertOneModel{
			Document: newUser(user),
		})
	}

	res, err := userCollection().BulkWrite(ctx, operations)

	return int(res.InsertedCount), err
}