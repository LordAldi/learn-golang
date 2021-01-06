package main

import "fmt"

func main() {

	//------------------VARIABLE-------------------------
	fmt.Println("Hello World!")
	fmt.Println(len("Hello, World"))
	fmt.Println("Hello, World"[1])
	fmt.Println("Hello, " + "World")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 + 1 =", 1.0+1.0)

	var x string = "Hello, World"
	fmt.Println(x)

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)

}
