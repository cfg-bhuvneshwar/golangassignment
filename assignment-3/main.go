package main

import (
	"fmt"
	"strconv"
)

// Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

// In progress

func main() {
	n := int64(2)

	binary := strconv.FormatInt(n, 2)

	fmt.Println(binary)

}
