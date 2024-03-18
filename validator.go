package main

import (
	"errors"
	"fmt"
)

func Validator(tokens [][]Token) error {
	var err error

	if err = validateParentheses(tokens); err != nil {
		return err
	}

	for i := 0; i < len(tokens); i++ {
		row := tokens[i]
		for j := 0; j < len(tokens[i]); j++ {
			if j == 0 {
				if i == 0 {
					continue
				}
				if row[j].Type == STRING {
					/*
						This covers cases for
						"key"-> Invalid because of the length check
						"key" "," -> Invalid because of the COLON check
					*/
					if len(row) == j+1 || row[j+1].Type != COLON {
						return errors.New("expected ':' after property/key")
					}
					if len(row) == j+2 {
						return errors.New("expected value after ':'")
					}
					remainingTokens := row[j+2:]

					fmt.Println(remainingTokens)
				}
			}
		}
	}

	return nil
}

func validateParentheses(tokens [][]Token) error {
	var stkForCurly []string
	var stkForSquare []string

	for _, tr := range tokens {
		for _, t := range tr {
			if t.Value == "{" {
				stkForCurly = append(stkForCurly, t.Value)
			}
			if t.Value == "}" {
				if len(stkForCurly) == 0 {
					return errors.New("extra '}'")
				}
				stkForCurly = stkForCurly[1:]
			}
			if t.Value == "[" {
				stkForSquare = append(stkForSquare, t.Value)
			}
			if t.Value == "]" {
				if len(stkForSquare) == 0 {
					return errors.New("extra ']'")
				}
				stkForSquare = stkForSquare[1:]
			}
		}
	}

	if len(stkForCurly) != 0 {
		return errors.New("extra '{'")
	}

	if len(stkForSquare) != 0 {
		return errors.New("extra '['")
	}
	return nil
}

func getValueType(tokens []Token) (string, error) {
	if tokens[0].Type == STRING ||
		tokens[0].Type == NUMBER ||
		tokens[0].Type == SYMBOL {
		if len(tokens) > 2 {
			return "", errors.New("use an array for multiple element")
		}
		//return ""
	}
	if tokens[0].Type == COMMA {
		return "", errors.New("expected a value before ','")
	}

	if tokens[0].Type == COLON {
		return "", errors.New("':' literal is repeated")
	}
	//if tokens[0].Type == LEFT_SQUARE_BRACE {
	//
	//}
	return "", nil
}
