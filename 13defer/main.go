package main

import "fmt"

func main() {
	defer fmt.Println("world") //if you use defer it works like a stack
	fmt.Println("hello")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
