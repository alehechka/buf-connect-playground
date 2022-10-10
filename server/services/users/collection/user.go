package collection

import (
	users "buf-connect-playground/proto/gen/users/v1"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RemoteUser struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	User *users.User `bson:",inline"`
}

func (r *RemoteUser) GetUser() *users.User {
	user := r.User
	user.UserId = r.ID.Hex()
	return user
}
