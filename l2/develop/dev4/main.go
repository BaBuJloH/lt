/*Написать функцию поиска всех множеств анаграмм по словарю.


Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.


Требования:
Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
Выходные данные: ссылка на мапу множеств анаграмм
Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func MakeAnagramMap(words []string) map[string][]string {
	data := make(map[string][]string)
	wordsCollection := make(map[string]interface{})

	for _, word := range words {
		word = strings.ToLower(word)
		_, ok := wordsCollection[word]
		if ok {
			continue
		}
		wordsCollection[word] = struct{}{}
		stringRunes := []rune(word)
		sort.Slice(stringRunes, func(i int, j int) bool { return stringRunes[i] < stringRunes[j] })
		sortedString := string(stringRunes)
		data[sortedString] = append(data[sortedString], word)
	}

	anagramMap := make(map[string][]string)
	for _, val := range data {
		anagramMap[val[0]] = val[1:]
	}
	return anagramMap
}

func main() {
	fmt.Printf("%v", MakeAnagramMap([]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}))
}
