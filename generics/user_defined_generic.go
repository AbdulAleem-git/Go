package main

import "fmt"

func newGenericFunc[age int64 | float64](myAge age) {
	fmt.Println(myAge)
}

func main() {
	fmt.Println("Go Generics Tutorial")
	var testAge int64 = 23
	var testAge2 float64 = 24.5

	newGenericFunc(testAge)
	newGenericFunc(testAge2)
}
