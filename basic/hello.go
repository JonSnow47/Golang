package main

import "fmt"

func hello() {
	var firstname, lastname string
	fmt.Print("Please input your name: ")
	fmt.Scanln(&firstname, &lastname)
	fmt.Printf("Hi,%s %s!\nThanks for use Golang.\n", firstname, lastname)
}
func main() {
	hello()
}