package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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
		sum += evaluateExpression(line)
	}
	fmt.Println(sum)
}
