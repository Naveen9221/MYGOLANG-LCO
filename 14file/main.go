package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to file handiling")

	content := "this is about file handling in go lang"

	file, err := os.Create("./myfiles.txt")
	checknilerr(err)

	len, err := io.WriteString(file, content)

	fmt.Println(len)

	checknilerr(err)

	defer file.Close()

	readFile("./myfiles.txt")

}
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checknilerr(err)
	fmt.Println("file data", string(databyte))

}

func checknilerr(err error) {
	if err != nil {
		panic(err)
	}

}
