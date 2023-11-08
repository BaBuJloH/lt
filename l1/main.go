package main

import "fmt"

// суть которого состоит в том, что классы должны зависеть от абстракций, а не от конкретных деталей
// суть которого состоит в том, что код должен зависеть от абстракций, а не от конкретных деталей
type Stack interface { 
    Pop() int
    Push(int)
}

func main() {
    var s Stack
    s = New(2) 
    // s = stack.New(2)
    // s = superstack.New(10)

    s.Push(11) // s.data[0] == 11
    fmt.Printf("s: %v\n", s)
    fmt.Printf("s.pop: %v\n", s.Pop()) // s.data[0] == 11 index == -1
}


// package superstack

// package stack
// stack

type Us

type stack struct {
    index int
    data []int
}

func New(num int) *stack {
    s := stack{}
    s.index = -1
    s.data = make([]int, num)
    return &s
}

func (s *stack) Clear() {}

func (s *stack) Push(i int) {
    s.index++
    if s.index >= len(s.data) {
        s.data = append(s.data, i)
        return 
    }
    s.data[s.index] = i
}

func (s *stack) Pop() int {
    if s.index < 0 {
        return 0
    }
    res := s.data[s.index]
    s.index--
    return res
}