package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := make([]int, 0)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("ERROR")
			return
		}
		numbers = append(numbers, i)
	}

	for pos := 25; pos < len(numbers); pos++ {
		target := numbers[pos]
		found := false
		for i := pos - 25; i < pos; i++ {
			for j := pos - 25; j < pos; j++ {
				if i == j {
					continue
				}
				if numbers[i]+numbers[j] == target {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			fmt.Println("number is", target)
			break
		}
	}

}
