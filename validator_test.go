package main

import (
	"os"
	"testing"
)

func TestValidatorParenthesis(t *testing.T) {
	inputs := []string{
		"{",
		"}",
		"[",
		"]",
		"{}",
		"[]",
		"{]",
		"[}",
		"{{}",
		"{}}",
		"[[]",
		"[]]",
	}

	for _, inp := range inputs {
		t.Run(inp, func(t *testing.T) {
			tokens, err := Tokenize(inp)
			if err != nil {
				t.Errorf("tokenizing error : %v", err)
			}

			err = Validator(tokens)
			if err != nil {
				t.Errorf("validating error : %v", err)
			}
		})
	}
}

func TestValidatorBasic(t *testing.T) {
	inputs := []string{
		"testing_files/step2/invalid_1.json",
	}

	for _, inp := range inputs {
		t.Run(inp, func(t *testing.T) {
			file, _ := os.ReadFile(inp)
			tokens, err := Tokenize(string(file))
			if err != nil {
				t.Errorf("tokenizing error : %v", err)
			}

			err = Validator(tokens)
			if err != nil {
				t.Errorf("validating error : %v", err)
			}
		})
	}
}
