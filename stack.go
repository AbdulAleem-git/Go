package main

import "fmt"


type mystack interface{
	push(int)
	pop() int 
}
type stack struct{
	data [] int
}

func (m *stack) push(i int) {
	m.data = append(m.data, i)	 
}
func (m *stack) pop() int {
	if len(m.data) == 0 {
		return -1 
	}
	item :=  m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	return item
}
func main(){
	m1 := &stack{}
	for i := 0 ; i < 10 ; i++{
		m1.push(i)
	}
	fmt.Println(m1.data)
	for i := 0 ; i < 10 ; i++{
		fmt.Printf("%d ",m1.pop())
	}
	var i interface{} = 32
	switch i.(type){
	case int:
		fmt.Println("Integer")
	}
	fmt.Println()
}