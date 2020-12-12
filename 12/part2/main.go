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
	xWaypoint int
	yWaypoint int
}

func (s *State) String() string {
	return fmt.Sprintf("ship: (%d, %d) waypoint (%d, %d)", s.xPos, s.yPos, s.xWaypoint, s.yWaypoint)
}

func (s *State) rotateWaypoint(instruction string, amount int) {
	// if waypoint is at (1, 0) and we rotate 90 degrees left
	// it goes to (0, -1)
	// if waypoint is at (1, 0) and we rotate 90 degrees right
	// it goes to (0, 1)
	xWaypointNew := 0
	yWaypointNew := 0

	if instruction == "L" {
		switch amount {
		case 90:
			xWaypointNew = s.yWaypoint
			yWaypointNew = -s.xWaypoint
		case 180:
			xWaypointNew = -s.xWaypoint
			yWaypointNew = -s.yWaypoint
		case 270:
			xWaypointNew = -s.yWaypoint
			yWaypointNew = s.xWaypoint
		default:
			fmt.Println("bad rotation")
		}
	} else {
		switch amount {
		case 90:
			xWaypointNew = -s.yWaypoint
			yWaypointNew = s.xWaypoint
		case 180:
			xWaypointNew = -s.xWaypoint
			yWaypointNew = -s.yWaypoint
		case 270:
			xWaypointNew = s.yWaypoint
			yWaypointNew = -s.xWaypoint
		default:
			fmt.Println("bad rotation")
		}
	}
	s.xWaypoint = xWaypointNew
	s.yWaypoint = yWaypointNew
}

func (s *State) move(instruction string, amount int) {
	switch instruction {
	case "N":
		s.yWaypoint -= amount
	case "S":
		s.yWaypoint += amount
	case "E":
		s.xWaypoint += amount
	case "W":
		s.xWaypoint -= amount
	case "L", "R":
		s.rotateWaypoint(instruction, amount)
	case "F":
		s.xPos += (amount * s.xWaypoint)
		s.yPos += (amount * s.yWaypoint)
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	state := State{xPos: 0, yPos: 0, xWaypoint: 10, yWaypoint: -1}

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
