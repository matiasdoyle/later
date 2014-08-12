package routes

import (
	"fmt"

	"github.com/martini-contrib/render"

	"github.com/matiasdoyle/checkout/models"
)

func RenderInbox(r render.Render) {
	u := models.User{Id: 1}

	items, err := models.FindItems(u)

	if err != nil {
		panic(err)
	}

	fmt.Println(items)

	r.HTML(200, "inbox", items)
}
