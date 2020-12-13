package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input_short.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	ts, _ := strconv.Atoi(line)
	fmt.Println("departure timestamp", ts)

	scanner.Scan()
	line = scanner.Text()
	buses := make([]int, 0)
	for _, c := range strings.Split(line, ",") {
		s := string(c)
		if s == "x" {
			continue
		}
		bus, _ := strconv.Atoi(s)
		buses = append(buses, bus)
	}
	fmt.Println(buses)

	// find the smallest multiple of bus that is larger than ts
	bestBus := 0
	bestWait := 0
	for _, bus := range buses {
		wait := ((ts / bus) + 1) * bus
		if bestWait == 0 || wait < bestWait {
			bestWait = wait
			bestBus = bus
		}
		fmt.Println(bus, wait)
	}
	fmt.Printf("best bus is %d, wait is %d\n", bestBus, bestWait)
	fmt.Printf("desired departure is %d\n", ts)
	fmt.Printf("extra wait is %d\n", bestWait-ts)
	fmt.Printf("extra wait times bus number is %d\n", (bestWait-ts)*bestBus)
}
