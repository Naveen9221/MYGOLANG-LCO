package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("my time conversion")
	present := time.Now()
	fmt.Println(present.Format("02 - 01 - 2006 15:04:05 Monday"))
	createdate := time.Date(2020, time.December, 12, 23, 23, 34, 0, time.UTC)
	fmt.Println(createdate.Format("01-02-2006 Monday"))
}
