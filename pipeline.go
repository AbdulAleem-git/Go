package main

import (
	"fmt"
	"time"
)
func main()  {
	counter := make(chan int)
	squares := make(chan int)

	go func(){
		for x := 0 ; x < 10; x++{
			counter <- x
			time.Sleep(500*time.Millisecond)
		}
		close(counter)
	}()

	go func (){
		for v := range counter{
			squares <- v * v
		}	
		close(squares)		
	}()

	for v := range squares{
		fmt.Printf("%d ",v)
	}
	fmt.Println()
}
