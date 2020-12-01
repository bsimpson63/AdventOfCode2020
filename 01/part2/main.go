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

	numbers := make(map[int]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		i, err := strconv.Atoi(t)
		if err != nil {
			fmt.Println("Failed to convert:", t)
			continue
		}

		// fmt.Println("Got integer", i)
		numbers[i] = ""
	}
	for i, _ := range numbers {
		for j, _ := range numbers {
			need := 2020 - i - j
			_, ok := numbers[need]
			if ok {
				fmt.Println("Match:", i, j, need)
				fmt.Println("Product:", i*j*need)
				return
			}
		}
	}
}
