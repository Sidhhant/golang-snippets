package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(string(body))
}
