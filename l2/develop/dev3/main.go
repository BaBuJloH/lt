/*
Утилита sort
Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительно

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type SortConfig struct {
	sortColumn         int
	sortByNumberValue  bool
	reverseSort        bool
	uniqueRows         bool
	sortByMonth        bool
	isRowsAlredySorted bool
	months             [12]string
	fileName           string
}

func NewSortConfig() *SortConfig {
	s := SortConfig{}
	s.months = [12]string{"янв", "фев", "мар", "апр", "май", "июн", "июл", "авг", "сен", "окт", "ноя", "дек"}
	flag.IntVar(&s.sortColumn, "k", 0, "Sets column for sort")
	flagN := flag.Bool("n", false, "Make sort by number value")
	flagR := flag.Bool("r", false, "Makes reverse sort")
	flagU := flag.Bool("u", false, "Ignore duplicate lines")
	flagM := flag.Bool("m", false, "Makes sort by month")
	flagC := flag.Bool("c", false, "Check if rows alredy sorted")

	flag.Parse()

	args := flag.Args()
	s.sortByNumberValue = *flagN
	s.reverseSort = *flagR
	s.uniqueRows = *flagU
	s.sortByMonth = *flagM
	s.isRowsAlredySorted = *flagC

	if len(args) == 1 {
		s.fileName = args[0]
	} else {
		log.Fatalf("The argument (path to the file name) must be one")
	}
	return &s
}

func Start(s *SortConfig) (string, error) {
	rows, err := readFile(s.fileName)
	if err != nil {
		return "", fmt.Errorf("can't read file '%s': %s", s.fileName, err.Error())
	}
	return sortRows(rows, s)
}

func readFile(fileName string) ([]string, error) {
	rows := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		return rows, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		rows = append(rows, sc.Text())
	}
	return rows, nil
}

func uniqualizer(rows []string) []string {
	tempBuf := make(map[string]bool)
	for _, row := range rows {
		tempBuf[row] = true
	}

	rows = rows[:0]
	for key := range tempBuf {
		rows = append(rows, key)
	}
	return rows
}

// поиск пробельных символов
func getColunmValue(row string, s *SortConfig) (string, error) {
	re := regexp.MustCompile(`\s+`)

	listOfColumns := re.Split(strings.TrimSpace(row), -1)
	if len(listOfColumns) >= s.sortColumn {
		return listOfColumns[s.sortColumn-1], nil
	}
	return "", fmt.Errorf("can not find column")
}

func sortRows(rows []string, s *SortConfig) (string, error) {
	var sourceRows []string
	if s.isRowsAlredySorted {
		sourceRows = make([]string, len(rows))
		_ = copy(sourceRows, rows)

		s.sortByMonth = false
		s.sortByNumberValue = false
		s.sortColumn = 0
		s.isRowsAlredySorted = false
		s.reverseSort = false
	}

	switch {

	case s.sortColumn > 0:
		{
			if s.uniqueRows {
				rows = uniqualizer(rows)
			}

			sort.SliceStable(rows, func(i, j int) bool {
				iVal, err := getColunmValue(rows[i], s)
				if err != nil {
					return false
				}
				jVal, err := getColunmValue(rows[j], s)
				if err != nil {
					return false
				}

				if s.sortByNumberValue {
					if s.reverseSort {
						return !iValNumLessThanjVal(iVal, jVal)
					}
					return iValNumLessThanjVal(iVal, jVal)
				}

				if s.sortByMonth {
					if s.reverseSort {
						return !iValMonthLessThanjVal(iVal, jVal, s)
					}
					return iValMonthLessThanjVal(iVal, jVal, s)
				}

				if s.reverseSort {
					return iVal < jVal
				}
				return iVal > jVal
			})
		}

	case s.sortByNumberValue:
		{
			if s.uniqueRows {
				rows := uniqualizer(rows)
			}

			sort.SliceStable(rows, func(i, j int) bool {
				if s.reverseSort {
					return !iValNumLessThanjVal(rows[i], rows[j])
				}
				return iValNumLessThanjVal(rows[i], rows[j])
			})
		}

	case s.uniqueRows:
		{
			rows = uniqualizer(rows)
			sort.SliceStable(rows, func(i, j int) bool {
				if s.sortByMonth {
					if s.reverseSort {
						return !iValNumLessThanjVal(rows[i], rows[j], s)
					}
					return iValNumLessThanjVal(rows[i], rows[j], s)
				}

				if s.reverseSort {
					return rows[i] > rows[j]
				}
				return rows[i] < rows[j]
			})
		}

	case s.sortByMonth:
		{
			sort.SliceStable(rows, func(i, j int) bool {
				if s.reverseSort {
					return !iValMonthLessThatjVal(rows[i], rows[j], s)
				}
				return iValMonthLessThanjVal(rows[i], rows[j], s)
			})
		}

	default:
		sort.SliceStable(rows, func(i, j int) bool {
			if s.reverseSort {
				return rows[i] > rows[j]
			}
			return rows[i] < rows[j]
		})
	}

	if s.isRowsAlredySorted {
		for i, row := range rows {
			if row != sourceRows[i] {
				return "false", nil
			}
		}
		return "true", nil
	}

	var result strings.Builder
	lenRows := len(rows)
	for i, row := range rows {
		if i < lenRows-1 {
			_, _ = result.WriteString(row + "\n")
		} else {
			_, _ = result.WriteString(row)
		}
	}
	return result.String(), nil
}

func iValNumLessThanjVal(strI, strJ string) bool {
	iVal, _ := strconv.Atoi(strI)
	jVal, _ := strconv.Atoi(strJ)
	return iVal < jVal
}

func iValMonthLessThanjVal(strI, strJ string, s *SortConfig) bool {
	for _, month := range s.months {
		switch {
		case month == strI:
			return true
		case month == strJ:
			return false
		default:
			continue
		}
	}
	return true
}

func main() {
	s := NewSortConfig()
	res, err := Start(s)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(res)
}
