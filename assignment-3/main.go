package main

import "fmt"

// Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

func main() {
	num := 5
	result := []int{}
	for i := 0; i <= num; i++ {
		pushNum := 0
		j := i
		for j > 0 {
			if j%2 == 1 {
				pushNum++
			}
			j /= 2
		}
		result = append(result, pushNum)
	}
	fmt.Println(result)
}
