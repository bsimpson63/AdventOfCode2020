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
	first := true
	for scanner.Scan() {
		answers := scanner.Text()
		if answers == "" {
			sum += len(group)
			group = make(map[rune]bool)
			first = true
			continue
		}

		if first {
			for _, c := range answers {
				group[c] = true
			}
			first = false
		} else {
			individual := make(map[rune]bool)
			for _, c := range answers {
				individual[c] = true
			}
			for k := range group {
				_, ok := individual[k]
				if !ok {
					delete(group, k)
				}
			}
		}
	}
	sum += len(group)
	fmt.Println(sum)
}
