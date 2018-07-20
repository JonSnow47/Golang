/*
 * Revision History:
 *     Initial: 2018/05/15        Chen Yanchen
 */

package main

import "fmt"

func main()  {
	addAccount()
}

var m map[string]float64

func init()  {
	m = make(map[string]float64)
}

// 在 map 中检查是否存在此人，不存在就添加
func addAccount()  {
	m = make(map[string]float64)
	m["Apple"] = 3.2
	m["Banana"] = 2.4
	fmt.Println(m["Apple"])
}