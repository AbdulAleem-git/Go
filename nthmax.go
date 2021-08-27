package main

import (
	"fmt"
)
type maxheap struct{
	 h [] int 
}

func (m *maxheap)buildmaxheap(arr []int) {
	if len(arr) == 0 {
		fmt.Println("empty array cant be create a heap")
	}
	fmt.Println(arr)
	for _ , v := range arr {
		m.insert(v)
	}

}
func (m *maxheap) insert(ele int) {
	m.h = append(m.h, ele)
	m.heapifyUP(len(m.h)-1)
}
func (m *maxheap) heapifyUP(index int ){
	for m.h[index] > m.h[parent(index)]{
		m.swap(index , parent(index))
		index = parent(index)
	}	
}
func (m *maxheap) swap(i1 , i2 int){
	m.h[i1] , m.h[i2] = m.h[i2] , m.h[i1]
}
func parent (index int)int{
	return (index - 1) / 2 
}
func left(index int )int{
	return 2 * index  + 1
}
func right(index int)int{
	return 2 * index + 2
}

func (m *maxheap)extract() int {
	if len(m.h) == 0 {
		fmt.Printf("Empty Heap")
		return -1
	}
	ele := m.h[0]
	m.h[0] , m.h[len(m.h)-1] = m.h[len(m.h)-1] , m.h[0]
	m.h = m.h[:len(m.h)-1]
	m.heapifyDown(0)
	return ele
}
func (m *maxheap)heapifyDown(index int) {
	l , r := left(index) , right(index)
	lastindex := len(m.h)-1
	var childToCompare int
	for l <= lastindex{
		if l == lastindex{
			childToCompare = l
		}else if(m.h[l] > m.h[r]){
			childToCompare = l
		}else{
			childToCompare =r 
		}
		if m.h[childToCompare] > m.h[index]{
			m.h[childToCompare] , m.h[index] = m.h[index] , m.h[childToCompare]
			index = childToCompare
			l , r = left(childToCompare),right(childToCompare)
		}else{
			return
		}
	}
}
func main(){
	heap := &maxheap{}
	arr := []int{32,45,35,39,85,95,20,23}
	heap.buildmaxheap(arr)
	for _,_ = range heap.h{
			fmt.Printf("%d ",heap.extract())
	}
}
