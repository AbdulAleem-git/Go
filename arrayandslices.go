package main

import "fmt"

func  add(s []int)  {
	s = append(s , 34)
	fmt.Println(s)
}
func  main()  {
	a := [...]int{12, 78, 50,2,3}   // an Array
	// c := []int{12, 78, 50,2,3}   // a slice
	
	// fmt.Println(a)
	// b := a // An array always creat copy for assignment becuase array are value type.
	// b[0] = 32
	// c = append(c, 34)
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	slice := a[:]
	add(slice	)
	fmt.Println(slice)
	
	

	cars := []string{}
    fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "aleem")
	var names []string
	names = append(names, "abdul")
	names =append(names, cars...)
	fmt.Println(names)

    
}