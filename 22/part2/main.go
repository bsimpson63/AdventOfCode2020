package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	decks := strings.Split(data, "\n\n")

	player1 := make([]int, 0)
	for _, num := range strings.Split(decks[0], "\n") {
		if num == "" {
			continue
		}
		val, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("oops")
		}
		player1 = append(player1, val)
	}

	player2 := make([]int, 0)
	for _, num := range strings.Split(decks[1], "\n") {
		if num == "" {
			continue
		}
		val, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("oops")
		}
		player2 = append(player2, val)
	}
	fmt.Println(player1)
	fmt.Println(player2)

	player1Wins := playGame(&player1, &player2, 1)
	if player1Wins {
		fmt.Println("Player 1 wins overall!")
	} else {
		fmt.Println("Player 2 wins overall!")
	}

	fmt.Printf("\nPlayer 1's deck: %v\n", player1)
	fmt.Printf("Player 2's deck: %v\n", player2)

	s1 := 0
	s2 := 0
	for i := 0; i < len(player1); i++ {
		s1 += (len(player1) - i) * player1[i]
	}
	for i := 0; i < len(player2); i++ {
		s2 += (len(player2) - i) * player2[i]
	}

	fmt.Printf("Player 1 score: %d\n", s1)
	fmt.Printf("Player 2 score: %d\n", s2)
}

func serializeState(player1 *[]int, player2 *[]int) string {
	s := ""
	for _, c1 := range *player1 {
		s += fmt.Sprintf("%d,", c1)
	}
	s += "|"
	for _, c2 := range *player2 {
		s += fmt.Sprintf("%d,", c2)
	}
	return s
}

func playGame(player1 *[]int, player2 *[]int, game int) bool {
	seenStates := make(map[string]bool)
	for round := 1; round < 10000; round++ {
		fmt.Printf("-- Round %d (Game %d) --\n", round, game)
		fmt.Printf("Player 1's deck: %v\n", player1)
		fmt.Printf("Player 2's deck: %v\n", player2)

		state := serializeState(player1, player2)
		if _, alreadySeen := seenStates[state]; alreadySeen {
			fmt.Printf("Already seen this state. Player 1 wins!\n\n")
			return true
		}
		seenStates[state] = true

		c1 := (*player1)[0]
		c2 := (*player2)[0]
		*player1 = (*player1)[1:]
		*player2 = (*player2)[1:]

		fmt.Printf("Player 1 plays: %d\n", c1)
		fmt.Printf("Player 2 plays: %d\n", c2)

		if len(*player1) >= c1 && len(*player2) >= c2 {
			// both players have at least as many cards
			// as the numbers they drew
			// recurse!
			fmt.Printf("Playing a sub-game to determine the winner...\n")
			subPlayer1 := make([]int, 0)
			for i := 0; i < c1; i++ {
				subPlayer1 = append(subPlayer1, (*player1)[i])
			}
			subPlayer2 := make([]int, 0)
			for i := 0; i < c2; i++ {
				subPlayer2 = append(subPlayer2, (*player2)[i])
			}
			player1Won := playGame(&subPlayer1, &subPlayer2, game+1)
			fmt.Printf("...anyway, back to game %d.\n", game)

			if player1Won {
				fmt.Printf("Player 1 wins round %d of game %d!", round, game)
				*player1 = append(*player1, c1)
				*player1 = append(*player1, c2)
			} else {
				fmt.Printf("Player 2 wins round %d of game %d!", round, game)
				*player2 = append(*player2, c2)
				*player2 = append(*player2, c1)
			}
		} else if c1 < c2 {
			fmt.Println("Player 2 wins the round")
			*player2 = append(*player2, c2)
			*player2 = append(*player2, c1)
		} else {
			fmt.Println("Player 1 wins the round")
			*player1 = append(*player1, c1)
			*player1 = append(*player1, c2)
		}

		if len(*player1) == 0 {
			fmt.Printf("Player 2 wins!\n\n")
			return false
		} else if len(*player2) == 0 {
			fmt.Printf("Player 1 wins!\n\n")
			return true
		}
		fmt.Println()
	}
	fmt.Println("exceeded iteration limit")
	return true
}
