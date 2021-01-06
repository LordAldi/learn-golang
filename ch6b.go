package main

import (
	"fmt"
)

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func main() {
	// defer second()
	// first()
	// fmt.Println("3rd")

	//defer is often used when resources need to be freed in some way. For example, when we open a file, we need to make sure to close it later.
	// f, _ := os.Open(filename)
	// defer f.Close()

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")

}
