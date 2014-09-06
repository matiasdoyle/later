package routes

import (
	"fmt"

	"github.com/martini-contrib/render"

	"github.com/matiasdoyle/later/models"
)

func RenderInbox(r render.Render) {
	u := models.FindUserById(1)
	if u == nil {
		r.HTML(200, "inbox", []models.Item{})
		return
	}

	items, err := models.FindItems(*u)
	if err != nil {
		panic(err)
	}

	fmt.Println(items)

	r.HTML(200, "inbox", items)
}
