package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type read_resp struct{}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// x := make([]byte, 99)
	// fmt.Println(resp.Body.Read(x))
	// fmt.Println(string(x))
	io.Copy(os.Stdout, resp.Body)
}
