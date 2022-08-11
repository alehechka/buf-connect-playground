package collection

import (
	"context"

	users "buf-connect-playground/proto/gen/users/v1"

	"buf-connect-playground/utils/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ListUsers retrieves a list of Users and sends them via channel
func ListItems(ctx context.Context, numUsers int64, page int64, userChan chan<- *users.User, errChan chan<- error) {
	defer close(userChan)
	skip := numUsers * page

	cursor, err := userCollection().Find(ctx, bson.M{}, &options.FindOptions{Limit: &numUsers, Skip: &skip})
	if err != nil {
		errChan <- err
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		internalUser := user{}
		err = cursor.Decode(&internalUser)
		if err != nil {
			errChan <- err
			return
		}

		userChan <- internalUser.User()
	}

	errChan <- database.EOD
	return
}
