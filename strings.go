package main

import (
	"fmt"
)

func printchars(s string) {
	fmt.Printf("Characters: ")
	for _, r := range s{
		fmt.Printf("%c ",r)
	}
}
func mutate(s []rune) string {  
    s[0] = 'k' 
    return string(s)
}
func main(){
	name := "Hello World"
    fmt.Printf("String: %s\n", name)
    printchars(name)
    fmt.Printf("\n")
    fmt.Printf("\n\n")
    name = "Señor"
    fmt.Printf("String: %s\n", name)
    printchars(name)
    fmt.Printf("\n")

	country := "india"
	fmt.Println(mutate([]rune(country)))
}