package main

import (
	"crypto/rand"
	"sync"
)

const N = 10 
func main() {
 m := make(map[int]int) 
 wg := &sync.WaitGroup{} 
 wg.Add(N) 
 for i := 0; i < N; i++ { 
 go func() { 
 defer wg.Done() 
 m[rand.Int()],_ = rand.Int() 
 }() 
 } 
 wg.Wait() 
 println(len(m)) 
} 
