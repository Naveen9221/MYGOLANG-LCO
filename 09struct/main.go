package main

import "fmt"

// no inheritance in go lang and capital letter for access modifier
func main() {
	user1 := user{"naveen", "naveenp.19ece@kk", 21}
	fmt.Printf("%+v", user1) // +v is to display like {Name:naveen Email:naveenp.19ece@kk Age:21}
	fmt.Printf("%v", user1.Name)

	login := 23
	if login < 10 {
		fmt.Println("regular")
	} else {
		fmt.Println("need to check")
	}

	if n := 1; n == 1 {
		fmt.Println("variable is initialized and used in the same place")
	}
}

type user struct {
	Name  string
	Email string
	Age   int
}
