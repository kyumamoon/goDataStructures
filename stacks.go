package main

import "fmt"

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(data ...interface{}) {
	for i := range data {
		s.data = append(s.data, data[i])
	}
}

func (s *Stack) Pop() interface{} {
	if len(s.data) == 1 {
		temp := s.data[0]
		s.data = nil
		return temp
	} else if len(s.data) == 0 {
		return nil
	} else {
		temp := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return temp
	}

}

func main() {
	stackOne := Stack{}
	fmt.Println(stackOne.data...)
	stackOne.Push(1515151)
	fmt.Println(stackOne.data...)
	stackOne.Pop()
	fmt.Println(stackOne.data...)
	stackOne.Pop()
	fmt.Println(len(stackOne.data))
	stackOne.Push(1515151)
	fmt.Println(stackOne.data...)
	fmt.Println(len(stackOne.data))
}
