package main

import (
	"fmt"
	"strings"
)

func step(grid map[[3]int]bool) map[[3]int]bool {
	xMax, xMin := 0, 0
	yMax, yMin := 0, 0
	zMax, zMin := 0, 0
	for coords, val := range grid {
		if !val {
			continue
		}

		if coords[0] > xMax {
			xMax = coords[0]
		}
		if coords[1] > yMax {
			yMax = coords[1]
		}
		if coords[2] > zMax {
			zMax = coords[2]
		}

		if coords[0] < xMin {
			xMin = coords[0]
		}
		if coords[1] < yMin {
			yMin = coords[1]
		}
		if coords[2] < zMin {
			zMin = coords[2]
		}

	}

	// consider grid points one beyond our current border
	xMax++
	yMax++
	zMax++
	xMin--
	yMin--
	zMin--

	nextGrid := make(map[[3]int]bool)
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			for z := zMin; z <= zMax; z++ {
				isActive := grid[[3]int{x, y, z}]
				neighborCount := 0
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						for dz := -1; dz <= 1; dz++ {
							if dx == 0 && dy == 0 && dz == 0 {
								// this is the cube we are looking at already
								continue
							}
							if state := grid[[3]int{x + dx, y + dy, z + dz}]; state == true {
								neighborCount++
							}
						}
					}
				}
				if isActive && (neighborCount == 2 || neighborCount == 3) {
					nextGrid[[3]int{x, y, z}] = true
				} else if !isActive && neighborCount == 3 {
					nextGrid[[3]int{x, y, z}] = true
				}
			}
		}
	}
	return nextGrid

}

func print(grid map[[3]int]bool) {
	xMax, xMin := 0, 0
	yMax, yMin := 0, 0
	zMax, zMin := 0, 0

	for coords, val := range grid {
		if !val {
			continue
		}

		if coords[0] > xMax {
			xMax = coords[0]
		}
		if coords[1] > yMax {
			yMax = coords[1]
		}
		if coords[2] > zMax {
			zMax = coords[2]
		}

		if coords[0] < xMin {
			xMin = coords[0]
		}
		if coords[1] < yMin {
			yMin = coords[1]
		}
		if coords[2] < zMin {
			zMin = coords[2]
		}

	}

	for z := zMin; z <= zMax; z++ {
		fmt.Printf("z: %d\n", z)
		for y := yMin; y <= yMax; y++ {
			for x := xMin; x <= xMax; x++ {
				if grid[[3]int{x, y, z}] == true {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

var shortInput = `.#.
..#
###`

var input = `#.##....
.#.#.##.
###.....
....##.#
#....###
.#.#.#..
.##...##
#..#.###`

func main() {
	grid := make(map[[3]int]bool)
	for y, line := range strings.Split(input, "\n") {
		for x, c := range line {
			if string(c) == "#" {
				grid[[3]int{x, y, 0}] = true
			}
		}
	}

	/*
		fmt.Println(grid)
		fmt.Println("initial state")
		print(grid)
		grid = step(grid)
		fmt.Println("1 step later:")
		print(grid)
		grid = step(grid)
		fmt.Println("2 steps later:")
		print(grid)
	*/

	for i := 1; i <= 6; i++ {
		fmt.Printf("step %d:\n", i)
		grid = step(grid)
		fmt.Printf("%d active cubes\n", len(grid))
	}
}
