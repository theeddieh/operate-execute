package main

import (
	"fmt"
	"strconv"
)

// rescurse down to rightmost expression with * or /,
// calculate, and return
// use that term in the prior expression

func execute(expression []string) int {

	var firstTerm, secondTerm, answerTerm int
	var operation string
	// var nextExpression[]string

	firstTerm, _ = strconv.Atoi(expression[0])
	operation = expression[1]
	secondTerm, _ = strconv.Atoi(expression[2])

	// nextExpression = expression[2:]
	// secondTerm = execute(nextExpression)

	switch operation {
	case "*":
		answerTerm = (firstTerm * secondTerm)
	case "/":
		answerTerm = (firstTerm / secondTerm)
	case "+":
		answerTerm = (firstTerm + secondTerm)
	case "-":
		answerTerm = (firstTerm - secondTerm)
	default:
		fmt.Println("operation error")
	}

	return answerTerm
}

func main() {

	expressions := [][]string{
		{"4", "/", "2"},
		{"9", "-", "5"},
		{"4", "*", "2"},
		{"9", "+", "7"},
		{"1", "*", "5", "-", "2"},
		{"8", "/", "2", "+", "3"},
		{"3", "+", "2", "*", "2"},
		{"6", "+", "9", "/", "3"},
	}

	for _, e := range expressions {
		// fmt.Println(e)
		fmt.Println(execute(e))
	}

	// a := []string{"4", "/", "2"}
	// b := []string{"1", "*", "2", "+", "3"}

}
