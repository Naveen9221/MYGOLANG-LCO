package main

import "fmt"

func main() {

	//days := []string{"m", "t", "wed"}

	///for i := 0; i < len(days); i++ {
	//	fmt.Println(days[i])
	//}

	//for j := range days {
	//	fmt.Println(days[j])
	//}

	//for index, day := range days {
	//	fmt.Printf("the index is %v and the value is %v", index, day)
	//}

	k := 2
	for k < 5 {
		fmt.Println(k)
		k++
	}
	if k == 5 {
		goto lco
	}

lco:
	fmt.Println("go to definition")
}
