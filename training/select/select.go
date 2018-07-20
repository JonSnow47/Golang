package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 6; i >= 0; i-- {
		switch_test(i)
	}
}

//
func switch_test(n int) {
	switch n {
	case 0, 1:
		fmt.Println("Done.")
		time.Sleep(time.Second)
		break
	default:
		fmt.Println("Now:", n)
		time.Sleep(time.Second)
	}
}

func select_test(ch1, ch2 chan int) {
	select {
	case <-ch1:
		fmt.Println("get ch1 element.")
	case <-ch2:
		fmt.Println("get ch2 element.")
	default:
		fmt.Println("waiting...")
		time.Sleep(time.Second * 1)
	}
}
