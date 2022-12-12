package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	databyte, err := ioutil.ReadAll(response.Body)

	content := string(databyte)

	fmt.Println(content)

	fmt.Printf("type is:%T", response)
	defer response.Body.Close()

}
