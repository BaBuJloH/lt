package main

import (
	"reflect"
	"testing"
)

func TestMakeAnagramMap(t *testing.T) {
	test := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	val_answ := MakeAnagramMap(test)
	val_answ_true := map[string][]string{"листок": []string{"слиток", "столик"}, "пятак": []string{"пятка", "тяпка"}}
	if !reflect.DeepEqual(val_answ, val_answ_true) {
		t.Fatalf("Test case: %v,\n function returns: %v,\n true answ: %v\n", test, val_answ, val_answ_true)
	}

	test1 := []string{}
	val_answ1 := MakeAnagramMap(test1)
	val_answ_true1 := map[string][]string{}
	if !reflect.DeepEqual(val_answ1, val_answ_true1) {
		t.Fatalf("Test case: %v,\n function returns: %v,\n true answ: %v\n", test1, val_answ1, val_answ_true1)
	}

	test3 := []string{"пятак", "пятка", "тяпка", "пятак", "пятак", "пятак", "пятак"}
	val_answ3 := MakeAnagramMap(test3)
	val_answ_true3 := map[string][]string{"пятак": []string{"пятка", "тяпка"}}
	if !reflect.DeepEqual(val_answ3, val_answ_true3) {
		t.Fatalf("Test case: %v,\n function returns: %v,\n true answ: %v\n", test3, val_answ3, val_answ_true3)
	}

	test4 := []string{"пятАк", "Пятка", "тяпКа", "листок", "слиток", "столик"}
	val_answ4 := MakeAnagramMap(test4)
	val_answ_true4 := map[string][]string{"листок": []string{"слиток", "столик"}, "пятак": []string{"пятка", "тяпка"}}
	if !reflect.DeepEqual(val_answ4, val_answ_true4) {
		t.Fatalf("Test case: %v,\n function returns: %v,\n true answ: %v\n", test4, val_answ4, val_answ_true4)
	}
}
