package main

import (
	"fmt"
	"math"
)

func main() {
	os()
}

func os() {
	var s []float64
	for i := float64(0); i < 10; i++ {
		fmt.Println(math.Pow(2, i))
		s = append(s, math.Pow(2, i))
	}
	fmt.Println(s)
}
