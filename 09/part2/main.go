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

	target := 27911108
	pos := 0
	for {
		sum := numbers[pos]
		end := 0
		for end = pos + 1; end < len(numbers); end++ {
			sum += numbers[end]
			if sum >= target {
				break
			}
		}
		if sum == target {
			fmt.Println("found sequence")

			for i := pos; i <= end; i++ {
				fmt.Println(numbers[i])
			}
			// now manually find the smallest and largest numbers and sum them

			break
		}
		pos++
	}

}
