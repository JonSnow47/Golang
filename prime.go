/*package main

import "fmt"

func prime(num int) {
	var result [num] int
	var i, j = 1, 2
	for i = 1; i < num; i++ {
		for j = 2; j < i; j++ {
			if i % j == 0; 
			//break;
		}
		result += i
	}
	return result
}

func main() {
	for true {
		fmt.Println("endless loop")
	}
}