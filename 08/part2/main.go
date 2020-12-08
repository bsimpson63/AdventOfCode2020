package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findExit(lines []string, flipPos int) bool {
	accumulator := 0
	pos := 0
	seen := make(map[int]bool)

	for {
		line := lines[pos]
		parts := strings.Split(line, " ")
		command := parts[0]
		i, _ := strconv.Atoi(parts[1])

		if pos == flipPos {
			if command == "jmp" {
				command = "nop"
			} else if command == "nop" {
				command = "jmp"
			}
		}

		switch command {
		case "nop":
			pos++
		case "acc":
			accumulator += i
			pos++
		case "jmp":
			pos += i
		}

		_, alreadySeen := seen[pos]
		if alreadySeen {
			fmt.Println("loop", accumulator)
			return false
		} else if pos >= len(lines) {
			fmt.Println("exiting", accumulator)
			return true
		} else {
			seen[pos] = true
		}
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	for i := 0; i < 1000; i++ {
		found := findExit(lines, i)
		if found {
			break
		}
	}

}
