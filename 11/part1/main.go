package main

import (
	"bufio"
	"fmt"
	"os"
)

func print(grid map[[2]int]string) {
	xMax := 0
	yMax := 0
	for coords, _ := range grid {
		if coords[0] > xMax {
			xMax = coords[0]
		}
		if coords[1] > yMax {
			yMax = coords[1]
		}
	}

	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			fmt.Printf("%s", grid[[2]int{x, y}])
		}
		fmt.Println()
	}
}

func isSame(grid map[[2]int]string, other map[[2]int]string) bool {
	xMax := 0
	yMax := 0
	for coords, _ := range grid {
		if coords[0] > xMax {
			xMax = coords[0]
		}
		if coords[1] > yMax {
			yMax = coords[1]
		}
	}

	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			if grid[[2]int{x, y}] != other[[2]int{x, y}] {
				return false
			}
		}
	}
	return true
}

func countOccupied(grid map[[2]int]string) int {
	count := 0
	for _, state := range grid {
		if state == "#" {
			count++
		}
	}
	return count
}

func step(grid map[[2]int]string) map[[2]int]string {
	xMax := 0
	yMax := 0
	for coords, _ := range grid {
		if coords[0] > xMax {
			xMax = coords[0]
		}
		if coords[1] > yMax {
			yMax = coords[1]
		}
	}

	nextGrid := make(map[[2]int]string)
	for x := 0; x <= xMax; x++ {
		for y := 0; y <= yMax; y++ {
			isEmpty := true
			switch state := grid[[2]int{x, y}]; state {
			case ".":
				// floor, not a seat
				nextGrid[[2]int{x, y}] = "."
				continue
			case "L":
				// empty seat
				isEmpty = true
			case "#":
				// occupied seat
				isEmpty = false
			}

			neighborCount := 0
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if dx == 0 && dy == 0 {
						// this is the seat we are looking at already
						continue
					}
					if state := grid[[2]int{x + dx, y + dy}]; state == "#" {
						neighborCount++
					}
				}
			}
			if isEmpty && neighborCount == 0 {
				nextGrid[[2]int{x, y}] = "#"
			} else if !isEmpty && neighborCount >= 4 {
				nextGrid[[2]int{x, y}] = "L"
			} else {
				nextGrid[[2]int{x, y}] = grid[[2]int{x, y}]
			}
		}
	}
	return nextGrid
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	grid := make(map[[2]int]string)
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, c := range line {
			grid[[2]int{x, y}] = string(c)
		}
		y++
	}

	i := 0
	print(grid)
	for {
		next := step(grid)
		fmt.Println(i)
		print(next)
		if isSame(grid, next) {
			fmt.Println("stable")
			fmt.Println(countOccupied(grid), "occupied seats")
			break
		}
		grid = next
		i++
		fmt.Println()
	}
}
