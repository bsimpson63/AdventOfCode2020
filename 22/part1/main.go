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
	var c1, c2 int

	for i := 0; i < 1000; i++ {
		fmt.Printf("Round %d:\n", i+1)
		fmt.Printf("Player 1's deck: %v\n", player1)
		fmt.Printf("Player 2's deck: %v\n", player2)

		c1, player1 = player1[0], player1[1:]
		c2, player2 = player2[0], player2[1:]

		fmt.Printf("Player 1 plays: %d\n", c1)
		fmt.Printf("Player 2 plays: %d\n", c2)
		if c1 < c2 {
			fmt.Println("Player 2 wins the round")
			player2 = append(player2, c2)
			player2 = append(player2, c1)
		} else {
			fmt.Println("Player 1 wins the round")
			player1 = append(player1, c1)
			player1 = append(player1, c2)
		}

		if len(player1) == 0 {
			fmt.Println("Player 2 wins!")
			break
		} else if len(player2) == 0 {
			fmt.Println("Player 1 wins!")
			break
		}
		fmt.Println()

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
	fmt.Printf("Player 1 score: %d\n", s2)

}
