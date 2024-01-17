package main

import "fmt"

/*
Здесь используется та же идея: создание программы посещения
инкапсулируется в объекте (назовем его «строителем»), а клиент использует этот объект для констру-
ирования структуры данных.
Требуется для реализации:

1. Класс Director, который будет распоряжаться строителем и отдавать ему команды в
	нужном порядке, а строитель будет их выполнять;
2. Базовый абстрактный класс Builder, который описывает интерфейс строителя,
	те команды, которые он обязан выполнять;
3. Класс ConcreteBuilder, который реализует интерфейс строителя и взаимодействует
	со сложным объектом;
4. Класс сложного объекта Product.

Плюсы:
- Инкапсуляция процесса создания сложного
	объекта.
- Возможность поэтапного конструирования
	объекта с переменным набором этапов (в отличие
	от «одноэтапных» фабрик).
- Сокрытие внутреннего представления продукта от
	клиента.
- Реализации продуктов могут свободно
	изменяться, потому что клиент имеет дело только
	с абстрактным интерфейсом
Минусы:
-ConcreteBuilder и создаваемый им продукт жестко связаны между собой,
	поэтому при внесеннии изменений в класс продукта скорее всего придется
	соотвествующим образом изменять и класс ConcreteBuilder.
Пример использования:
-В системе могут существовать сложные объекты, создание которых за одну
	операцию затруднительно или невозможно. Требуется поэтапное построение объектов
	с контролем результатов выполнения каждого этапа.
*/

type Director struct {
	b Builder
}

func NewDirector() *Director {
	return &Director{b: &ConcreteBuilder{p: new(Person)}}
}

func (d *Director) Build(name, surname, patronymic string, age int) *Person {
	d.b.SetName(name)
	d.b.SetSurname(surname)
	d.b.SetPatronymic(patronymic)
	d.b.SetAge(age)
	return d.b.GetPerson()
}

type Builder interface {
	SetName(name string)
	SetSurname(surname string)
	SetPatronymic(patronymic string)
	SetAge(age int)
	GetPerson() *Person
}

type ConcreteBuilder struct {
	p *Person
}

func (c *ConcreteBuilder) SetAge(age int) {
	c.p.Age = age
}

func (c *ConcreteBuilder) SetName(name string) {
	c.p.Name = name
}

func (c *ConcreteBuilder) SetSurname(surname string) {
	c.p.Surname = surname
}

func (c *ConcreteBuilder) SetPatronymic(patronymic string) {
	c.p.Patronymic = patronymic
}

func (c *ConcreteBuilder) GetPerson() *Person {
	return c.p
}

type Person struct {
	Name, Surname, Patronymic string
	Age                       int
}

func main() {
	d := NewDirector()
	fmt.Println(d.Build("Nick", "Ivanov", "Petrovitch", 25))
}
