package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type frame struct {
	id   int
	data map[[2]int]string
}

func rotateClockwise(data map[[2]int]string, dim int) map[[2]int]string {
	newData := make(map[[2]int]string)
	for x := 0; x < dim; x++ {
		for y := 0; y < dim; y++ {
			newData[[2]int{x, y}] = data[[2]int{dim - y - 1, x}]
		}
	}
	return newData
}

func (f *frame) rotateClockwise() {
	/*
		to rotate:
		1 0 2
		0 0 0
		4 0 3

		90 degrees we want to end up with:
		4 0 1
		0 0 0
		3 0 2

		0, 0 --> 2, 0
		2, 0 --> 2, 2
		2, 2 --> 0, 2
		0, 2 --> 0, 0

		see https://stackoverflow.com/questions/42519/how-do-you-rotate-a-two-dimensional-array
	*/

	f.data = rotateClockwise(f.data, 10)
}

func flipHorizontal(data map[[2]int]string, dim int) map[[2]int]string {
	newData := make(map[[2]int]string)
	for x := 0; x < dim; x++ {
		for y := 0; y < dim; y++ {
			newData[[2]int{x, y}] = data[[2]int{x, dim - y - 1}]
		}
	}
	return newData
}

func (f *frame) flipHorizontal() {
	/*
		flip along the horizontal axis

		start:
		1 0 2
		0 0 0
		4 0 3

		end:
		4 0 3
		0 0 0
		1 0 2

		0, 0 --> 0, 2
		2, 0 --> 2, 2
		2, 2 --> 2, 0
		0, 2 --> 0, 0
	*/

	f.data = flipHorizontal(f.data, 10)
}

func (f *frame) getBorder(b int) string {
	/*
		0: north
		1: east
		2: south
		3: west
	*/
	border := ""
	switch b {
	case 0:
		for x := 0; x < 10; x++ {
			border += f.data[[2]int{x, 0}]
		}
	case 1:
		for y := 0; y < 10; y++ {
			border += f.data[[2]int{9, y}]
		}
	case 2:
		for x := 0; x < 10; x++ {
			border += f.data[[2]int{x, 9}]
		}
	case 3:
		for y := 0; y < 10; y++ {
			border += f.data[[2]int{0, y}]
		}
	default:
		fmt.Println("bad border", b)
	}

	// switch to the canonical sorting
	reversed := reverseString(border)
	if reversed < border {
		border = reversed
	}

	return border
}

func (f *frame) getBorders() []string {
	borders := make([]string, 0)

	borders = append(borders, f.getBorder(0))
	borders = append(borders, f.getBorder(1))
	borders = append(borders, f.getBorder(2))
	borders = append(borders, f.getBorder(3))

	return borders
}

func (f *frame) String() string {
	s := fmt.Sprintf("Tile %d\n", f.id)
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			s += fmt.Sprintf(f.data[[2]int{x, y}])
		}
		s += fmt.Sprintf("\n")
	}
	return s
}

func reverseString(s string) string {
	result := ""
	for _, c := range s {
		result = string(c) + result
	}
	return result
}

func parseFrame(frameStr string) frame {
	lines := strings.Split(frameStr, "\n")
	idStrPieces := strings.Split(lines[0], " ")
	idStr := idStrPieces[1][:len(idStrPieces[1])-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("oops")
	}

	data := make(map[[2]int]string)
	for y := 0; y < len(lines)-1; y++ {
		// add 1 to y because line 0
		for x, c := range lines[y+1] {
			data[[2]int{x, y}] = string(c)
		}
	}
	return frame{id, data}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := ""
	for scanner.Scan() {
		line := scanner.Text()
		data += line + "\n"
	}
	framesRaw := strings.Split(data, "\n\n")

	// NOTE: in golang map entries are not addressable
	// so we can't use frame values, we must use pointers
	frames := make(map[int]*frame)
	for _, frameStr := range framesRaw {
		frame := parseFrame(frameStr)
		frames[frame.id] = &frame
		//fmt.Println(frame)
	}

	// find the 4 the borders of each frame
	// and their flips (so 8 total) and
	// find those shared by tiles
	frameIdsByBorder := make(map[string][]int)
	for id, frame := range frames {
		borders := frame.getBorders()
		for _, border := range borders {
			if frameIdsByBorder[border] == nil {
				frameIdsByBorder[border] = make([]int, 0)
			}
			frameIdsByBorder[border] = append(frameIdsByBorder[border], id)
		}
	}

	// find the frame with the most shared borders--the
	// middle frame must have 4 shared borders
	sharedCountByID := make(map[int]int)
	for _, ids := range frameIdsByBorder {
		if len(ids) == 1 {
			// this frame is only see on one frame
			continue
		}

		for _, id := range ids {
			sharedCountByID[id]++
		}
	}

	nextFrameID := 0
	for id, count := range sharedCountByID {
		//fmt.Printf("%d has %d shared borders\n", id, count)
		if count == 2 {
			if nextFrameID == 0 {
				nextFrameID = id
				break
			}
		}
	}

	// frameIdsByBorder is border: [frame ids] (map[string][]int)

	grid := make(map[[2]int]int)
	grid[[2]int{0, 0}] = nextFrameID

	// flip the first corner until its unmatched edges are
	// pointed north, west
	// or it's matched edges are pointed east, south
	borders := frames[nextFrameID].getBorders()
	matchedBorders := make([]string, 0)
	for _, border := range borders {
		if len(frameIdsByBorder[border]) == 2 {
			matchedBorders = append(matchedBorders, border)
		}
	}
	desiredBorders := make([]string, 4)
	desiredBorders[0] = ""
	desiredBorders[1] = matchedBorders[0]
	desiredBorders[2] = matchedBorders[1]
	desiredBorders[3] = ""

	// print before/after align() to see the change
	//fmt.Println(frames[nextFrameID])
	frames[nextFrameID].align(desiredBorders)
	//fmt.Println(frames[nextFrameID])

	fmt.Printf("(%d, %d): %d\n", 0, 0, nextFrameID)

	// next one shares this frame's east border
	for _, id := range frameIdsByBorder[matchedBorders[0]] {
		if id != nextFrameID {
			nextFrameID = id
		}
	}

	gridDimension := int(math.Sqrt(float64(len(frames))))

	for y := 0; y < gridDimension; y++ {
		for x := 0; x < gridDimension; x++ {
			if x == 0 && y == 0 {
				// first corner, we already did this
				continue
			}

			if x != 0 {
				// frame at x, y shares this frame's east border of
				// frame at x-1, y
				previousFrame := frames[grid[[2]int{x - 1, y}]]
				previousBorder := previousFrame.getBorder(1)
				for _, id := range frameIdsByBorder[previousBorder] {
					if id != previousFrame.id {
						nextFrameID = id
					}
				}
			} else {
				// next one shares previous row's south border
				previousFrame := frames[grid[[2]int{0, y - 1}]]
				previousBorder := previousFrame.getBorder(2)
				for _, id := range frameIdsByBorder[previousBorder] {
					if id != previousFrame.id {
						nextFrameID = id
					}
				}
			}

			fmt.Printf("(%d, %d): %d\n", x, y, nextFrameID)
			grid[[2]int{x, y}] = nextFrameID

			// zero these back out
			desiredBorders[0] = ""
			desiredBorders[1] = ""
			desiredBorders[2] = ""
			desiredBorders[3] = ""

			borders = frames[nextFrameID].getBorders()

			// for y != 0 match the border to our west
			// and the border to our north
			// for y == 0 the border to our north is the
			// one that has no other match
			// for x == 0 the border to our west is the
			// one that has no other match
			// corners are special--they have 2 non matching borders
			// so we need to specify both the matching borders instead
			if x == gridDimension-1 && y == 0 {
				// north east corner
				// must specify matching border to west
				// other matched border to south
				desiredBorders[3] = frames[grid[[2]int{x - 1, y}]].getBorder(1)

				for _, border := range borders {
					if border == desiredBorders[3] {
						continue
					}
					if len(frameIdsByBorder[border]) == 2 {
						desiredBorders[2] = border
					}
				}
			} else if x == 0 && y == gridDimension-1 {
				// south west corner
				// must specify matching border to north and
				// other matched border to east
				desiredBorders[0] = frames[grid[[2]int{x, y - 1}]].getBorder(2)

				for _, border := range borders {
					if border == desiredBorders[0] {
						continue
					}
					if len(frameIdsByBorder[border]) == 2 {
						desiredBorders[1] = border
					}
				}
			} else {
				if y != 0 {
					// match the border to our north (that frame's south border)
					desiredBorders[0] = frames[grid[[2]int{x, y - 1}]].getBorder(2)
				} else {
					// north border has no match
					for _, border := range borders {
						if len(frameIdsByBorder[border]) == 1 {
							desiredBorders[0] = border
						}
					}
				}

				if x != 0 {
					// match the border to our west (that frame's east border)
					desiredBorders[3] = frames[grid[[2]int{x - 1, y}]].getBorder(1)
				} else {
					// west border has no match
					for _, border := range borders {
						if len(frameIdsByBorder[border]) == 1 {
							desiredBorders[3] = border
						}
					}
				}
			}

			frames[nextFrameID].align(desiredBorders)
		}
	}

	// print the grid!
	// start with the first row of the grid: (0, 0), (1, 0), (2, 0)
	// print the first row of each frame
	// print the second row of each frame
	for y := 0; y < gridDimension; y++ {
		// print the heaters
		for x := 0; x < gridDimension; x++ {
			frameID := grid[[2]int{x, y}]
			fmt.Printf("%d       ", frameID)
		}
		fmt.Println()

		for row := 0; row < 10; row++ {
			for x := 0; x < gridDimension; x++ {
				frameID := grid[[2]int{x, y}]
				frame := frames[frameID]
				for col := 0; col < 10; col++ {
					fmt.Printf("%s", frame.data[[2]int{col, row}])
				}
				fmt.Printf(" ")
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}

	combinedGrid := make(map[[2]int]string)
	for y := 0; y < gridDimension; y++ {
		for x := 0; x < gridDimension; x++ {
			// only use the inner 8x8 of each frame
			// after removing the border
			frameID := grid[[2]int{x, y}]
			frame := frames[frameID]

			for row := 1; row < 9; row++ {
				for col := 1; col < 9; col++ {
					cX := x*8 + col - 1
					cY := y*8 + row - 1
					d, isThere := frame.data[[2]int{col, row}]
					if !isThere {
						fmt.Println("oops")
					}
					combinedGrid[[2]int{cX, cY}] = d
				}
			}
		}
	}

	fmt.Println("Combined image:")
	for y := 0; y < 8*gridDimension; y++ {
		for x := 0; x < 8*gridDimension; x++ {
			fmt.Printf("%s", combinedGrid[[2]int{x, y}])
		}
		fmt.Println()
	}

	monsterOffsets := make(map[[2]int]string)
	for y, line := range strings.Split(monster, "\n") {
		for x, c := range line {
			if string(c) == "#" {
				monsterOffsets[[2]int{x, y}] = "#"
			}
		}
	}
	fmt.Println(monsterOffsets)

	maxMonsterCount := 0
	for i := 0; i < 8; i++ {
		monsterCount := 0
		for x := 0; x < 8*gridDimension; x++ {
		y:
			for y := 0; y < 8*gridDimension; y++ {
				for offset := range monsterOffsets {
					xM := x + offset[0]
					yM := y + offset[1]
					if combinedGrid[[2]int{xM, yM}] != "#" {
						continue y
					}
				}
				monsterCount++
			}
		}

		// this orientation's result
		fmt.Printf("Found %d monsters\n", monsterCount)
		if monsterCount > maxMonsterCount {
			maxMonsterCount = monsterCount
		}

		// rotate
		combinedGrid = rotateClockwise(combinedGrid, 8*gridDimension)
		if i == 3 {
			combinedGrid = flipHorizontal(combinedGrid, 8*gridDimension)
		}
	}
	tilesPerMonster := len(monsterOffsets)
	monsterTiles := tilesPerMonster * maxMonsterCount
	totalTiles := 0
	for _, val := range combinedGrid {
		if val == "#" {
			totalTiles++
		}
	}
	fmt.Printf("%d total tiles, %d occupied by monsters, %d remaining\n", totalTiles, monsterTiles, totalTiles-monsterTiles)

}

var monster = `                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `

func (f *frame) align(desiredBorders []string) {
outer:
	for i := 0; i < 8; i++ {
		borders := f.getBorders()
		for j, desiredBorder := range desiredBorders {
			if desiredBorder == "" {
				// not constrained, check remaining desiredBorders
				continue
			}
			if borders[j] != desiredBorder {
				// this border isn't correct--go to next flip of frame
				f.rotateClockwise()
				if i == 3 {
					f.flipHorizontal()
				}
				continue outer
			}
		}
		// all the borders matched
		return
	}
	fmt.Println("whoops!")
}
