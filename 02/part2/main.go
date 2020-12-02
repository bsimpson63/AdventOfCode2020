package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	validCount := 0
	invalidCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		fields := strings.Fields(t)
		limitsText, char, password := fields[0], string(fields[1][0]), fields[2]
		s := strings.Split(limitsText, "-")
		p1, _ := strconv.Atoi(s[0])
		p2, _ := strconv.Atoi(s[1])
		// positions are 1-indexed, convert to 0-index
		p1--
		p2--

		if string(password[p1]) == char && string(password[p2]) != char {
			validCount++
		} else if string(password[p1]) != char && string(password[p2]) == char {
			validCount++
		} else {
			invalidCount++
		}
	}
	fmt.Println("There are", validCount, "valid passwords and", invalidCount, "invalid passwords.")
}
