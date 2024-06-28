package main

import (
	"fmt"
	"io"
	"net/http"
)

func readURL(url string) error {
	response, err := http.Get(url)
	if err != nil {
		// fmt.Println("ERROR: ", err)
		return err
	}
	defer response.Body.Close()
	body, errors := io.ReadAll(response.Body)
	if err != nil {
		// fmt.Println("ERROR: ", errors)
		return errors
	}
	fmt.Println(string(body))

	return nil

}

func main() {
	var url string
	// url = "https://jsonplaceholder.typicode.com/posts"
	url = "https://googl.com"
	if err := readURL(url); err != nil {
		fmt.Println("ERROR", err)
	}
}
