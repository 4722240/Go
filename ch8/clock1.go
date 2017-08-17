package main

import (
	"io"
	"log"
	"time"
	"net"
	"fmt"
)


func main() {
	listener ,err := net.Listen("tcp","localhost:8888")
	if(err != nil){
		log.Fatal(err)
	}
	for{
		conn ,err := listener.Accept()
		if(err != nil){
			log.Fatal(err)
			continue
		}
		handleConn(conn)
	}
}

func  handleConn(c net.Conn ){
	defer c.Close()
	for  {
		fmt.Println(c)
		_,err :=io.WriteString(c,time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1*time.Second)
	}
}
