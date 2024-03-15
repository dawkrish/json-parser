package main

import (
	"testing"
)

type testCaseTokenizer struct {
	input    string
	expected [][]Token
}

func assertCorrectMessageForTokenizer(t testing.TB, got, want [][]Token) {
	t.Helper()
	if !compareTokens(got, want) {
		t.Errorf("want <%q> but got <%q>", want, got)
	}
	//fmt.Println("Successful Result => ", got)
}

func TestNumber(t *testing.T) {
	testCases := []testCaseTokenizer{
		{input: "1", expected: [][]Token{{{Type: NUMBER, Value: "1"}}}},
		{input: "2", expected: [][]Token{{{Type: NUMBER, Value: "2"}}}},
		{input: "99", expected: [][]Token{{{Type: NUMBER, Value: "99"}}}},
		{input: "00", expected: [][]Token{{{Type: NUMBER, Value: "00"}}}},
		{input: "00,", expected: [][]Token{{{Type: NUMBER, Value: "00"}}}},
		{input: "00\"", expected: [][]Token{{{Type: NUMBER, Value: "00"}}}},
		{input: "4.4", expected: [][]Token{{{Type: NUMBER, Value: "4.4"}}}},
		{input: "4e4", expected: [][]Token{{{Type: NUMBER, Value: "4e4"}}}},
		{input: "4e44", expected: [][]Token{{{Type: NUMBER, Value: "4e44"}}}},
		{input: "44e4", expected: [][]Token{{{Type: NUMBER, Value: "44e4"}}}},
		{input: "4.", expected: [][]Token{}},
		{input: "4e", expected: [][]Token{}},
		{input: "4.a", expected: [][]Token{}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := Tokenize(tc.input)
			if err != nil {
				t.Fatalf("unexpected error  : %v", err)
			}
			assertCorrectMessageForTokenizer(t, got, tc.expected)
		})
	}
}

func TestString(t *testing.T) {
	testCases := []testCaseTokenizer{
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
			assertCorrectMessageForTokenizer(t, got, tc.expected)
		})
	}
}

func TestSymbol(t *testing.T) {
	testCases := []testCaseTokenizer{
		{input: "348mptqw 1", expected: [][]Token{}},
		{input: "true 1", expected: [][]Token{{Token{Type: SYMBOL, Value: "true"}, Token{Type: WHITE_SPACE, Value: " "}, Token{Type: NUMBER, Value: "1"}}}},
		{input: "348true ", expected: [][]Token{{Token{Type: NUMBER, Value: "348"}, Token{Type: SYMBOL, Value: "true"}}}},
		{input: "348true", expected: [][]Token{{Token{Type: NUMBER, Value: "348"}, Token{Type: SYMBOL, Value: "true"}}}},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := Tokenize(tc.input)
			if err != nil {
				t.Fatalf("unexpected error  : %v", err)
			}
			assertCorrectMessageForTokenizer(t, got, tc.expected)
		})
	}
}

func TestPunctuationsAndUndefined(t *testing.T) {
	testCases := []testCaseTokenizer{
		{input: ",", expected: [][]Token{{Token{Type: COMMA, Value: ","}}}},
		{input: ":", expected: [][]Token{{Token{Type: COLON, Value: ":"}}}},
		{input: "[", expected: [][]Token{{Token{Type: LEFT_SQUARE_BRACE, Value: "["}}}},
		{input: "}", expected: [][]Token{{Token{Type: RIGHT_CURLY_BRACE, Value: "}"}}}},
		{input: "{", expected: [][]Token{{Token{Type: LEFT_CURLY_BRACE, Value: "{"}}}},
		{input: "]", expected: [][]Token{{Token{Type: RIGHT_SQUARE_BRACE, Value: "]"}}}},
		{input: "\".\"", expected: [][]Token{{Token{Type: STRING, Value: "."}}}},
		{input: ".", expected: [][]Token{}},
		{input: "''", expected: [][]Token{}},
		{input: "<", expected: [][]Token{}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := Tokenize(tc.input)
			if err != nil {
				t.Fatalf("unexpected error  : %v", err)
			}
			assertCorrectMessageForTokenizer(t, got, tc.expected)
		})
	}
}

// fileContent, err := os.ReadFile(TEST_DIRECTORY + "/" + "one.json")
// if err != nil {
// 	t.Fatal(err)
// }
