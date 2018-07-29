package main

import (
	"./student"
	"fmt"
	"reflect"
)

type S interface {
	SetName(string)
	GetName() string
	GetSchool() string
}

func main() {
	fmt.Println([]byte("oj  bk"))
	s := student.New()
	f(s)
	fmt.Println(s)
	f1()
	f2()
}

func f(i S) {
	i.SetName("jon")
}

type T []interface{}

// 判断类型
func f1() {
	t := make(T, 5)
	t[0] = "stringggggg"
	t[1] = 11111111111
	t[2] = 1.0000000001
	t[3] = make(chan int)
	t[4] = T{}
	for i, _ := range t {
		switch t[i].(type) {
		case string:
			fmt.Printf("t[%d] type of string, value is %v\n", i, t[i])
		case float64:
			fmt.Printf("t[%d] type of float64, value is %v\n", i, t[i])
		case int:
			fmt.Printf("t[%d] type of int, value is %v\n", i, t[i])
		case chan int:
			fmt.Printf("t[%d] type of chan int, value is %v\n", i, t[i])
		default:
			fmt.Printf("t[%d] type of %s, value is %v\n", reflect.TypeOf(t[i]), t[i])
		}
	}
}

//////////////////////////////////////////////////////////////
type cat struct{}

func (*cat) Speak() {
	fmt.Println("miao~")
}

type dog struct{}

func (*dog) Speak() {
	fmt.Println("wang!")
}

type Animal interface {
	Speak()
}

func f2() {
	animals := []Animal{&cat{}, &dog{}}
	for _, v := range animals {
		v.Speak()
	}
}
