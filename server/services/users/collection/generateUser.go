package collection

import (
	_type "github.com/alehechka/buf-connect-playground/proto/gen/google/type"
	users "github.com/alehechka/buf-connect-playground/proto/gen/users/v1"
	"github.com/brianvoe/gofakeit/v6"
)

func GenerateUser(userID string) *users.User {
	user := GenerateNewUser()
	user.UserId = userID
	return user
}

func GenerateNewUser() *users.User {
	birthday := gofakeit.Date()

	return &users.User{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Gender:    users.Gender(gofakeit.Number(1, 2)),
		Birthday: &_type.Date{
			Month: int32(birthday.Month()),
			Day:   int32(birthday.Day()),
			Year:  int32(birthday.Year()),
		},
	}
}
