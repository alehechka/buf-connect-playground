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

func GetRemoteUser(user *users.User) (remoteUser RemoteUser, err error) {

	if len(user.GetUserId()) > 0 {
		id, err := primitive.ObjectIDFromHex(user.GetUserId())
		if err != nil {
			return RemoteUser{}, err
		}
		user.UserId = ""
		remoteUser.ID = id
	}

	remoteUser.User = user
	return
}

func GetNewRemoteUser(user *users.User) RemoteUser {
	user.UserId = ""
	return RemoteUser{User: user}
}
