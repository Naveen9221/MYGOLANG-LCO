package main

import "fmt"

func main() {
	fmt.Println("My function the excecution starts here")
	result := adder(2, 5)
	proresult, mess := proadder(12, 4, 5, 6, 7) //it act like slice
	//dont care about the output mention like proresult,_:=proadder(12, 4, 5, 6, 7)
	fmt.Printf("The result is :%v", result)
	fmt.Printf("\n pro result:%v and %v", proresult, mess)

}
func adder(v1 int, v2 int) int {
	return v1 + v2
}
func proadder(values ...int) (int, string) { //need to mention type of values getting and type of values the function goting to return from the
	total := 0
	for val := range values {
		total += val
	}
	return total, "The net total"

}
