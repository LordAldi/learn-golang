package main

import (
	"fmt"
	"strings"
)

func main() {
	//To search for a smaller string in a bigger string
	// Contains(s, substr string) bool
	fmt.Println(strings.Contains("test", "es"))
	// => true

	// func Count(s, sep string) int
	//To count the number of times a smaller string occurs in a bigger string,
	fmt.Println(strings.Count("test", "t"))
	// => 2

	// To determine if a bigger string starts with a smaller string
	// func HasPrefix(s, prefix string) bool
	fmt.Println(strings.HasPrefix("test", "te"))
	// => true
}
