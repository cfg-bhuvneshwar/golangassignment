package main

import "fmt"

// Given two integers a and b, return the sum of the two integers without using the operators + and -.
func main() {
	a := 5
	b := 10

	for i := 1; i <= b; i++ {
		a++
	}
	fmt.Println(a)

}
