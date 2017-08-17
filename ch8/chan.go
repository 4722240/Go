package main

import "fmt"

func main() {
	ch := make(chan int,1)
	ch1 := make(chan int,1)
	cha := [...]chan(int){ch,ch1}

	for i:=1;i<=10;i++{
		select {
		case x := <-cha[0]:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
