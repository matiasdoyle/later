package models

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/coopernurse/gorp"
)

type User struct {
	Id       int64
	Email    string `form:"email"`
	Password string `form:"password"`
	Token    string
	Created  time.Time
	Updated  time.Time
}

func (u *User) PreInsert(s gorp.SqlExecutor) error {
	u.Created = time.Now()
	u.Updated = time.Now()

	if len(u.Token) == 0 {
		u.Token = createToken(u.Email)
	}

	return nil
}

func (u *User) PreUpdate(s gorp.SqlExecutor) error {
	u.Updated = time.Now()
	return nil
}

func CreateUser(u *User) (*User, error) {
	if EmailUsed(u.Email) {
		return nil, errors.New("Email is used")
	}

	p, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		panic(err)
	}

	u.Password = string(p)

	err = db.Insert(u)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return u, nil
}

func FindUserByToken(token string) (*User, error) {
	users := []User{}

	_, err := db.Select(&users, "SELECT * FROM users WHERE Token=?", token)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if len(users) != 1 {
		return nil, nil
	}

	return &users[0], nil
}

func EmailUsed(email string) bool {
	count, err := db.SelectInt("SELECT count(*) FROM users WHERE Email=?", email)
	if err != nil {
		panic(err)
	}

	return count > 0
}

func createToken(email string) string {
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%s%d", email, time.Now().Unix())))
	return hex.EncodeToString(h.Sum(nil))
}
