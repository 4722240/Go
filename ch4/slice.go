package main

import (
	"fmt"
)

func main() {
	a := []int {1,1,1,1,1}
	b := a[0:2]
	//c := make([]int , 3)
	//copy(c,a[0:3])
	//a[1] = 0
	d := append(b,6);
	//d[2] = 0;
	printSlice(a)
	printSlice(b)
	//printSlice(c)
	printSlice(d)
}

func printSlice(s []int){
	len := len(s)
	for i := 0;i < len ;i++ {
		fmt.Printf("%d ",s[i])
	}
	fmt.Printf("\n")
}
