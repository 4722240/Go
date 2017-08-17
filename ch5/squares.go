package main

import (
	"fmt"
)

func square() func() int{
	var x int
	return func() int {
		x++
		return x*x
	}
}

func aaa() {
	var i = 1
	var f = func() int {
		i++
		return i
	}
	f()
	fmt.Println(i)
}

func main() {
	//f := square()
	//f1 :=square()
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f1())
	//fmt.Println(f1())
	aaa()
}
