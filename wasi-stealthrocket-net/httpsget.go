package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	_ "github.com/stealthrocket/net/http"
	"github.com/stealthrocket/net/wasip1"
)

func main() {
    // No certAuthority certificates inside WASM module.
    // So just switching TLS certificate verification off.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext:     wasip1.DialContext,
	}
	client := &http.Client{Transport: tr}

	response, err := client.Get("https://httpbin.org/anything")
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
