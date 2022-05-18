package main

import (
	"fmt"
	"strconv"
)

// rescurse down to rightmost expression with * or /,
// calculate, and return
// use that term in the prior expression
// cases:
// - evaluate as is and return subresult
// - pass remaining expressing to be evaluated
// - * as sections first, then +
func execute(expression []string) int {

	var firstTerm int
	// var secondTerm int
	var operation string
	// var nextExpression []string
	var result int

	firstTerm, _ = strconv.Atoi(expression[0])

	if len(expression) == 1 {
		//fmt.Println(firstTerm, ".")
		result = firstTerm
	}
	if len(expression) == 3 {
		// at end of sub section?
	}

	if len(expression) >= 3 {
		operation = expression[1]
		//fmt.Println(firstTerm, " ", operation, " ")

		switch operation {
		case "*":
			result = firstTerm * execute(expression[2:])
		case "/":
			result = firstTerm / execute(expression[2:])
		case "+":
			result = firstTerm + execute(expression[2:])
		case "-":
			result = firstTerm - execute(expression[2:])
		default:
			fmt.Println("???")
		}
	}

	return result
}

// peek ahead helpers

func nextMulDiv(expression []string) int {
	for i, s := range expression {
		if s == "*" || s == "/" {
			return i
		}
	}
	return 0
}

func nextAddSub(expression []string) int {
	for i, s := range expression {
		if s == "+" || s == "-" {
			return i
		}
	}
	return 0
}

func main() {

	expressions := [][]string{
		{"4", "/", "2"},
		{"9", "-", "5"},
		{"4", "*", "2"},
		{"9", "+", "7"},
		{"2", "*", "5", "-", "9"},
		{"8", "/", "4", "+", "3"},
		{"1", "+", "2", "*", "3"},
		{"6", "+", "9", "/", "3"},
	}

	for _, e := range expressions {
		// fmt.Println(e)
		fmt.Println(execute(e))
	}

	// a := []string{"4", "/", "2"}
	// b := []string{"1", "*", "2", "+", "3"}

}
