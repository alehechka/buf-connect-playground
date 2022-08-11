package collection

import (
	"buf-connect-playground/utils/database"

	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func userCollection() *mongo.Collection {
	if collection == nil && database.Database != nil {
		collection = database.Database.Collection("users")
	}

	return collection
}
