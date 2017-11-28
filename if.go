package main

import "fmt"

func test2() {
	var a int = 4
	print("start\n")
	//var b int = 6
	if a == 3 {
		fmt.Println("a =",a)
	}else if a == 4 {
		fmt.Println("a =",a)
	}
}
func main() {
	test2()
}