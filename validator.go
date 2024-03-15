package main

import (
	"errors"
)

func Validator(tokens [][]Token) error {
	var err error

	if err = validateParentheses(tokens); err != nil {
		return err
	}
	for i, tr := range tokens {
		var possibleEntry []Token
		for j, t := range tr {
			if i == 0 && j == 0 {
				continue
			}
			if t.Type == COMMA {
				if len(possibleEntry) == 0 {
					return errors.New("expected property name")
				}
				if len(possibleEntry) == 1 {
					if possibleEntry[0].Type == STRING {
						return errors.New("expected ':' after property name")
					}
					if possibleEntry[0].Type == COLON {
						return errors.New("expected key/property before ':'")
					}
					return errors.New("expected property name")
				}
				if len(possibleEntry) == 2 {
					if possibleEntry[0].Type == COLON {
						return errors.New("expected key/property before ':'")
					}
					if possibleEntry[0].Type == STRING {
						if possibleEntry[1].Type == COLON {
							return errors.New("expected value after ':'")
						}
						return errors.New("expected ':' after property name")
					}
					return errors.New("expected property name")
				}

				if len(possibleEntry) == 3 {
					if possibleEntry[0].Type == STRING {
						if possibleEntry[1].Type != COLON {
							return errors.New("expected ':' after property name")
						}
						//	check for a valid value !!
						if possibleEntry[2].Type == STRING ||
							possibleEntry[2].Type == NUMBER ||
							possibleEntry[2].Type == SYMBOL {

						}
					}
				}

				possibleEntry = []Token{}
				continue
			}
			possibleEntry = append(possibleEntry, t)
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
