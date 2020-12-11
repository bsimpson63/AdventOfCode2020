package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var countFromPoint = make(map[int]int)

func findTerminationCount(adaptorSet map[int]bool, start int, end int) int {
	memoizedCount, ok := countFromPoint[start]
	if ok {
		return memoizedCount
	}

	if start == end {
		return 1
	}

	count := 0
	for offset := 1; offset <= 3; offset++ {
		needed := start + offset
		if needed > end {
			continue
		}

		if _, ok := adaptorSet[needed]; ok {
			count += findTerminationCount(adaptorSet, needed, end)
		}
	}
	countFromPoint[start] = count
	return count
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	adaptors := make([]int, 0)

	// add an adaptor to represent the 0 start point
	adaptors = append(adaptors, 0)

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("ERROR")
			return
		}
		adaptors = append(adaptors, i)
	}

	// don't worry about the end, it's automatically a choke point

	sort.Ints(adaptors)
	adaptorSet := make(map[int]bool)
	for _, adaptor := range adaptors {
		adaptorSet[adaptor] = true
	}
	fmt.Println(findTerminationCount(adaptorSet, 0, adaptors[len(adaptors)-1]))
}
