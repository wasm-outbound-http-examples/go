package main

import (
	"fmt"
	"io"
	"net/http"

	_ "github.com/stealthrocket/net/http"
)

func main() {
	response, err := http.Get("http://httpbin.org/anything")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err == nil {
		fmt.Println(string(body))
	} else {
        fmt.Println(err)
    }
}
