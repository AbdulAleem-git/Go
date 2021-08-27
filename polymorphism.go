package main

import "fmt"
type Income interface{
	calculate () int 
	source () string
}
type fixedBilling struct{
	projectName string
	bidAmount int
}
type timeandmaterial struct{
	projectName string
	noOfhours int
	hourlyRate int
}

func (fb fixedBilling) calculate() int{
	return fb.bidAmount
}
func (fb timeandmaterial) calculate()int{
	return fb.noOfhours * fb.hourlyRate
}
func (fb fixedBilling) source () string{
	return fb.projectName
}
func (fb timeandmaterial) source () string{
	return fb.projectName
}

func main(){
	p1 := fixedBilling{projectName: "koo",bidAmount: 123}
	p2 := timeandmaterial{projectName: "kookoo",noOfhours: 12,hourlyRate: 23}
	fmt.Println(p1.calculate() , p1.source())
	fmt.Println(p2.calculate() , p2.source())

}