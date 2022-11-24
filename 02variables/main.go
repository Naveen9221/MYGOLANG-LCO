package main

import "fmt"

const Logintoken string = "naveen123" // First letter is capital so it is reffered as public variable

func main() {
	var user string = "naveen"
	fmt.Println(user)
	var user1 string
	fmt.Println(user1)

	var isuser bool = true
	fmt.Println(isuser)
	fmt.Printf("type : %T ", isuser)

	var a = 10 // without type declaration lexer assign the type
	fmt.Println(a)

	b := 10 // without var but not allowed to use outside the method
	fmt.Println(b)

	fmt.Println(Logintoken)

}
