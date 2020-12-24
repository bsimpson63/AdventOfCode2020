package main

import (
	"fmt"
	"strconv"
)

func main() {
	//input := "389125467"
	input := "789465123"
	cups := make([]int, 0)
	for _, c := range input {
		v, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Println("oops")
		}
		cups = append(cups, v)
	}

	// prevent wraparound by pinning the current cup to index 0
	for move := 1; move <= 100; move++ {
		fmt.Printf("Move %d\n", move)
		pickedUp := make([]int, 0)
		newCups := make([]int, 0)

		for i, cup := range cups {
			if i == 0 || i >= 4 {
				newCups = append(newCups, cup)
			} else {
				pickedUp = append(pickedUp, cup)
			}
		}
		fmt.Printf("cups %v\n", cups)
		fmt.Printf("picked up %v\n", pickedUp)
		fmt.Printf("cups remaining %v\n", newCups)

		destination := newCups[0] - 1
		if destination <= 0 {
			destination = 9
		}

	outer:
		for {
			for _, cup := range pickedUp {
				if destination == cup {
					destination--
					if destination <= 0 {
						destination = 9
					}
					continue outer
				}
			}
			break
		}
		fmt.Printf("destination cup is %d\n", destination)
		newerCups := make([]int, 0)
		for _, cup := range newCups {
			newerCups = append(newerCups, cup)
			if cup == destination {
				// place the picked up cups after the destination cup
				newerCups = append(newerCups, pickedUp...)
			}
		}
		fmt.Printf("after placing picked up %v\n", newerCups)
		// shift forward to make new current cup
		cups = make([]int, 0)
		for i, cup := range newerCups {
			if i == 0 {
				continue
			}
			cups = append(cups, cup)
		}
		cups = append(cups, newerCups[0])
		fmt.Println()
	}
}
