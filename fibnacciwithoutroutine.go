package main

import (
	"fmt"
	"time"
)


func spinner(delay time.Duration){
	for {
		for _, r := range `/-\`{
			fmt.Printf("\r%c",r)
			time.Sleep(delay)
		}
	}
}
func main(){

	go spinner(100 * time.Millisecond)
	n := 45
	f := fib(n)
	fmt.Println(f)
}
func fib (n int)int{
	if n == 0 || n == 1 {
		return n
	}
	return fib (n-1) + fib(n-2)
}