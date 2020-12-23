package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type frame struct {
	id   int
	data map[[2]int]string
}

func (f frame) getBorders() []string {
	borders := make([]string, 0)

	border := ""
	for x := 0; x < 10; x++ {
		border += f.data[[2]int{x, 0}]
	}
	borders = append(borders, border)

	border = ""
	for y := 0; y < 10; y++ {
		border += f.data[[2]int{9, y}]
	}
	borders = append(borders, border)

	border = ""
	for x := 0; x < 10; x++ {
		border += f.data[[2]int{x, 9}]
	}
	borders = append(borders, border)

	border = ""
	for y := 0; y < 10; y++ {
		border += f.data[[2]int{0, y}]
	}
	borders = append(borders, border)

	return borders
}

func (f frame) String() string {
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
	frames := make(map[int]frame)
	for _, frameStr := range framesRaw {
		frame := parseFrame(frameStr)
		frames[frame.id] = frame
		fmt.Println(frame)
	}

	// find the 4 the borders of each frame
	// and their flips (so 8 total) and
	// find those shared by tiles
	frameIdsByBorder := make(map[string][]int)
	for id, frame := range frames {
		// don't treat flips as a different border
		// so sort the 2 options and take the "smaller"
		borders := frame.getBorders()
		for _, border := range borders {
			reversed := reverseString(border)
			if reversed < border {
				border = reversed
			}

			if frameIdsByBorder[border] == nil {
				frameIdsByBorder[border] = make([]int, 0)
			}
			frameIdsByBorder[border] = append(frameIdsByBorder[border], id)
		}
	}

	for border, ids := range frameIdsByBorder {
		fmt.Printf("border %s: %v\n", border, ids)
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

	product := 1
	for id, count := range sharedCountByID {
		fmt.Printf("%d has %d shared borders\n", id, count)
		if count == 2 {
			product *= id
		}
	}
	fmt.Printf("product of corners: %d\n", product)

	// 3457 has 2 shared borders
	// 3709 has 2 shared borders
	// 2111 has 2 shared borders
	// 1093 has 2 shared borders
}
