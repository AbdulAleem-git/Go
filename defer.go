package main

import (
	"fmt"
	"sync"
)

// func finished(){
// 	fmt.Println("Finished the function")
// }
// func largest(nums [] int) {
// 	 max := 0
// 	for _, v  := range nums{
// 		if max < v {
// 			max = v
// 		}

// 	}
// 	fmt.Println("The largest number is ",max)
// }

// func main(){
// 	nums := []int{12,3,4,5,4,3,5,45}
// 	largest(nums)
// }

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		// wg.Done()
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		// wg.Done()
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
	// wg.Done()
}
func main() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rec := []rect{r1,r2,r3}
	for _ , v := range rec{
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
    fmt.Println("All go routines finished executing")
}
