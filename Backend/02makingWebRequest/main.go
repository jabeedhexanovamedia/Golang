package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Making a web request...")

	url:= "https://jsonplaceholder.typicode.com/todos/4"

	res,err := http.Get(url)

	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	defer res.Body.Close()

	// Read the response body

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Println("Response : ",string(data))
}