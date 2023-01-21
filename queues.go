package main

import "fmt"

type Queues struct {
	data []interface{}
}

func (q *Queues) Enqueue(data ...interface{}) {
	for i := range data {
		q.data = append(q.data, data[i])
	}
}

func (q *Queues) Dequeue() interface{} {
	if len(q.data) == 1 {
		temp := q.data[0]
		q.data = nil
		return temp
	} else if len(q.data) == 0 {
		return nil
	} else {
		temp := q.data[0]
		q.data = q.data[1:]
		return temp

	}

}

func main() {
	stackOne := Queues{}
	stackOne.Enqueue(1515151)
	fmt.Println(stackOne.data...)
	stackOne.Dequeue()
	fmt.Println(stackOne.data...)
	stackOne.Dequeue()
	fmt.Println(stackOne.data...)
}
