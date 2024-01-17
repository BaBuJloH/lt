package main

import "fmt"

/*
Состояние (англ. State) — поведенческий шаблон проектирования.
Используется в тех случаях, когда во время выполнения программы объект
должен менять своё поведение в зависимости от своего состояния.
Создается впечатление, что объект изменил свой класс.
Паттерн State является объектно-ориентированной реализацией конечного автомата.

- Вводит класс Context, в котором определяется интерфейс для внешнего мира.
- Вводит абстрактный класс State.
- Представляет различные "состояния" конечного автомата в виде подклассов State.
- В классе Context имеется указатель на текущее состояние, который изменяется
при изменении состояния конечного автомата.

Пример использования:
- Когда у вас есть объект, поведение которого кардинально меняется в зависимости
от внутреннего состояния, причём типов состояний много, и их код часто меняется.
*/

type TVChannel interface {
	WriteAMessage()
}

type TV struct {
	TVChannel
}

func NewTV() *TV {
	return &TV{TV1{}}
}

func (tv *TV) SetChannel(tvchannel TVChannel) {
	tv.TVChannel = tvchannel
}

func (tv *TV) NextChannel() {
	switch tv.TVChannel.(type) {
	case TV1:
		tv.TVChannel = Russia1{}
	case Russia1:
		tv.TVChannel = Friday{}
	case Friday:
		tv.TVChannel = TV1{}
	}
}

type TV1 struct {
}

func (tv1 TV1) WriteAMessage() {
	fmt.Println("You are at TV1 channel now")
}

type Russia1 struct {
}

func (ru1 Russia1) WriteAMessage() {
	fmt.Println("You are at Russia 1 channel now")
}

type Friday struct {
}

func (fr Friday) WriteAMessage() {
	fmt.Println("You are at Friday channel now")
}

func main() {
	tv := NewTV()
	tv.SetChannel(Friday{})
	tv.WriteAMessage()
	tv.NextChannel()
	tv.WriteAMessage()
	tv.NextChannel()
	tv.WriteAMessage()
}
