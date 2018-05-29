/*
 * Revision History:
 *     Initial: 2018/05/28        Chen Yanchen
 */

package main

import (
	"time"
	"fmt"
)

func main()  {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 2500)   //阻塞，则执行次数为sleep的休眠时间/ticker的时间
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
