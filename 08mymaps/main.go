package main

import "fmt"

func main() {
	fmt.Println("my maps")
	lan := make(map[string]string)
	lan["ja"] = "javascript"
	lan["go"] = "go lang"
	lan["ls"] = "lang"
	lan["ru"] = "ruby"
	fmt.Println(lan["ja"])
	delete(lan, "ls")
	fmt.Println(lan)
	for key, value := range lan {
		fmt.Printf("for the key %v: value is %v \n", key, value)
	}
}
