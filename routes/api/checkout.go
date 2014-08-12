package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/matiasdoyle/checkout/models"
)

func CreateCheckoutItem(i models.Item, r *http.Request) (int, string) {
	qs := r.URL.Query()
	token := qs.Get("token")

	if token == "" {
		return 400, "Missing token"
	}

	u, err := models.FindUserByToken(token)
	if err != nil {
		return 500, "Oh no!"
	}

	if u == nil {
		return 401, "Unknown token"
	}

	if i.Url != "" && i.Title == "" {
		fmt.Println("No title")
		i.Title = GetHTMLTitle(i.Url)
	}

	fmt.Println(i)

	_, err = models.CreateItem(&i, *u)

	if err != nil {
		panic(err)
	}

	return 200, "Created checkout item"
}

// func IsValidUrl() {

// }

func GetHTMLTitle(url string) string {
	response, err := http.Get(url)
	r, _ := regexp.Compile("<title>(.+)</title>")

	if err != nil {
		panic(err)
	}

	if response.StatusCode != 200 {
		fmt.Println("Oh no!")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	matches := r.FindStringSubmatch(string(body))

	if len(matches) == 0 {
		return url
	}

	return matches[1]
}
