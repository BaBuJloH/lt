package main

import (
	"testing"
)

func TestUnpackageOfString(t *testing.T) {
	str := "45"
	val, err := UnpackageOfString(str)
	if err == nil {
		t.Fatalf("Error: %v, string example: %s, result: %s", err, str, val)
	}
	str1 := "a4bc2d5e"
	val1, err1 := UnpackageOfString(str1)
	if err1 != nil || val1 != "aaaabccddddde" {
		t.Fatalf("Error: %v, string example: %s, result: %s", err1, str1, val1)
	}
	str2 := "abcd"
	val2, err2 := UnpackageOfString(str2)
	if err2 != nil || val2 != "abcd" {
		t.Fatalf("Error: %v, string example: %s, result: %s", err2, str2, val2)
	}
	str3 := ""
	val3, err3 := UnpackageOfString(str3)
	if err3 != nil || val3 != "" {
		t.Fatalf("Error: %v, string example: %s, result: %s", err3, str3, val3)
	}
	str4 := "qwe\\4\\5"
	val4, err4 := UnpackageOfString(str4)
	if err4 != nil || val4 != "qwe45" {
		t.Fatalf("Error: %v, string example: %s, result: %s", err4, str4, val4)
	}
	str5 := "qwe\\\\5"
	val5, err5 := UnpackageOfString(str5)
	if err5 != nil || val5 != "qwe\\\\\\\\\\" {
		t.Fatalf("Error: %v, string example: %s, result: %s", err5, str5, val5)
	}
	str6 := "qwe\\45"
	val6, err6 := UnpackageOfString(str6)
	if err6 != nil || val6 != "qwe44444" {
		t.Fatalf("Error: %v, string example: %s, result: %s", err6, str6, val6)
	}
}
