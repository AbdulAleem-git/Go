package main

import (
	"fmt"
	"net"
)

func server(input chan string){
	listener ,_ := net.Listen("tcp", ":8080")
	conn ,_ := listener.Accept()
	fmt.Fprintf(conn, fmt.Sprintf("server : %s", <-input))
}
func client(output chan string){
	conn ,_ := net.Dial("tcp","127.0.0.1:8080")
	output <- "hello world"
	fmt.Fprintf("client:%s",<-output)
}
func main(){
	fmt.Println("going start server")

}