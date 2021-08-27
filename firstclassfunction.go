package main

import "fmt"

//Anonymous function

// func main(){
// 	a := func (){
// 		fmt.Println("hello world")
// 	}
// 	a()
// 	func (n string)  {
// 		fmt.Println(n)
// 	}("abdul")
// }

//User defined function

// type add func(a int , b int) int
// type sub func(a int , b int) int

// func main(){
// 	var a add
// 	a = func(a int, b int)int{return a + b}
// 	fmt.Println(a(2,3))
// 	var b sub
// 	b = func (a, b int )int{return a-b}
// 	fmt.Println(b(23,45))
// }

//Passing Function to other function

// func print(f func(a , b int)int){
// 	fmt.Println(f(54,34))
// }

// func main(){
// 	a := func(a int , b int)int{
// 		return a+b
// 	}
// 	print(a)
// 	b := func (a, b int)int  {
// 		return a-b
// 	}
// 	print(b)
// }

// Returning a function from a function
// func simple()func(a, b int) int{
// 	f := func(a , b int ) int{return a+b}
// 	return f
// }

// func main() {
//     s := simple()
//     fmt.Println(s(60, 7))
// }

// function Closure

// func main(){
// 	a := 77
// 	f := func(){
// 		fmt.Println(a)
// 	}
// 	f()
// }

// func appenStr() func(string)string{
// 	hello := "Hello"

// 	f := func(str string) string{
// 		return hello +  " "+ str
// 	}
// 	return f
// }

// func main(){
// 	a:= appenStr()
// 	b:= appenStr()
// 	fmt.Println(a("Abdul"))
// 	fmt.Println(b("Gopher"))
// 	fmt.Println(a("Gopher"))
//     fmt.Println(b("!"))
// }

// Example of first Class function

// type student struct{
// 	firstname string
// 	lastname string
// 	grade string
// 	country string
// }

// func filter (s [] student ,f func(student) bool)[]student{
// 	var r []student
// 	for _ , v := range s {
// 		if f(v) == true{
// 			r = append(r, v)
// 		}
// 	}
// 	return r
// }
// func main(){
// 	s1 := student{
//         firstname: "Naveen",
//         lastname:  "Ramanathan",
//         grade:     "A",
//         country:   "India",
//     }
//     s2 := student{
//         firstname: "Samuel",
//         lastname:  "Johnson",
//         grade:     "B",
//         country:   "USA",
//     }
// 	s := []student{s1 , s2 }
// 	f := func (s student)  bool{
// 		if s.grade =="A"{
// 			return true
// 		}else{
// 			return false
// 		}

// 	}
// 	res := filter(s,f)
// 	fmt.Print(res)
// }


type Employee struct{
	firstname string
	lastname string
	gender string
}
func filter(e []Employee , female func(Employee)bool)[]Employee{
	var r[] Employee
	for _ , v := range e {
		if female(v) == true{
			r = append(r, v)
		}
	}
	return r
}
func main(){
	e1 := Employee{
		"anjali","kumar","female",
	}
	e2 := Employee{
		"rama","singhal","female",
	}
	e3 := Employee{
		"raja","sinha", "male",
	}
	e := []Employee{e1,e2,e3}
	f := func(e Employee)bool{
		if e.gender == "female"{
			return true
		}else{
			return false
		}
	}
	fmt.Println(filter(e,f))
}

// - targets: ['m3dbnode-0.m3dbnode.clamp:7203', 'm3dbnode-1.m3dbnode.clamp:7203', 'm3dbnode-2.m3dbnode.clamp:7203'