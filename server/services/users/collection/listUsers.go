package collection

import (
	"context"
	"strconv"

	users "github.com/alehechka/buf-connect-playground/proto/gen/users/v1"
	"github.com/alehechka/buf-connect-playground/utils/database"
	"github.com/brianvoe/gofakeit/v6"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func listItems(ctx context.Context, numUsers int32, userChan chan<- *users.User, errChan chan<- error) {
	num := int(numUsers)

	for i := 0; i < num; i++ {
		userChan <- GenerateUser(strconv.Itoa(gofakeit.Number(1, 9999)))
	}

	close(userChan)
	errChan <- database.EOD
	return
}

// ListUsers retrieves a list of Users and sends them via channel
func ListItems(ctx context.Context, numUsers int32, userChan chan<- *users.User, errChan chan<- error) {
	num := int64(numUsers)
	cursor, err := userCollection().Find(ctx, bson.M{}, &options.FindOptions{Limit: &num})
	if err != nil {
		close(userChan)
		errChan <- err
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		internalUser := user{}
		err = cursor.Decode(&internalUser)
		if err != nil {
			close(userChan)
			errChan <- err
			return
		}

		userChan <- internalUser.User()
	}

	close(userChan)
	errChan <- database.EOD
	return
}
