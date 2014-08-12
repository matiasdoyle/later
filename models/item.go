package models

import (
	"time"

	"github.com/coopernurse/gorp"
)

type Item struct {
	Id          int64
	UserId      int64
	Title       string
	Url         string
	Description string
	State       int // Could be: 0 => unread/inbox, 1 => read, 2 => checked-out/archived
	Created     time.Time
	Updated     time.Time
}

func (i *Item) PreInsert(s gorp.SqlExecutor) error {
	i.Created = time.Now()
	i.Updated = time.Now()
	i.State = 0

	return nil
}

func (i *Item) PreUpdate(s gorp.SqlExecutor) error {
	i.Updated = time.Now()
	return nil
}

func CreateItem(i *Item, u User) (*Item, error) {
	i.UserId = u.Id

	err := db.Insert(i)

	if err != nil {
		return nil, err
	}

	return i, nil
}

func FindItems(u User) ([]Item, error) {
	items := []Item{}

	_, err := db.Select(&items, "SELECT * FROM items WHERE UserId = $1", u.Id)

	if err != nil {
		panic(err)
	}

	return items, nil
}
