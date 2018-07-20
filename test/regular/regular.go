// 常用正则表达式
package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "asdgaioi oisdnglkanfb" +
		"aiwjofinaewflakfasmdoifm"

	fmt.Println(regular(s))
}

// 匹配第一个空格前的所有字符
func regular(s string) string {
	reg := regexp.MustCompile(`(?U)^.*\s`)
	return reg.FindString(s)
}
