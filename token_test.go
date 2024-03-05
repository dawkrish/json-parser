package main

import (
	"testing"
)

type testCase struct {
	input    string
	expected [][]Token
}

func assertCorrectMessage(t testing.TB, got, want [][]Token) {
	t.Helper()
	if !compareTokens(got, want) {
		t.Errorf("want <%q> but got <%q>", want, got)
	}
	//fmt.Println("Successful Result => ", got)
}

func TestNumber(t *testing.T) {

	testCases := []testCase{
		{input: "1", expected: [][]Token{{{Type: NUMBER, Value: "1"}}}},
		{input: "2", expected: [][]Token{{{Type: NUMBER, Value: "2"}}}},
		{input: "99", expected: [][]Token{{{Type: NUMBER, Value: "99"}}}},
		{input: "00", expected: [][]Token{{{Type: NUMBER, Value: "00"}}}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := Tokenize(tc.input)
			if err != nil {
				t.Fatalf("unexpected error  : %v", err)
			}
			assertCorrectMessage(t, got, tc.expected)
		})
	}
}

func TestString(t *testing.T) {
	testCases := []testCase{
		{input: "\"", expected: [][]Token{}},
		{input: "\"\"", expected: [][]Token{{{Type: STRING, Value: ""}}}},
		{input: "\"hello 123 .-?\"", expected: [][]Token{{{Type: STRING, Value: "hello 123 .-?"}}}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := Tokenize(tc.input)
			if err != nil {
				t.Fatalf("unexpected error  : %v", err)
			}
			assertCorrectMessage(t, got, tc.expected)
		})
	}
}

// fileContent, err := os.ReadFile(TEST_DIRECTORY + "/" + "one.json")
// if err != nil {
// 	t.Fatal(err)
// }
