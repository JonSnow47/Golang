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

func regular(s string) string {
	reg := regexp.MustCompile(`(?U)^.*\s`)
	return reg.FindString(s)
}
