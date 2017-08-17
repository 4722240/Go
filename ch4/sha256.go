package main

import (
	"fmt"
	"crypto/sha256"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf(" %x \n %x \n %t \n %T \n %d \n",c1,c2,c1 == c2,c1,diff(c1,c2))
}

func diff(c1 [32]uint8, c2 [32]uint8) int {
	count := 0
	for i := 0; i < 32; i++ {
		if(c1[i] == c2[i]){
			count ++
		}
	}
	return count
}