package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
every tile has six neighbors: east, southeast, southwest, west,
northwest, and northeast. These directions are given in your list,
respectively, as e, se, sw, w, nw, and ne.

how to lay out the grid?
x,y

east moves 1 unit in x direction
west moves -1 unit in x direction
southeast moves -1 unit in y direction, 0.5 unit in x direction
southwest moves -1 unit in y direction, -0.5 unit in x direction
northeast moves 1 unit in y direction, 0.5 unit in x direction
northwest moves 1 unit in y direction, -0.5 unit in x direction

*/

func getDestination(directionsStr string) [2]float64 {
	directions := make([]string, 0)
	i := 0
	for {
		if i >= len(directionsStr) {
			break
		}

		if string(directionsStr[i]) == "s" || string(directionsStr[i]) == "n" {
			direction := string(directionsStr[i]) + string(directionsStr[i+1])
			directions = append(directions, direction)
			i += 2
		} else {
			directions = append(directions, string(directionsStr[i]))
			i++
		}
	}

	x, y := 0., 0.
	for _, direction := range directions {
		switch direction {
		case "e":
			x++
		case "w":
			x--
		case "se":
			y--
			x += 0.5
		case "sw":
			y--
			x -= 0.5
		case "ne":
			y++
			x += 0.5
		case "nw":
			y++
			x -= 0.5
		default:
			fmt.Println("OOPS")
		}
	}

	return [2]float64{x, y}
}

func takeStep(flipped map[[2]float64]bool) map[[2]float64]bool {
	// go through each flipped tile and update the
	// counts for each of its 6 neighbors
	neighborOffsets := make([][2]float64, 6)
	neighborOffsets[0] = [2]float64{1, 0}
	neighborOffsets[1] = [2]float64{-1, 0}
	neighborOffsets[2] = [2]float64{0.5, -1}
	neighborOffsets[3] = [2]float64{-0.5, -1}
	neighborOffsets[4] = [2]float64{0.5, 1}
	neighborOffsets[5] = [2]float64{-0.5, 1}
	neighborCounts := make(map[[2]float64]int)
	for pos := range flipped {
		for _, offset := range neighborOffsets {
			neighborPos := [2]float64{pos[0] + offset[0], pos[1] + offset[1]}
			neighborCounts[neighborPos]++
		}
	}

	next := make(map[[2]float64]bool)

	for pos, neighborCount := range neighborCounts {
		isFlipped := flipped[pos]
		if isFlipped {
			if neighborCount == 0 || neighborCount > 2 {
				// flipped back to white
			} else {
				next[pos] = true
			}
		} else {
			if neighborCount == 2 {
				next[pos] = true
			} else {
				// stays white
			}
		}
	}
	return next
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	/*
		fmt.Println(getDestination("esew"))
		fmt.Println(getDestination("nwwswee"))
	*/

	flipped := make(map[[2]float64]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		destination := getDestination(line)
		if _, isFlipped := flipped[destination]; isFlipped {
			// this was already flipped, flip it back by deleting it
			delete(flipped, destination)
		} else {
			flipped[destination] = true
		}
	}
	fmt.Println(len(flipped))
	for i := 1; i <= 100; i++ {
		flipped = takeStep(flipped)
		fmt.Println(len(flipped))
	}

}
