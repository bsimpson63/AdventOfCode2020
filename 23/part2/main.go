package main

import (
	"fmt"
	"strconv"
)

type cup struct {
	value int
	next  *cup
}

func printCups(c *cup, MaxCups int) {
	fmt.Printf("cups: ")
	for i := 0; i < MaxCups; i++ {
		fmt.Printf("%d ", c.value)
		c = c.next
	}
	fmt.Println()
}

func main() {

	input := "789465123"
	//input := "389125467"
	MaxCups := 1000000

	cupIDs := make([]int, MaxCups)
	for i, c := range input {
		v, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Println("oops")
		}
		cupIDs[i] = v
	}

	for i := 10; i <= MaxCups; i++ {
		cupIDs[i-1] = i
	}

	cups := make(map[int]*cup)
	cups[cupIDs[0]] = &cup{cupIDs[0], nil}
	for i := 1; i < MaxCups; i++ {
		prevID := cupIDs[i-1]
		id := cupIDs[i]

		cups[id] = &cup{id, nil}

		if prev, exists := cups[prevID]; exists {
			prev.next = cups[id]
		} else {
			fmt.Println("OOPS")
		}
	}

	// last cup loops around to the first cup
	cups[cupIDs[MaxCups-1]].next = cups[cupIDs[0]]

	current := cups[cupIDs[0]]
	for move := 1; move <= 10000000; move++ {
		if move%1000 == 0 {
			fmt.Printf("move %d\n", move)
		}
		//
		//printCups(current, MaxCups)

		pickedUp := make([]int, 3)

		pickedUp[0] = current.next.value
		pickedUp[1] = current.next.next.value
		pickedUp[2] = current.next.next.next.value

		//fmt.Printf("pick up: %v\n", pickedUp)

		pickedUpIDs := make(map[int]bool)
		pickedUpIDs[pickedUp[0]] = true
		pickedUpIDs[pickedUp[1]] = true
		pickedUpIDs[pickedUp[2]] = true

		destination := current.value - 1
		for {
			if _, isPickedUp := pickedUpIDs[destination]; isPickedUp {
				destination--
				continue
			}
			if destination <= 0 {
				destination = MaxCups
				continue
			}
			break
		}

		//fmt.Printf("destination: %d\n", destination)

		// the current cup used to point to the first picked
		// up cup, now it points to the cup the last
		// picked up cup was pointing to
		current.next = cups[pickedUp[2]].next
		// the last picked up cup now points to the cup
		// previously after the destination cup
		cups[pickedUp[2]].next = cups[destination].next
		// put the picked up cups after the destination cup
		cups[destination].next = cups[pickedUp[0]]

		// the next iteration moves one cup over
		current = current.next
		//fmt.Println()
	}

	//fmt.Println("final")
	printCups(cups[1], 3)
}
