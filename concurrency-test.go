package main

import (
	"fmt"
	"sync"
	"time"
)

// func printchars(s string, wg *sync.WaitGroup){
// 	defer wg.Done()
// 	for _, c := range s {
// 		time.Sleep(time.Millisecond * 200)
// 		fmt.Printf("%c",c)
// 	}
// 	fmt.Println()
// }

// func printchars_channel(s string , ch chan rune){
// 		for _ , c := range s {
// 			time.Sleep(time.Millisecond * 200)
// 			ch <- c
// 		}
// 		close(ch)
// }

// func main(){
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go printchars("Hello World",&wg)
// 	wg.Wait()

// 	ch := make(chan rune)
// 	go  printchars_channel("whats up",ch)
// 	// for {
// 	// 	if c , ok := <-ch;!ok{
// 	// 		break
// 	// 	}else{
// 	// 		fmt.Printf("%c ",c)
// 	// 	}
// 	// }
// 	for c := range ch{
// 		fmt.Printf("%c ",c)
// 	}
// 	fmt.Println()
// 	fmt.Println("Main thread exited")
// }


func 