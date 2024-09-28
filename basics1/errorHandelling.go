package basics1

import (
	"errors"
	"fmt"
)

type User struct {
	id    string
	email string
}

func getUser(id string) (User, error) {
	user := User{
		id:    id,
		email: "random@gmail.com",
	}

	if user.id == "1" {
		return user, errors.New("the user is an inappropriate user")
	}

	return user, nil
}

func errorHandelling() {
	user1, err := getUser("1")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user1)
}
