package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	NUMBER = "NUMBER"
	STRING = "STRING"

	COMMA              = "COMMA"
	DOUBLE_QUOTE       = "DOUBLE_QUOTE"
	COLON              = "COLON"
	LEFT_CURLY_BRACE   = "LEFT_CURLY_BRACE"
	RIGHT_CURLY_BRACE  = "RIGHT_CURLY_BRACE"
	LEFT_SQUARE_BRACE  = "LEFT_SQUARE_BRACE"
	RIGHT_SQUARE_BRACE = "RIGHT_SQUARE_BRACE"
	WHITE_SPACE        = "WHITE_SPACE"
	SYMBOL             = "SYMBOL"
	LETTER             = "LETTER"

	CARRIAGE_RETURN = "CARRIAGE_RETURN"
	UNDEFINED       = "UNDEFINED"
)

type Token struct {
	Type  string
	Value string
}

func (t Token) String() string {
	return fmt.Sprintf("Type-%v", t.Type)
}

func Tokenize(input string) ([][]Token, error) {
	//fmt.Println("Starting Tokenizing with input -> ", input)
	var tokens [][]Token
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.Trim(line, " ")
		// fmt.Println("line -> ", line)
		var tr []Token

		for i := 0; i < len(line); i++ {
			pseudoToken := getTokenType(line[i])

			if pseudoToken.Type == DOUBLE_QUOTE {
				isStringClosed := false
				possibleString := ""

				for i+1 < len(line) {
					i++
					pseudoToken = getTokenType(line[i])
					if pseudoToken.Type == DOUBLE_QUOTE {
						isStringClosed = true
						break
					}
					possibleString += pseudoToken.Value
				}
				if !isStringClosed {
					return [][]Token{}, errors.New("missing closing quote \"")
				}
				tr = append(tr, Token{Type: STRING, Value: possibleString})
			}

			if pseudoToken.Type == LETTER {
				possibleSymbol := pseudoToken.Value
				for i+1 < len(line) {
					i++
					pseudoToken = getTokenType(line[i])
					if pseudoToken.Value == " " {
						break
					}
					possibleSymbol += pseudoToken.Value
				}

				if possibleSymbol == "true" ||
					possibleSymbol == "false" ||
					possibleSymbol == "null" {
					tr = append(tr, Token{
						Type:  SYMBOL,
						Value: possibleSymbol,
					})
				} else {
					return [][]Token{}, errors.New("unidentified symbol " + possibleSymbol)
				}

			}

			if pseudoToken.Type == NUMBER {
				possibleNumber := pseudoToken.Value

				for i+1 < len(line) {
					i++
					pseudoToken = getTokenType(line[i])

					/*
						if the type is a number, continue
						if the value is ".", continue
						if the value is "e", continue
					*/
					if pseudoToken.Type == NUMBER || pseudoToken.Value == "e" || pseudoToken.Value == "." {
						possibleNumber += pseudoToken.Value
					} else {
						break
					}

				}
				lastCharacterOfPossibleNumber := possibleNumber[len(possibleNumber)-1]
				if getTokenType(lastCharacterOfPossibleNumber).Value == "." {
					return [][]Token{}, errors.New("unterminated fraction number")
				}
				if getTokenType(lastCharacterOfPossibleNumber).Value == "e" {
					return [][]Token{}, errors.New("exponent part is missing")
				}
				tr = append(tr, Token{Type: NUMBER, Value: possibleNumber})
			}

			if pseudoToken.Type == COMMA ||
				pseudoToken.Type == COLON ||
				pseudoToken.Type == LEFT_CURLY_BRACE ||
				pseudoToken.Type == RIGHT_CURLY_BRACE ||
				pseudoToken.Type == LEFT_SQUARE_BRACE ||
				//pseudoToken.Type == WHITE_SPACE ||
				pseudoToken.Type == RIGHT_SQUARE_BRACE {
				tr = append(tr, pseudoToken)
			}
			if pseudoToken.Type == UNDEFINED {
				return [][]Token{}, errors.New(fmt.Sprintf("unknow literal \"%v\"", pseudoToken.Value))
			}
		}
		// fmt.Println("Token Rows", tr)
		tokens = append(tokens, tr)
	}
	//fmt.Println("Ending Tokenizing with tokens -> ", tokens)
	return tokens, nil
}

func getTokenType(c uint8) Token {
	val := fmt.Sprintf("%c", c)
	switch {
	case c >= 48 && c <= 57:
		return Token{Type: NUMBER, Value: val}

	case c >= 65 && c <= 90 || c >= 97 && c <= 122:
		return Token{Type: LETTER, Value: val}

	case c == 34:
		return Token{Type: DOUBLE_QUOTE, Value: val}

	case c == 13: // A carriage return is a part of line break
		return Token{Type: CARRIAGE_RETURN, Value: val}

	case c == 32:
		return Token{Type: WHITE_SPACE, Value: val}

	case c == 44:
		return Token{Type: COMMA, Value: val}

	case c == 58:
		return Token{Type: COLON, Value: val}

	case c == 91:
		return Token{Type: LEFT_SQUARE_BRACE, Value: val}

	case c == 93:
		return Token{Type: RIGHT_SQUARE_BRACE, Value: val}

	case c == 123:
		return Token{Type: LEFT_CURLY_BRACE, Value: val}

	case c == 125:
		return Token{Type: RIGHT_CURLY_BRACE, Value: val}

	default:
		// fmt.Println("undefined val", c)
		return Token{Type: UNDEFINED, Value: val}
	}

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

//fmt.Printf("%T\t%c\n", line[i], line[i])
