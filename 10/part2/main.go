package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func findTerminationCount(adaptorSet map[int]bool, start int, end int) int {
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
	return count
}

func findsChokePoints(adaptors []int) []int {
	/*
		a point is a choke if all paths must go through it
		all paths must go through a point if the
		previous point can only go to the current point
	*/
	adaptorSet := make(map[int]bool)
	for _, adaptor := range adaptors {
		adaptorSet[adaptor] = true
	}

	chokePoints := make([]int, 0)

	// is 0 a choke point?
	optionCount := 0
	for offset := 1; offset <= 3; offset++ {
		needed := 0 + offset
		if _, ok := adaptorSet[needed]; ok {
			optionCount++
		}
	}
	if optionCount == 1 {
		chokePoints = append(chokePoints, 0)
	}

	for i, adaptor := range adaptors {
		if i == 0 {
			// already considered this above
			continue
		}
		previous := adaptors[i-1]
		optionCount := 0
		for offset := 1; offset <= 3; offset++ {
			needed := previous + offset
			if _, ok := adaptorSet[needed]; ok {
				optionCount++
			}
		}
		if optionCount == 1 {
			chokePoints = append(chokePoints, adaptor)
		}
	}
	return chokePoints
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

	chokePoints := findsChokePoints(adaptors)
	chokePointsSet := make(map[int]bool)
	for _, chokePoint := range chokePoints {
		chokePointsSet[chokePoint] = true
	}

	fmt.Println("choke points at", chokePoints)

	isChoked := true
	lastChoke := 0
	paths := 1
	for _, adaptor := range adaptors {
		_, isChokePoint := chokePointsSet[adaptor]

		if !isChokePoint {
			if isChoked {
				// transitioning from choked to open
				isChoked = false

			} else {
				// still unchoked
			}
		} else {
			if isChoked {
				// still choked
				lastChoke = adaptor

			} else {
				// transitioning from open to choked
				fmt.Printf("found gap %d->%d\n", lastChoke, adaptor)
				isChoked = true
				paths *= findTerminationCount(adaptorSet, lastChoke, adaptor)
			}

		}
	}
	fmt.Println("paths", paths)
}
