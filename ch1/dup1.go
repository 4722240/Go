package main

import (
	"fmt"
	"os"
	"bufio"
)

func main()  {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		if input.Text() == "exit" || input.Text() == ""{
			break
		}
		count[input.Text()]++
		//fmt.Println(input.Text(),count[input.Text()])
	}

	for line,n := range count {
		if(n>=1){
			fmt.Printf("%d\t%s\n", n,line)
		}
	}
}
