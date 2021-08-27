package main

import (
	"fmt"
	"testing"
)

func multiply (a int , b int) int{
	return a * b
}
func Test_multiply(t *testing.T){
	got := multiply(2,3)
	r := 2 * 3
	if got != r {
		fmt.Printf("got %d correct %d",got , r)
	}
}