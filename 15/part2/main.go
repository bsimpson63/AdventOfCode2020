package main

import "fmt"

func main() {
	input := []int{0, 6, 1, 7, 2, 19, 20}
	lastSeenAt := make(map[int]int)
	for i, n := range input {
		if i == len(input)-1 {
			break
		}
		lastSeenAt[n] = i + 1
	}

	previousNumber := input[len(input)-1]
	for i := len(input) + 1; i <= 30000000; i++ {
		nextNumber := 0
		if seenAt, haveSeen := lastSeenAt[previousNumber]; haveSeen {
			// saw previousNumber on turn i-1, we saw it previously at seenAt
			turnsSinceSeen := i - 1 - seenAt
			nextNumber = turnsSinceSeen
			//fmt.Printf("%d: %d (saw %d turn %d)\n", i, nextNumber, previousNumber, seenAt)
		} else {
			//fmt.Printf("%d: 0 (have not seen %d)\n", i, previousNumber)
		}
		// saw previousNumber on turn i-1
		lastSeenAt[previousNumber] = i - 1
		previousNumber = nextNumber
	}
	fmt.Println(previousNumber)
}
