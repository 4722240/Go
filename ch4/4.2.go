package main

import (
	"fmt"
	"os"
	"crypto/sha256"
	"crypto/sha512"
	"crypto"
)

func main() {
	len := len(os.Args)
	if(len == 2){
		fmt.Printf("%x \n",sha256.Sum256([]byte(os.Args[1])))
	}else if os.Args[2] == "SHA384" {
		fmt.Printf("%x \n",crypto.SHA384.New().Sum([]byte(os.Args[1])))
	}else if os.Args[2] == "SHA512" {
		fmt.Printf("%x \n",sha512.Sum512([]byte(os.Args[1])))
	}

}
