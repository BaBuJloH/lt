package main

import (
	"fmt"
	"log"
)

/*
Порождающий шаблон проектирования, предоставляющий подклассам (дочерним классам) интерфейс
для создания экземпляров некоторого класса. В момент создания наследники могут определить,
какой класс создавать. Иными словами, данный шаблон делегирует создание объектов наследникам
родительского класса. Это позволяет использовать в коде программы не конкретные классы,
а манипулировать абстрактными объектами на более высоком уровне.

Структура:
- Product — продукт
	определяет интерфейс объектов, создаваемых абстрактным методом;
- ConcreteProduct — конкретный продукт
	реализует интерфейс Product;
- Creator — создатель
	объявляет фабричный метод, который возвращает объект типа Product. Может также содержать реализацию этого метода «по умолчанию»;
	может вызывать фабричный метод для создания объекта типа Product;
- ConcreteCreator — конкретный создатель
	переопределяет фабричный метод таким образом, чтобы он создавал и возвращал объект класса ConcreteProduct.

Плюсы:
- позволяет сделать код создания объектов более универсальным,
	не привязываясь к конкретным классам (ConcreteProduct), а оперируя лишь общим интерфейсом (Product);
	позволяет установить связь между параллельными иерархиями классов.
Минусы:
- Может привести к созданию больших параллельных иерархий классов,
	так как для каждого класса продукта надо создать свой подкласс создателя.
Примеры использования:
- Система должна оставаться расширяемой путем добавления объектов новых типов.
	Непосредственное использование выражения new является нежелательным, так как в этом случае код
	создания объектов с указанием конкретных типов может получиться разбросанным по всему приложению.
	Тогда такие операции как добавление в систему объектов новых типов или замена объектов одного
	типа на другой будут затруднительными (подробнее в разделе Порождающие паттерны).
	Паттерн Factory Method позволяет системе оставаться независимой как от самого процесса
	порождения объектов, так и от их типов.
- Заранее известно, когда нужно создавать объект, но неизвестен его тип.
*/

type Product interface {
	Name() string
}

// concrete product 1
type Pineapple struct {
}

func NewPinewapple() Product {
	return &Pineapple{}
}

func (a *Pineapple) Name() string {
	return "pineapple"
}

// concrete product 2
type Apple struct {
}

func NewApple() Product {
	return &Apple{}
}

func (a *Apple) Name() string {
	return "apple"
}

// concrete product 3
type Cherry struct {
}

func NewCherry() Product {
	return &Cherry{}
}

func (c *Cherry) Name() string {
	return "cherry"
}

type Creator interface {
	CreateProduct(str string) Product
}

type ConcreteCreator struct {
}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

func (c *ConcreteCreator) CreateProduct(str string) Product {
	var product Product
	switch str {
	case "apple":
		product = &Apple{}
	case "pineapple":
		product = &Pineapple{}
	case "cherry":
		product = &Cherry{}
	default:
		log.Fatal("unknown product")
	}
	return product
}

func main() {
	c := NewCreator()
	fmt.Println(c.CreateProduct("apple").Name())
	fmt.Println(c.CreateProduct("cherry").Name())
	fmt.Println(c.CreateProduct("pineapple").Name())
	fmt.Println(c.CreateProduct("strawberry").Name())
}
