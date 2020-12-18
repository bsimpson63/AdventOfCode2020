package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func evaluateSimple(expression string) int {
	// expression has no parens, next strip out the addition
	r := regexp.MustCompile(`(\d+) \+ (\d+)`)
	for {
		m := r.FindAllStringSubmatchIndex(expression, 1)
		if m == nil {
			// no more addition
			break
		}
		matchOne := m[0][2:4]
		matchTwo := m[0][4:6]
		one, err := strconv.Atoi(expression[matchOne[0]:matchOne[1]])
		if err != nil {
			fmt.Println("oops 4")
		}

		two, err := strconv.Atoi(expression[matchTwo[0]:matchTwo[1]])
		if err != nil {
			fmt.Println("oops 5")
		}

		matchString := m[0][0:2]
		subResult := one + two
		remainingExpression := expression[0:matchString[0]] + fmt.Sprintf("%d", subResult) + expression[matchString[1]:len(expression)]
		expression = remainingExpression
	}

	//fmt.Println("evaluating", expression)
	pieces := strings.Split(expression, " ")
	result, err := strconv.Atoi(pieces[0])
	if err != nil {
		fmt.Println("oops")
	}

	operation := ""
	for i := 1; i < len(pieces); i++ {
		if pieces[i] == "+" {
			operation = "+"
			continue
		}
		if pieces[i] == "*" {
			operation = "*"
			continue
		}

		v, err := strconv.Atoi(pieces[i])
		if err != nil {
			fmt.Println("failed to convert", pieces[i])
		}
		if operation == "+" {
			result += v
		} else if operation == "*" {
			result *= v
		} else {
			fmt.Println("oops 3")
		}
		operation = ""
	}
	return result
}

func evaluateExpression(expression string) int {
	// find parens
	r := regexp.MustCompile(`(\([\d\+\* ]+\))`)

	for {
		m := r.FindStringIndex(expression)
		if m == nil {
			// no more parens
			break
		}

		// strip out the leading and trailing parens
		subExpression := expression[m[0]+1 : m[1]-1]
		subResult := evaluateExpression(subExpression)
		remainingExpression := expression[0:m[0]] + fmt.Sprintf("%d", subResult) + expression[m[1]:len(expression)]
		expression = remainingExpression
	}
	return evaluateSimple(expression)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		val := evaluateExpression(line)
		sum += val
	}
	fmt.Println(sum)
}
