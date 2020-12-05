package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func getSeat(s string) [2]int {
	rowString := s[:7]
	colString := s[7:]

	row := 0
	for i, c := range rowString {
		if string(c) == "F" {
			continue
		}
		part := 128 / int(math.Pow(float64(2), float64(i+1)))
		row += part
	}

	col := 0
	for i, c := range colString {
		if string(c) == "L" {
			continue
		}
		part := 8 / int(math.Pow(float64(2), float64(i+1)))
		col += part
	}

	return [2]int{row, col}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	seatIDs := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seat := getSeat(scanner.Text())
		seatID := seat[0]*8 + seat[1]
		seatIDs = append(seatIDs, seatID)
	}
	sort.Ints(seatIDs)
	for i, seatID := range seatIDs {
		if i == 0 {
			continue
		}
		if seatIDs[i-1] != seatID-1 {
			fmt.Println("missing", seatID-1)
		}
	}
}
