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

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	slopes := [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	treeCountBySlope := make(map[[2]int]int, len(slopes))

	for _, slope := range slopes {
		dx := slope[0]
		dy := slope[1]

		x := 0
		treeCount := 0

		for y, line := range lines {
			if y == 0 {
				// skip the first row, this is where we start
				continue
			}

			if y%dy != 0 {
				// skip this row, we never actually land here
				continue
			}

			x = x + dx
			// valid x is 0 to 30 inclusive
			x = x % 31

			if string(line[x]) == "#" {
				treeCount++
			}
		}
		fmt.Printf("(%d,%d) hit %d trees\n", dx, dy, treeCount)
		treeCountBySlope[slope] = treeCount
	}
	product := 1
	for _, count := range treeCountBySlope {
		product *= count
	}
	fmt.Printf("product of all hits is %d\n", product)
}
