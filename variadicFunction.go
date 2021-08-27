package main

import (
	"fmt"
)

func find(num int, nums ...int) {  
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}
func change_variadic(s... string) {  
	s = append(s, "playground")
	fmt.Println("slice in function", s)
	fmt.Println("len in function", len(s))
}
func change(s []string) {  
	s = append(s, "playground")
	fmt.Println("slice in function", s)
	fmt.Println("len in function", len(s))
}

func main() {  
    find(89, 89, 90, 95)
    find(45, 56, 67, 45, 90, 109)
    find(78, 38, 56, 98)
    find(87)

	welcome := make([]string ,10,10)
	welcome[0] = "hello"
	welcome[1] =  "world"
    change_variadic(welcome...)
	fmt.Println("slice in main", welcome)
	fmt.Println("len in main", len(welcome))
	change_variadic(welcome[0:3]...)
	fmt.Println("slice in main", welcome)
	fmt.Println("len in main", len(welcome))
	change_variadic(welcome[0:4]...)
	fmt.Println("slice in main",welcome)
	fmt.Println("len in main", len(welcome))
	change_variadic(welcome[0:5]...)
	fmt.Println("slice in main", welcome)
	fmt.Println("len in main", len(welcome))

	s := make ([]string , 5,5)
	s[0] ="main0"
	s[1] ="main1"
	s[2] ="main1"
	s[3] ="main1"
	s[4] ="main1"
	change(s[0:2])
	fmt.Println("slice in main",s)
	fmt.Println("len in main",len(s))
	
}