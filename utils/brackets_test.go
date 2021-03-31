package utils

import (
	"testing"
)

func TestValidateBrackets(t *testing.T) {
	testMap := map[string]bool{
		"[()]{}{[()()]()}":        true,
		"[(])":                    false,
		"[({":                     false,
		")}]":                     false,
		"[()]{}{[()()]()}}}}}}}}": false,
		"":                        true,
		"as":                      true,
	}

	for key, value := range testMap {
		isCorrect := ValidateBrackets(key)
		if isCorrect != value {
			t.Errorf("ValidateBrackets(\"%v\") = %v, expected %v", key, isCorrect, value)
		}
	}
}

func TestFixBrackets(t *testing.T) {
	testMap := map[string]string{
		"[()]{}{[()()]()}":        "[()]{}{[()()]()}",
		"[(])":                    "[()]",
		"[({":                     "[({})]",
		")}]":                     "",
		"[()]{}{[()()]()}}}}}}}}": "[()]{}{[()()]()}",
		"[23(56])":                "[23(56)]",
		"[1(2{3":                  "[1(2{3})]",
		"1)2}3]":                  "123",
		"":                        "",
		"as":                      "as",
	}

	for key, value := range testMap {
		result := FixBrackets(key)
		if result != value {
			t.Errorf("TestFixBrackets(\"%v\") = %v, expected %v", key, result, value)
		}
	}
}
