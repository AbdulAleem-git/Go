package main

import (
	"bufio"
	"fmt"
	"net"
)

func main(){
	fmt.Println("starting server on localhost:8080")
	ls,_ := net.Listen("tcp",":8080")

	conn , _ := ls.Accept()

	for {
		message ,_ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf( "Message Recieved: %s", message)
		fmt.Fprintf(conn , message+"\n")		
	}
}