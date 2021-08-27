package main

import "fmt"

// type Vowelfinder interface{
// 	FindVowels() [] rune

// }
// type myString string

// func (m myString) FindVowels() []rune{
// 	var vowel []rune
// 	for _, s  := range m {
// 		if s  == 'a' || s == 'e' || s == 'i' || s =='o' || s == 'u'{
// 			vowel = append(vowel , s)
// 		}
// 	}
// 	return vowel
// }
// func (m myString) FindVowels() []string{
// 	var vowel []string
// 	for _, s  := range m {
// 		if s  == 'a' || s == 'e' || s == 'i' || s =='o' || s == 'u'{
// 			vowel = append(vowel , string(s))
// 		}
// 	}
// 	return vowel
// }

// func main(){
// 	s := myString("i love you")
// 	v := s.FindVowels()
// 	for _, c := range v{
// 		fmt.Printf("%s ",c)
// 	}
// 	fmt.Printf("\n")
// }

// type SalaryCalculator interface{
// 	calculateSalary() int
// }

// type Permanent struct{
// 	empId int
// 	basicpay int
// 	pf int
// }
// type contract struct{
// 	empId int
// 	basicpay int
// }
// type test struct{
// 	name string
// }

// func (p Permanent) calculateSalary() int{
// 	return p.basicpay + p.pf
// }
// func (c contract) calculateSalary() int{
// 	return c.basicpay
// }
// func ( c test) calculateSalary() string{
// 	return c.name
// }
// func totalExpense(s [] SalaryCalculator){
// 	expense := 0

// 	for _ , v := range s {
// 		expense = expense + v.calculateSalary()
// 	}
// 	fmt.Printf("Total Expenses : %d\n",expense)
// }
// func main(){
// 	pemp1 := Permanent{
// 		empId: 1,
// 		basicpay: 4500,
// 		pf:20,
// 	}
// 	pemp2 := Permanent{
// 		empId: 20,
// 		basicpay:7800,
// 		pf:20,
// 	}
// 	pemp3 := Permanent{
// 		empId: 3,
// 		basicpay: 7500,
// 		pf:28,
// 	}
// 	cemp1 := contract{
// 		empId: 1,
// 		basicpay: 23334,
// 	}
// 	cemp2 := contract{
// 		empId: 2,
// 		basicpay: 2334,
// 	}

// 	employees := [] SalaryCalculator{pemp1,pemp2,pemp3,cemp1,cemp2}
// 	totalExpense(employees)
// }

// type MilageCalculator interface{
// 	calculateDiesel() int
// 	calculatePetrol() int
// 	calculateCNG() int
// }
// type PetrolCar struct{
// 	carname string
// 	totalpetrol int
// 	totalKM int
// }
// type DieselCar struct{
// 	carname string
// 	totalDiesal int
// 	totalKM int
// }
// type CNGCar struct{
// 	carname string
// 	totalCNG int
// 	totalKM int
// }
// func(p PetrolCar)calculatePetrol() int{
// 	return (p.totalKM / p.totalpetrol)
// }
// func(d DieselCar)calculateDiesel() int{
// 	return (d.totalKM / d.totalDiesal)
// }
// func(c CNGCar)calculateCNG() int{
// 	return (c.totalKM / c.totalCNG)
// }
// func totalExpense(m []MilageCalculator){
// 	expense := 0
// 	for _, v := range m{
// 		switch v.(type){
// 			PetrolCar:
// 				expense :=
// 		}
// 	}
// }
// func main(){

// }

// Implementing the type assertion
// func assert(i interface{}) {

// 	switch i.(type) {
// 	case int:
// 		s := i.(int)
// 		fmt.Printf("Integer Value: %d\n", s)
// 	case string:
// 		s := i.(string)
// 		fmt.Printf("String Value: %s\n", s)
// 	default:
// 		s, ok := i.(float64)
// 		if !ok {
// 			fmt.Printf("not a correct value")
// 		} else {
// 			fmt.Printf("Float Value: %f\n", s)
// 		}
// 	}
// }
// func main() {
// 	var i interface{} = 23
// 	assert(i)
// 	var s interface{} = "hello"
// 	assert(s)
// 	var f interface{} = 23.45
// 	assert(f)
// }

// Implementing interfaces using pointer receivers vs value receivers
// type Describer interface{
// 	Describe ()
// }
// type Person struct {
// 	name string
// 	age int
// }

// func(p Person) Describe(){
// 	fmt.Printf("name: %s , age: %d\n",p.name,p.age)
// 	p.name = "aleem"
// }
// type Address struct{
// 	state string
// 	country string
// }
// func (a *Address) Describe(){
// 	a.state = "uttarakhand"
// 	fmt.Printf("in method: State %s Country %s\n", a.state,a.country)
// }

// func main(){
// 	var d1 Describer
// 	p1 := Person{
// 		name: "abdul",
// 		age: 25,
// 	}
// 	d1 = p1
// 	d1.Describe()
// 	p2 := Person{
// 		name: "aleem",
// 		age: 26,
// 	}
// 	d1 = &p2
// 	d1.Describe()
// 	a := Address{
// 		state: "uttar pradesh",
// 		country: "india",
// 	}
// 	fmt.Printf("In main: state %s country %s\n", a.state,a.state)
// 	a.Describe()
// 	fmt.Printf("In main: state %s country %s\n   ", a.state,a.state)

// }

// Implementing multiple interfaces

// type SalaryCalculator interface{
// 	DisplaySalary()
// }
// type LeaveCalculator interface{
// 	CalculateLeaveLeft() int
// }
// type EmployeeOps  interface{
// 	SalaryCalculator
// 	LeaveCalculator
// }

// type Employee struct {
//     firstName string
//     lastName string
//     basicPay int
//     pf int
//     totalLeaves int
//     leavesTaken int
// }

// func (e Employee) DisplaySalary(){
// 	fmt.Printf("%s%s has Salary %d", e.firstName ,e.lastName,(e.basicPay + e.pf))

// }
// func (e Employee) CalculateLeaveLeft() int{
// 	return (e.totalLeaves - e.leavesTaken)
// }

// func main(){
// 	e := Employee {
//         firstName: "Abdul",
//         lastName: "Aleem",
//         basicPay: 5000,
//         pf: 200,
//         totalLeaves: 30,
//         leavesTaken: 5,
//     }
// 	var s EmployeeOps = e
// 	s.DisplaySalary()

// 	fmt.Println("\nLeaves left =", s.CalculateLeaveLeft())

// }

// Zero Valu of Interface


type Describer interface{
	Describe()
}
func main(){
	var d1 Describer
	if d1 == nil {
		fmt.Print("empty")
	}
	d1.Describe()
}