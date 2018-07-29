package main

import (
	"fmt"
	"strings"
)

func main() {
	Contains()
	ContainsAny()
	//ContainsRune()
	LastIndex()
	LastIndexAny()
	Join()
}

// string.Contains 检测 s 中是否包含 substr
func Contains() {
	s := `abcde`
	substr := `abbb` // a:true; f:flase; ab:true; ac:false
	result := strings.Contains(s, substr)
	fmt.Println("[strings.Contains]	", result)
}

// string.ContainsAny 检测 s 中是否包含 chars 中所含字符
func ContainsAny() {
	s := `abcde`
	chars := `acccc` // a:true; f:flase; ab:true; accc:true
	result := strings.ContainsAny(s, chars)
	fmt.Println("[strings.ContainsAny]	", result)
}

// ContainsRune reports whether the Unicode code point r is within s.
func ContainsRune() {
	s := `abcde0`
	var r rune = 0
	result := strings.ContainsRune(s, r)
	fmt.Println("[strings.ContainsRune]	", result)
}

// strings.LastIndex
func LastIndex()  {
	s := `abcde abcde abcde`
	substr := `abcde`
	result := strings.LastIndex(s,substr)
	fmt.Println("[strings.LastIndex]	", result)
}

// strings.LastIndexAny
func LastIndexAny()  {
	s := `abcde abcde abcde`
	char := `abcde`
	result := strings.LastIndexAny(s,char)
	fmt.Println("[strings.LastIndexAny]	", result)
}

// string.Join 将字符串数组a的每个元素通过特定字符sep连接起来
func Join() {
	a := []string{`aaa`, `bbb`}
	sep := "----"
	result := strings.Join(a, sep)
	fmt.Println("[strings.Join]	", result)
}
