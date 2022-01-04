package restlib

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type AuthSuccess struct {
	ID, Message string
}

func Get(url string) {

	client := resty.New()
	resp, err := client.R().Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Header().Get("server"))

}

func Post(url string) {

	data := map[string]string{
		"ID": "123",
	}

	_ = data
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").SetResult(&AuthSuccess{}).SetBody(data).SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").Post(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}

func CustomPost(url string) {

}
