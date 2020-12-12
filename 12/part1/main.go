package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// State of the ship
type State struct {
	xPos      int
	yPos      int
	direction int
}

func (s *State) String() string {
	return fmt.Sprintf("(%d, %d) heading %d", s.xPos, s.yPos, s.direction)
}

func (s *State) turn(turn string, degrees int) {
	if turn == "L" {
		s.direction += degrees
	} else {
		s.direction -= degrees
	}
	if s.direction < 0 {
		s.direction += 360
	}
	if s.direction >= 360 {
		s.direction -= 360
	}
}

func (s *State) move(instruction string, amount int) {
	switch instruction {
	case "N":
		s.yPos -= amount
	case "S":
		s.yPos += amount
	case "E":
		s.xPos += amount
	case "W":
		s.xPos -= amount
	case "L", "R":
		s.turn(instruction, amount)
	case "F":
		switch s.direction {
		case 0:
			s.xPos += amount
		case 90:
			s.yPos -= amount
		case 180:
			s.xPos -= amount
		case 270:
			s.yPos += amount
		default:
			fmt.Println("bad direction")
		}
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	state := State{xPos: 0, yPos: 0, direction: 0}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		instruction := string(line[0])
		c := string(line[1:len(line)])
		amount, err := strconv.Atoi(c)
		if err != nil {
			fmt.Println("oops")
			return
		}

		fmt.Println("starting at", state)
		state.move(instruction, amount)
		fmt.Println(instruction, amount)
		fmt.Println("moved to", state)
	}
}
