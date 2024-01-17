package main

import "fmt"

/*
Паттерн Command позволяет представить запрос в виде объекта. Из этого следует, что команда
- это объект. Такие запросы, например, можно ставить в очередь, отменять или возобновлять.

Требуется для реализации:
- Базовый абстрактный класс Command описывающий интерфейс команды;
- Класс ConcreteCommand, реализующий команду;
- Класс Receiver, реализующий получателя и имеющий набор действий,
которые команда можем запрашивать;

Сначала клиент создает объект ConcreteCommand, конфигурируя его получателем запроса.
Этот объект также доступен инициатору. Инициатор использует его при отправке запроса,
вызывая метод execute(). Этот алгоритм напоминает работу функции обратного вызова в
процедурном программировании – функция регистрируется, чтобы быть вызванной позднее.

Паттерн Command отделяет объект, инициирующий операцию, от объекта, который знает,
как ее выполнить. Единственное, что должен знать инициатор, это как отправить команду.
Это придает системе гибкость: позволяет осуществлять динамическую замену команд, использовать
сложные составные команды, осуществлять отмену операций.

*/

type Command interface {
	Execute() string
}

type ConcreteCommand1 struct {
	rec *Receiver
}

func (cc *ConcreteCommand1) Execute() string {
	return cc.rec.Action1()
}

type ConcreteCommand2 struct {
	rec *Receiver
}

func (cc *ConcreteCommand2) Execute() string {
	return cc.rec.Action2()
}

type Receiver struct {
}

func (r *Receiver) Action1() string {
	return "action1"
}

func (r *Receiver) Action2() string {
	return "action2"
}

func main() {
	// make queue of commands
	receivers := []Command{&ConcreteCommand1{}, &ConcreteCommand2{}}
	for _, rec := range receivers {
		fmt.Println(rec.Execute())
	}
}
