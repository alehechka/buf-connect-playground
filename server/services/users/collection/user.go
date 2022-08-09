package collection

import (
	_type "github.com/alehechka/buf-connect-playground/proto/gen/google/type"
	users "github.com/alehechka/buf-connect-playground/proto/gen/users/v1"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type user struct {
	ID primitive.ObjectID `bson:"_id"`

	FirstName string       `bson:"f,omitempty"`
	LastName  string       `bson:"l,omitempty"`
	Gender    users.Gender `bson:"g,omitempty"`
	Birthday  *_type.Date  `bson:"b,omitempty"`
}

func newUser(u *users.User) user {
	return updateUser(u, primitive.NewObjectID())
}

func updateUser(u *users.User, id primitive.ObjectID) user {
	return user{
		ID:        id,
		FirstName: u.GetFirstName(),
		LastName:  u.GetLastName(),
		Gender:    u.GetGender(),
		Birthday:  u.GetBirthday(),
	}
}

func (u *user) User() *users.User {
	return &users.User{
		UserId:    u.ID.Hex(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Gender:    u.Gender,
		Birthday:  u.Birthday,
	}
}
