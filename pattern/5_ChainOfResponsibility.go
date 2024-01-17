package main

import (
	"fmt"
)

/*
Цепочка обязанностей — это поведенческий паттерн проектирования,
который позволяет передавать запросы последовательно по цепочке обработчиков.
Каждый последующий обработчик решает, может ли он обработать запрос сам и
стоит ли передавать запрос дальше по цепи.

Плюсы
- уменьшает зависимость между клиентом и обработчиками. Каждый обработчик
сам выполняет свою логику независимо.
- реализует принцип единственной обязанности.
- реализует принцип открытости и закрытости.

Минусы
- запрос может остаться необработанным.

Примеры использования:
- в разрабатываемой системе имеется группа объектов, которые могут обрабатывать
сообщения определенного типа;
- все сообщения должны быть обработаны хотя бы одним объектом системы;
- сообщения в системе обрабатываются по схеме «обработай сам либо перешли другому»,
то есть одни сообщения обрабатываются на том уровне, где они получены,
а другие пересылаются объектам иного уровня.
*/

type Handler interface {
	Handle(com Command) int
	SetNext(h Handler)
}

type Adder struct {
	next Handler
}

func (a *Adder) SetNext(h Handler) {
	a.next = h
}

func (a *Adder) Handle(com Command) int {
	if com.handler == "adder" {
		return com.num1 + com.num2
	}
	if a.next == nil {
		return com.num1
	}
	return a.next.Handle(com)
}

type Multiplier struct {
	next Handler
}

func (m *Multiplier) SetNext(h Handler) {
	m.next = h
}

func (m *Multiplier) Handle(com Command) int {
	if com.handler == "multiplier" {
		return com.num1 * com.num2
	}
	if m.next == nil {
		return com.num1
	}
	return m.next.Handle(com)
}

type Subtracter struct {
	next Handler
}

func (s *Subtracter) SetNext(h Handler) {
	s.next = h
}

func (s *Subtracter) Handle(com Command) int {
	if com.handler == "subtracter" {
		return com.num1 - com.num2
	}
	if s.next == nil {
		return com.num1
	}
	return s.next.Handle(com)
}

type Command struct {
	handler    string
	num1, num2 int
}

func main() {
	a := Adder{}
	m := Multiplier{}
	m.SetNext(&Subtracter{})
	a.SetNext(&m)
	fmt.Println(a.Handle(Command{"adder", 5, 1}))
	fmt.Println(a.Handle(Command{"multiplier", 5, 12}))
	fmt.Println(a.Handle(Command{"subtracter", 5, 10}))
}
