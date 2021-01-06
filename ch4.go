package main

import "fmt"

func main() {
	//-------------------------Control Structures-----------------------//
	fmt.Println("For statement")
	i := 1
	for i <= 10 {
		if i%2 == 0 {
			fmt.Println("even")

		} else {
			fmt.Println("odd")
		}
		switch i {
		case 0:
			fmt.Println("zero")
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		case 4:
			fmt.Println("four")
		case 5:
			fmt.Println("five")

		}
		i++
	}

}
