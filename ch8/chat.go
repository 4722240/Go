package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
)

func main() {
	listen ,err := net.Listen("tcp","localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	go broadcast()
	for  {
		conn,err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string

var (
	enter = make(chan client)
	leaving =make(chan client)
	message = make(chan  string)
)

func broadcast()  {
	clients := make(map[client]bool)
	for{
		select {
		case msg := <- message :
			for cli := range clients{
				cli <-msg
			}
		case cli := <- enter:
			clients[cli] = true
		case cli := <- leaving:
			delete(clients,cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn)  {
	ch := make(chan string)
	go clientWriter(conn,ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	message <- who + " has arrived"
	enter <- ch

	input := bufio.NewScanner(conn)
	for input.Scan()  {
		message <- who + " : " + input.Text()
	}

	leaving <- ch
	message <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn,ch <-chan string)  {
	for msg := range ch{
		fmt.Fprintln(conn,msg)
	}
}
