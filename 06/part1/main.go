package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	group := make(map[rune]bool)
	sum := 0
	for scanner.Scan() {
		answers := scanner.Text()
		if answers == "" {
			sum += len(group)
			group = make(map[rune]bool)
			continue
		}
		for _, c := range answers {
			group[c] = true
		}
	}
	sum += len(group)
	fmt.Println(sum)
}
