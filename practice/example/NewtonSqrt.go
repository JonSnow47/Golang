package example

import (
	"fmt"
	"math"
)

func Sqrt(x float64) {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = (z*z - x) / (2*z)
		fmt.Printf("z[%v] = %v\n", i, z)
	}
	return
}

func NewtonSqrt() {
	fmt.Println("sqrt(2) =",math.Sqrt(2))
	Sqrt(2)
}
