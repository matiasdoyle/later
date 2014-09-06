package api

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/matiasdoyle/later/models"
)

func VerifyToken(params martini.Params, r *http.Request) (int, string) {
	token := params["token"]

	if token == "" {
		return http.StatusBadRequest, "Missing token"
	}

	u, err := models.FindUserByToken(token)
	if err != nil {
		return http.StatusInternalServerError, "Something bad happened"
	}

	if u == nil || u.Token == "" {
		return http.StatusUnauthorized, "Could not verify token"
	} else {
		return http.StatusOK, "Verified token"
	}
}
