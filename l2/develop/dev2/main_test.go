package main

import "testing"

func TestUnpackString(t *testing.T) {
	testTable := []struct {
		name      string
		input     string
		output    string
		haveError bool
	}{
		{name: "simple unpack", input: "a4bc2d5e", output: "aaaabccddddde"},
		{name: "no unpack operations", input: "abcd", output: "abcd"},
		{name: "incorrect string", input: "45", output: "", haveError: true},
		{name: "empty string", input: "", output: ""},
		{name: "escape - последовательность 1", input: "qwe\\4\\5", output: "qwe45"},
		{name: "escape - последовательность 2", input: "qwe\\45", output: "qwe44444"},
		{name: "escape - последовательность 3", input: "qwe\\\\5", output: "qwe\\\\\\\\\\"},
	}

	for _, testingCase := range testTable {
		t.Run(testingCase.name, func(t *testing.T) {
			result, err := UnpackString(testingCase.input)
			if !testingCase.haveError {
				if err != nil {
					t.Errorf("expected err == nil; got '%s'", err.Error())
				}
				if result != testingCase.output {
					t.Errorf("expected result '%s'; got '%s'", testingCase.output, result)
				}

			} else {
				if err != nil {
					if err.Error() != "wrong string" {
						t.Errorf("expected err.Error() == 'wrong string'; got '%s'", err.Error())
					}

				} else {
					t.Error("expected err.Error() == 'wrong string', err == nil")
				}
			}
		})
	}
}
