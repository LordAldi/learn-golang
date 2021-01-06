package main

import "fmt"

func average(arr []float64) float64 {
	total := 0.0
	for _, v := range arr {
		total += v
	}
	return (total / float64(len(arr)))
}
func f() (int, int) {
	return 5, 6
}

//variadic parameter.
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	x, y := f()
	fmt.Println(x, y)
	fmt.Println(add(1, 2, 3))

	// fmt.Println(add(xs...))

	//Closure
	z := 0
	increment := func() int {
		z++
		return z
	}
	fmt.Println(increment())

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4

	//recursion

	fmt.Print(factorial(4))

}
