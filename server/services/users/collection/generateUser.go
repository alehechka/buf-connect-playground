package collection

import (
	_type "buf-connect-playground/proto/gen/google/type"

	users "buf-connect-playground/proto/gen/users/v1"

	"github.com/brianvoe/gofakeit/v6"
)

func GenerateUser(userID string) *users.User {
	user := GenerateNewUser()
	user.UserId = userID
	return user
}

func GenerateNewUser() *users.User {

	return &users.User{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Gender:    users.Gender(gofakeit.Number(1, 2)),
		Birthday: &_type.Date{
			Month: int32(gofakeit.Number(1, 12)),
			Day:   int32(gofakeit.Number(1, 28)),
			Year:  int32(gofakeit.Number(1940, 2022)),
		},
	}
}
