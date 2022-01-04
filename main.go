package main

import (
	"fmt"

	"github.com/maijaza/consumeapi/pkg/restlib"
	"github.com/maijaza/consumeapi/pkg/stdhttp"
)

func main() {
	fmt.Println("GET : ")
	stdhttp.Get("https://httpbin.org/get")
	fmt.Println("POST : ")
	stdhttp.Post("https://httpbin.org/post")
	fmt.Println("CUSTOM POST : ")
	stdhttp.CustomPost("https://httpbin.org/post")

	fmt.Println("================================================")
	fmt.Println("GET : ")
	restlib.Get("https://httpbin.org/get")
	fmt.Println("POST : ")
	restlib.Post("https://httpbin.org/post")
	fmt.Println("CUSTOM POST : ")
	restlib.CustomPost("https://httpbin.org/post")
}
