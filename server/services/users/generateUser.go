package users

import (
	_type "github.com/alehechka/buf-connect-playground/proto/gen/google/type"
	users "github.com/alehechka/buf-connect-playground/proto/gen/users/v1"
	"github.com/brianvoe/gofakeit/v6"
)

func generateUser(userID string) *users.User {
	return &users.User{
		UserId:    userID,
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Gender:    users.Gender(gofakeit.Number(1, 2)),
		Birthday: &_type.Date{
			Month: int32(gofakeit.Number(0, 12)),
			Day:   int32(gofakeit.Number(0, 28)),
			Year:  int32(gofakeit.Number(1940, 2022)),
		},
	}
}
