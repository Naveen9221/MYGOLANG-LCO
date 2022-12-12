package main

import (
	"fmt"
	//"net/http"
	"net/url"
)

const myurl string = "http://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456jhb"

func main() {
	fmt.Println(myurl)

	result, err := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	if err != nil {
		panic(err)
	}

	qpar := result.Query()
	fmt.Printf("type is:%T \n", qpar)
	fmt.Println(qpar["coursename"])
	//for val, key := range qpar {
	fmt.Printf("values %+v ", qpar)
	//}

	partsofurl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	//anotherurl := partsofurl.String()
	fmt.Println(partsofurl)

}
