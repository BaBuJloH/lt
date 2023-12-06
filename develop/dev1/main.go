/*Создать программу печатающую точное время с использованием NTP -библиотеки. Инициализировать как go module. Использовать
 библиотеку github.com/beevik/ntp. Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Требования:
Программа должна быть оформлена как go module
Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS*/

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func ExactTime(host string) (time.Time, error) {
	return ntp.Time(host)
}

func main() {

	time, err := ExactTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		_, err := fmt.Fprint(os.Stderr, err.Error())
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
	fmt.Println(time)
}
