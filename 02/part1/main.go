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
		min, _ := strconv.Atoi(s[0])
		max, _ := strconv.Atoi(s[1])

		count := strings.Count(password, char)

		fmt.Println(password, char, "appears", count, "times.", min, max)

		if count >= min && count <= max {
			validCount++
		} else {
			invalidCount++
		}
	}
	fmt.Println("There are", validCount, "valid passwords and", invalidCount, "invalid passwords.")
}
