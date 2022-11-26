package main

import "fmt"

func main() {
	mynumber := 23
	var ptr *int
	ptr = &mynumber
	fmt.Println(mynumber)
	fmt.Println(ptr)
	fmt.Println(*ptr + 2)
}
