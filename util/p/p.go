package p

import (
	"sync"
)

func P(p, n int, f func(int int))  {
	// p is process num
	if p<1 || n <1 ||f==nil {
		return
	}

	ch := make(chan int, n)
	for i := 0; i<n; i++ {
		ch <- i
	}
	defer close(ch)

	wg := sync.WaitGroup{}
	wg.Add(p)

	if p > n {p = n }
	for i := 0; i<p; i++{
		// defer runtime
		defer wg.Done()
		for j := range ch {
			go func() {
				f(j)
			}()
		}
	}
	wg.Wait()
}