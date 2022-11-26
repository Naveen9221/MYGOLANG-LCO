package main

import "fmt"

func main() {
	mynumber := 23
	var ptr *int
	ptr = &mynumber
	fmt.Println(mynumber)
	fmt.Println(ptr)
	fmt.Println(*ptr + 2)

	//arrays

	var list [4]string // or  var list=[4] string{"123","345","56","5"}
	list[0] = "123"
	list[3] = "456"
	fmt.Println(list, "length", len(list))
}
