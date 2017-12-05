package main

import "fmt"
import "math"

var a,b int = 4,6

func bbc() {
	var a, b int = 3, 4
	var c float64 = math.Sqrt(float64(a * a + b * b))
	print(a, b, c)
}

func hello() {
	var firstname, lastname string
	fmt.Print("Please input your name: ")
	fmt.Scanln(&firstname, &lastname)
	fmt.Printf("Hi,%s %s!\nThanks for use Golang.\n", firstname, lastname)
}
func main() {
	hello()
}