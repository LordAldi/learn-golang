package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}
type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}
func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}
func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
func main() {
	kids := []Person{
		{"Jill", 9},
		{"ack", 10},
		{"Jacky", 7},
		{"John", 6},
		{"milly", 12},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}