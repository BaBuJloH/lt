/*Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.*/

package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	text := "a4bc2d5e"
	unpacked, err := UnpackString(text)

	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("\nInput: '%s'\nResult: '%s'", text, unpacked)
}

func UnpackString(text string) (string, error) {

	var unpackedString strings.Builder
	var lastRune rune
	var isEscapeSymbol bool

	for _, symbol := range text {
		switch {
		case isEscapeSymbol:
			{
				isEscapeSymbol = false
				lastRune = symbol
			}
		case symbol <= '9' && symbol >= '0':
			{
				if lastRune != 0 {
					iterationCount := int(symbol - '0')
					for i := 0; i < iterationCount; i++ {
						unpackedString.WriteRune(lastRune)
					}
					lastRune = 0
				} else {
					return "", fmt.Errorf("wrong string")
				}
			}
		case symbol == '\\':
			{
				isEscapeSymbol = true
				if lastRune != 0 {
					unpackedString.WriteRune(lastRune)
				}
			}

		default:
			if lastRune != 0 {
				unpackedString.WriteRune(lastRune)
			}
			lastRune = symbol
		}
	}
	if lastRune != 0 {

		unpackedString.WriteRune(lastRune)
	}
	return unpackedString.String(), nil

}
