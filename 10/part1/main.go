package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	adaptors := make([]int, 0)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("ERROR")
			return
		}
		adaptors = append(adaptors, i)
	}
	sort.Ints(adaptors)
	fmt.Println(adaptors)

	currentJolts := 0
	oneCount := 0
	threeCount := 0
	for _, adaptor := range adaptors {
		offset := adaptor - currentJolts
		if offset > 3 || offset < 0 {
			fmt.Println("OOPS")
			return
		}
		if offset == 1 {
			oneCount++
		} else if offset == 3 {
			threeCount++
		}
		currentJolts = adaptor
	}
	// device is 3+ last adaptor
	threeCount++
	fmt.Printf("one: %d, three: %d\n", oneCount, threeCount)

}
