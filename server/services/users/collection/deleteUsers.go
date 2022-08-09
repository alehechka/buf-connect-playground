package collection

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func DeleteAllUsers(ctx context.Context) (deleted int, err error) {
	res, err := userCollection().DeleteMany(ctx, bson.M{})

	return int(res.DeletedCount), err
}
