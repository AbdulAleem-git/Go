package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
)

// {"datetime":"2023-06-27 22:22:19.62710.193625","V":"10","partition":"p2"}

type Timeseries struct {
	Datetime  string
	Value     string
	Partition string
}

type CirculerQueue struct {
	items  []Timeseries
	front  int
	rear   int
	length int
	lock   sync.Mutex
}

func initQueue(capacity int) *CirculerQueue {
	return &CirculerQueue{
		items:  make([]Timeseries, capacity),
		front:  0,
		rear:   -1,
		length: 0,
	}
}

func (c *CirculerQueue) Enqueue(V Timeseries) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.length == len(c.items) {
		return false
	}
	c.rear = (c.rear + 1) % len(c.items)
	c.items[c.rear] = V
	c.length++
	return true
}

func (c *CirculerQueue) Dequeue() (Timeseries, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.length == 0 {
		return Timeseries{}, errors.New("Empty queue")
	}
	item := c.items[c.front]
	c.front = (c.front + 1) % len(c.items)
	c.length--
	return item, nil
}

func (cq *CirculerQueue) IsEmpty() bool {
	return cq.length == 0
}

func (cq *CirculerQueue) IsFull() bool {
	return cq.length == len(cq.items)
}

func Writer(file string, q *CirculerQueue, wg *sync.WaitGroup) {
	defer wg.Done()
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(f)
	defer f.Close()
	// var buffer bufio.Writer
	var ts Timeseries

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()

		json.Unmarshal(line, &ts)
		q.Enqueue(ts)
	}
}
func Reader(q *CirculerQueue, wg *sync.WaitGroup)

func main() {
	// queue := initQueue(5)
	var wg sync.WaitGroup
	wg.Add(2)
	q := initQueue(20)
	go Writer("data.json", q, &wg)
}
