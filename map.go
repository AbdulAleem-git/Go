package main

import "fmt"

// func main(){
// 	var m map[string][]int
// 	m = make(map[string][]int)
// 	m["abdul"] = []int{122,23,4,5}
// 	m["aleem"] = []int{23,34,5}
// 	for k, v := range m {
// 		fmt.Println(k,v)
// 	}
// 	if k ,ok := m["abdul"]; ok{
// 		fmt.Println(k)
// 	}
// 	if k ,ok := m["saleem"]; ok{
// 		fmt.Println(k)
// 	}else{
// 		fmt.Println("no data")
// 	}
// 	delete(m, "abdul")
// 	for k, v := range m {
// 		fmt.Println(k,v)
// 	}
// 	employeesalary := map[string]int{
// 		"munna":23,
// 		"waseem":34,
// 	}
// 	for k, v := range employeesalary {
// 		fmt.Println(k,v)
// 	}
// }

//Lets create one string to struct map

type Employee struct{
	salary int
	country string
}

func main(){
	emp1 := Employee{
		200,
		"india",
	}
	emp2 := Employee{
		300,
		"afghanistan",
	}
	emp3 := Employee{
		400,
		"indonesia",
	}
	m := map[string]Employee{
		"Abdul":emp1,
		"rashid khan": emp2,
		"glane maxwell": emp3,
	}
	for k ,v := range m{
		fmt.Printf("Name: %s\n",k)
		fmt.Printf("Salary: %d\n",v.salary)
		fmt.Printf("Country: %s\n",v.country)
		fmt.Println()
	}

	employeesalary := map [string]int{
		"abdul":2000,
		"aleem":3000,
		"rashid":3445,
	}
	m2 := employeesalary
	m2["abdul"] = 34554
	fmt.Println(employeesalary["abdul"])
	delete(m2 , "abdul")
	fmt.Println(employeesalary["abdul"])
	
	// comparing of two map
	// if m2 == employeesalary{
	// 	fmt.Println("true")
	// }else{
	// fmt.Println("false")

	// }
}