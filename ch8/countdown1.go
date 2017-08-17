package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	fmt.Println("commencing countdown")
	tick := time.Tick(1*time.Second)
	absort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		absort <- struct{}{}
	}()
	for countdown := 10; countdown > 0 ; countdown--  {
		fmt.Println(countdown)
		select{
		case y := <- tick:
			fmt.Println(y)
		case <- absort:
			fmt.Println("abort")
			return
		}
	}
	//lanch()
}
