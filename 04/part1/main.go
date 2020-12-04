package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isValid(tokens map[string]string) bool {
	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		//"cid",
	}

	for _, field := range requiredFields {
		_, ok := tokens[field]
		if !ok {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tokens := make(map[string]string)
	validCount, invalidCount := 0, 0
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			if isValid(tokens) {
				validCount++
			} else {
				invalidCount++
			}
			tokens = make(map[string]string)
			continue
		}

		fields := strings.Split(t, " ")
		for _, field := range fields {
			token := strings.Split(field, ":")
			key, value := token[0], token[1]
			tokens[key] = value
		}
	}
	if isValid(tokens) {
		validCount++
	} else {
		invalidCount++
	}
	fmt.Printf("found %d valid passports, %d invalid passports\n", validCount, invalidCount)

}
