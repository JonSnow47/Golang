package main

import (
	"fmt"
	"math"
)

func main() {
	fib()
}

func os() {
	var s []float64
	for i := float64(0); i < 10; i++ {
		fmt.Println(math.Pow(2, i))
		s = append(s, math.Pow(2, i))
	}
	fmt.Println(s)
}

func fib() {
	a, b := 0, 1
	for b < 2000 {
		a, b = b, a+b
		fmt.Println(b)
	}
}
