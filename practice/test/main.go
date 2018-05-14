package main

import (
	"fmt"
)

func main() {
	l := []int{1, 2, 3, 4}
	fmt.Println(Sum(0, l))
}

func Sum(sum int, l []int) int {
	if len(l) < 1 {
		return sum
	}
	sum = sum + l[0]
	return Sum(sum, l[1:])
}
