package main

import (
	"fmt"
	"math"
)

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}
type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}
func circleAreaP(c *Circle) float64 {
	c.r = 3

	return math.Pi * c.r * c.r
}

// type Shape interface {
// 	area() float64
// }

func main() {
	//initialization
	// var c Circle
	// c2 := new(Circle)
	c := Circle{x: 0, y: 0, r: 5}
	c2 := Circle{0, 0, 7}
	fmt.Println(c2.area())
	fmt.Println(circleAreaP(&c))
	fmt.Println(circleArea(c))
	a := new(Android)
	a.Person.Talk()
	a.Talk()

	//for pointer
	// c := &Circle{0, 0, 5}
}
