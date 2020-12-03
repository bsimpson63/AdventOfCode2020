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
	x, y := 0, 0
	treeCount := 0
	dx, dy := 3, 1
	for scanner.Scan() {
		t := scanner.Text()
		y++
		if y == 1 {
			// skip the first row, this is where we start
			continue
		}

		x, y = x+dx, y+dy

		// valid x is 0 to 30 inclusive
		x = x % 31

		fmt.Println(t)
		fmt.Printf("x=%d, y=%d\n", x, y)

		if string(t[x]) == "#" {
			treeCount++
		}
	}
	fmt.Println("hit", treeCount, "trees")
}
