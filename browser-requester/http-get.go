package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Nigel2392/requester"
)

func main() {
	var client = requester.NewAPIClient()
	ch := make(chan struct{})
	client.Get("https://httpbin.org/anything").Do(func(response *http.Response) {
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err == nil {
			fmt.Println(string(body))
		} else {
			fmt.Println(err)
		}

		ch <- struct{}{}
	})
	<-ch
}
