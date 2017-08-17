package main

import (
	"fmt"
	//"time"
)

func main() {
	naturals :=make(chan int)
	squares  :=make(chan int)

	go func() {
		for x:=0; x<100 ; x++  {
			naturals <- x
		}
		//close(naturals)
	}()
	
	//go func() {
	//	for x := range naturals{
	//		squares <- x * x
	//	}
	//	close(squares)
	//}()

	go func() {
		for  y:=0; y<100 ; y++{
			x := <- naturals
			squares <- x * x
			//time.Sleep(100*time.Microsecond)
		}
	}()

	for y:=0; y<100 ; y++{
		fmt.Println(<-squares)
	}
}
