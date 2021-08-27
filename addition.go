package main

import (
	"fmt"
	"time"
)

/*
func add(a, b int){
	fmt.Println(a+b)
}
func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	go func ()  {
		for i := 0 ; i < 5 ; i++{
			add(2,4)
		}
		wg.Done()
	}()
	wg.Wait()
}
*/
 func print(s string , c chan string){
	 for i :=0 ; i < 5 ; i++{
		 c <- s
		 time.Sleep(time.Millisecond*500)
	 }
	 close(c)
 }
 func main(){
	 c := make(chan string)
	 go print("weep" , c)
	 for msg := range c{
		//  msg , open := <- c
		//  if !open{
		// 	 break
		//  }
		 fmt.Println(msg)	
	 }
 }
	