package main

import "fmt"

type collection interface {
	size() int
}

type Queue interface {
	collection
	New()
	add(int)
	remove() int
}

type Stack interface {
	collection
	New()
	push(int)
	pop() int
}

type que []int

func (q *que) New() {
	*q = make(que, 0)
}

func (q *que) add(value int) {
	*q = append(*q, value)
}

func (q *que) remove() int {
	if len(*q) == 0 {
		return -1
	}
	i := (*q)[0]
	*q = (*q)[1:]
	return i
}
func (q *que) size() int {
	return len(*q)
}

func (s *que) push(value int) {
	*s = append(*s, value)
}

func (s *que) pop() int {
	if len(*s) == 0 {
		return -1
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func main() {
	var q Queue
	queue := new(que)
	q = queue
	q.New()
	q.add(45)
	q.add(46)
	q.add(47)
	q.add(48)
	q.add(49)
	q.add(50)

	for q.size() != 0 {
		fmt.Printf("%d  ", q.remove())
	}

	var s Stack
	stack := new(que)
	s = stack

	s.push(45)
	s.push(46)
	s.push(47)
	s.push(467)
	s.push(48)
	s.push(49)
	s.push(445)

	for s.size() != 0 {
		fmt.Printf("%d  ", s.pop())
	}
}
