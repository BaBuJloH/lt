package main

import "fmt"

type stack struct {
	index int
	data  []int
}

func newStack(num int) stack {
	s := stack{}
	s.index = -1
	s.data = make([]int, num)
	return s
}

func (s *stack) push(i int) {
	s.index++
	if s.index >= len(s.data) {
		s.data = append(s.data, i)
		return
	}
	s.data[s.index] = i
}

func (s *stack) pop() int {
	if s.index < 0 {
		return 0
	}
	res := s.data[s.index]
	s.index--
	return res
}

func (s *stack) clear() {
	s.index = -1
}

func (s *stack) clearNum(num int) {
	s.index = s.index - num
}

func main() {
	s := newStack(2)
	s.push(11) // s.data[0] == 11
	fmt.Printf("s: %v\n", s)
	fmt.Printf("s.pop: %v\n", s.pop()) // s.data[0] == 11 index == -1

	s.push(12) // s.index == 0, s.data[0] == 12
	s.push(13) // s.index == 1, s.data[1] == 13
	s.push(14) // s.index == 2, s.data[2] == 14 (append)
	fmt.Printf("s: %v\n", s)
	s.clear()
	fmt.Printf("s: %v\n", s)
	s.push(15)
	s.push(16)
	s.push(17)
	fmt.Printf("s: %v\n", s)
	s.clearNum(2) // s.pop(), s.pop()
	s.pop()
	s.push(18)
	fmt.Printf("s: %v\n", s)
}
