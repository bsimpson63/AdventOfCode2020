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
		need := 2020 - i
		_, ok := numbers[need]
		if ok {
			fmt.Println("Match:", i, need)
			fmt.Println("Product:", i*need)
		}
	}
}
