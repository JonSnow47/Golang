package example

import "fmt"

func prime(num int) {
	//var result string = ""
	var i, j = 1, 2
	for i = 1; i < num; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("prime: %d\n", i)
		}
	}
}
