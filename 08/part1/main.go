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

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	accumulator := 0
	pos := 0
	seen := make(map[int]bool)

	for {
		line := lines[pos]
		parts := strings.Split(line, " ")
		command := parts[0]
		i, _ := strconv.Atoi(parts[1])

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
			break
		} else {
			seen[pos] = true
		}
	}
}
