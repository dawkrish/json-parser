package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	NUMBER             = "NUMBER"
	STRING             = "STRING"
	BOOLEAN            = "BOOLEAN"
	COMMA              = "COMMA"
	DOUBLE_QUOTE       = "DOUBLE_QUOTE"
	SEMI_COLON         = "SEMI_COLON"
	LEFT_CURLY_BRACE   = "LEFT_CURLY_BRACE"
	RIGHT_CURLY_BRACE  = "RIGHT_CURLY_BRACE"
	LEFT_SQUARE_BRACE  = "LEFT_SQUARE_BRACE"
	RIGHT_SQUARE_BRACE = "RIGHT_SQUARE_BRACE"

	LETTER    = "LETTER"
	UNDEFINED = "UNDEFINED"
)

type Token struct {
	Type  string
	Value string
}

func Tokenize(input string) ([][]Token, error) {
	var tokens [][]Token
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var tr []Token

		for i := 0; i < len(line); i++ {
			pseudoToken := getTokenType(line[i])
			if pseudoToken.Type == NUMBER {
				possibleNumber := pseudoToken.Value
				// keep checking for the next character till we find a NOT-NUMBER
				for i+1 < len(line) {
					i++
					pt := getTokenType(line[i])
					if pt.Type != NUMBER {
						break
					}
					possibleNumber += pt.Value

				}
				// now we have the maximum possibleNumber
				tr = append(tr, Token{Type: NUMBER, Value: possibleNumber})
				break
			}

			if pseudoToken.Type == DOUBLE_QUOTE {
				isStringClosed := false
				possibleString := ""
				for i+1 < len(line) {
					i++
					pt := getTokenType(line[i])
					if pt.Type == DOUBLE_QUOTE {
						isStringClosed = true
						break
					}
					possibleString += pt.Value
				}
				if !isStringClosed {
					return [][]Token{}, errors.New("missing closing quote \"")
				}
				tr = append(tr, Token{Type: STRING, Value: possibleString})
			}

		}
		tokens = append(tokens, tr)
	}
	return tokens, nil
}

func getTokenType(c uint8) Token {
	val := fmt.Sprintf("%c", c)
	if c >= 48 && c <= 57 {
		return Token{Type: NUMBER, Value: val}
	}
	if c >= 65 && c <= 95 || c >= 97 && c <= 122 {
		return Token{Type: LETTER, Value: val}
	}
	if c == 34 {
		return Token{Type: DOUBLE_QUOTE, Value: val}
	}
	return Token{Type: UNDEFINED, Value: val}
}

func compareTokens(t1, t2 [][]Token) bool {
	if len(t1) != len(t2) {
		return false
	}
	for i := 0; i < len(t1); i++ {
		if len(t1[i]) != len(t2[i]) {
			return false
		}
		for j := 0; j < len(t1[i]); j++ {
			if t1[i][j].Type != t2[i][j].Type {
				return false
			}
			if t1[i][j].Value != t2[i][j].Value {
				return false
			}
		}
	}
	return true
}

//for _, c := range line {
//	//fmt.Printf("%T\t%c\n", c, c)
//	if c >= 48 && c <= 57 {
//		fmt.Println("c is a number")
//		// now check for the next token, if its a number then add it and keep doing it...

//	}
//}

//fmt.Printf("%T\t%c\n", line[i], line[i])
