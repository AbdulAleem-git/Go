package main

import (
	"fmt"
	"sync"
)

// Using Mutex

// var x = 0
// func increment(wg *sync.WaitGroup , m *sync.Mutex){
// 	m.Lock()
// 	x = x + 1
// 	m.Unlock()
// 	wg.Done()
// }

// func main(){
// 	var wg sync.WaitGroup
// 	var m sync.Mutex
// 	for i := 0 ; i < 1000; i++{
// 		wg.Add(1)
// 		go increment(&wg , &m)
// 	}
// 	wg.Wait()
// 	fmt.Printf("final Value of X %d\n",x)
// }

// Using Channel

var x = 0
func increment(wg *sync.WaitGroup , ch chan bool){
	ch <- true 
	x = x + 1 
	<- ch 	
	wg.Done()
}
func  main(){
	var wg sync.WaitGroup 
	var ch chan bool
	ch = make(chan bool , 1)
	for i := 0 ; i < 10000 ; i++{
		wg.Add(1)
		go increment(&wg , ch)
		wg.Wait()
	}
	fmt.Println("final valu = ", x)
}