package stdhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// refer https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7

func Get(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}
	fmt.Println(string(body))
	fmt.Println(resp.Header.Get("server"))

}

func Post(url string) {

	data, _ := json.Marshal(map[string]string{
		"ID": "123",
	})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}
	fmt.Println(string(body))
}

func CustomPost(url string) {

	data, _ := json.Marshal(map[string]string{
		"ID": "123",
	})

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Add("Authorization", "Bearer aaa")

	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}
	fmt.Println(string(body))
}
