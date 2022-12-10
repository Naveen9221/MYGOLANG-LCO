package main

import (
	"fmt"
	"sort"
)

func main() {
	var myslice = []string{"1", "2", "3", "4"}
	fmt.Println(myslice)
	myslice = append(myslice, "5", "6")
	fmt.Println(myslice)
	myslice = append(myslice[1:]) //append method to increase the size of the slice
	fmt.Println(myslice)
	hel := make([]int, 4) // using make for creating slice
	hel[0] = 1
	hel[1] = 2
	fmt.Println(hel)
	fmt.Println(sort.IntsAreSorted(hel))

	str := []string{"n", "a", "v", "e", "e", "n"}
	index := 2
	str = append(str[:index], str[index+1:]...) //remove the particular index value
	fmt.Println(str)

}
