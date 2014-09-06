package routes

import (
	"fmt"

	"github.com/matiasdoyle/later/models"
)

func Signup(u models.User) (int, string) {
	fmt.Println(u)

	_, err := models.CreateUser(&u)

	if err != nil {
		fmt.Println(err)
		return 400, err.Error()
	}

	return 200, "Created"
}
