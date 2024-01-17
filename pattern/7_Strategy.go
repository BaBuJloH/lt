package main

import "fmt"

/*
Паттерн Strategy переносит в отдельную иерархию классов все детали, связанные с реализацией алгоритмов.

Плюсы:
- Систему проще поддерживать и модифицировать, так как семейство алгоритмов перенесено
в отдельную иерархию классов.
- Паттерн Strategy предоставляет возможность замены одного алгоритма другим в процессе
выполнения программы.
- Паттерн Strategy позволяет скрыть детали реализации алгоритмов от клиента.

Минусы:
- Для правильной настройки системы пользователь должен знать об особенностях всех алгоритмов.
- Число классов в системе, построенной с применением паттерна Strategy, возрастает.

Пример использования:
- Отделение процедуры выбора алгоритма от его реализации.
Это позволяет сделать выбор на основании контекста.
*/

type Strategy interface {
	CountCurrency(value float64) float64
}

type Client struct {
	Strategy
}

func (c *Client) SetStrategy(str string) {
	if str == "eur to rub" {
		c.Strategy = &EurToRub{}
	} else if str == "usd to rub" {
		c.Strategy = &UsdToRub{}
	}
}

type EurToRub struct {
}

func (etr *EurToRub) CountCurrency(num float64) float64 {
	return num * 62.73
}

type UsdToRub struct {
}

func (utr *UsdToRub) CountCurrency(num float64) float64 {
	return num * 61.12
}

func main() {
	client := Client{}
	client.SetStrategy("usd to rub")
	fmt.Println(client.CountCurrency(1000))
}
