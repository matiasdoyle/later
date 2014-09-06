package routes

import (
	"github.com/martini-contrib/render"
	"github.com/matiasdoyle/later/models"
)

type SignupResponse struct {
	Error string `json:"error"`
	Token string `json:"token"`
}

func Signup(u models.User, r render.Render) {
	_, err := models.CreateUser(&u)
	res := SignupResponse{}
	statusCode := 200

	if err != nil {
		res.Error = err.Error()
		statusCode = 400
	}

	if u.Token != "" {
		res.Token = u.Token
	}

	r.JSON(statusCode, res)
}
