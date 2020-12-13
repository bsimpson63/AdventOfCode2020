package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
1st multiple	3	7	offset
1				3	7	4
2				6	7	1
3				9	14	5
4				12	14	2
5				15	21	6
6				18	21	3
7				21	21	0
8				24	28	4
9				27	28	1
10				30	35	5
11				33	35	2
12				36	42	6
13				39	42	3
14				42	42	0
15				45	49	4
16				48	49	1
17				51	56	5
18				54	56	2
19				57	63	6
20				60	63	3
21				63	63	0
22				66	70	4

find first time we get the desired offset
next time will be jumping forward 7 spots
2 + C*7

example:
want an offset of 1
happens first at multiple of 2
happens next at multiple of 9

start at 2 and jump forward by 7 each time

*/

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	scanner.Scan()
	line = scanner.Text()
	buses := make([][2]int, 0)
	for i, c := range strings.Split(line, ",") {
		s := string(c)
		if s == "x" {
			continue
		}
		bus, _ := strconv.Atoi(s)
		buses = append(buses, [2]int{i, bus})
	}
	fmt.Println(buses)

	// sort buses in descending order so we can take bigger steps
	busesOnly := make([]int, 0)
	offsetByBus := make(map[int]int)
	for _, t := range buses {
		busesOnly = append(busesOnly, t[1])
		offsetByBus[t[1]] = t[0]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(busesOnly)))
	busesAndOffsets := make([][2]int, 0)
	for _, bus := range busesOnly {
		offset := offsetByBus[bus]
		busesAndOffsets = append(busesAndOffsets, [2]int{offset, bus})
	}
	fmt.Println(busesAndOffsets)

	// find T such that T+Oi is divisible by Bi for all i
	// 1068781 + 4 = 59 * 18115
	//jumpSize := busesAndOffsets[1][1]
	jumpConstant := 0

	// find the jumpConstant
	for c := 1; c > 0; c++ {
		time := busesAndOffsets[0][1]*c - busesAndOffsets[0][0]
		offset := busesAndOffsets[1][0]
		bus := busesAndOffsets[1][1]
		if (time+offset)%bus != 0 {
			continue
		}
		//fmt.Printf("found first match at %d\n", time)
		jumpConstant = c
		break
	}
	//fmt.Println(jumpConstant)

outer:
	for i := 1; i > 0; i++ {
		c := jumpConstant + i*busesAndOffsets[1][1]
		time := busesAndOffsets[0][1]*c - busesAndOffsets[0][0]

		if i%10000 == 0 {
			fmt.Println(time)
		}

		offset := busesAndOffsets[1][0]
		bus := busesAndOffsets[1][1]
		if (time+offset)%bus != 0 {
			fmt.Println("didn't work!")
			return
		}
		for i := 2; i < len(busesAndOffsets); i++ {
			offset := busesAndOffsets[i][0]
			bus := busesAndOffsets[i][1]
			if (time+offset)%bus != 0 {
				//fmt.Printf("%d/%d doesn't work\n", bus, offset)
				continue outer
			}
		}
		fmt.Printf("%d works\n", time)
		return
	}

	return
}

func main1() {
	file, err := os.Open("./input_short.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	scanner.Scan()
	line = scanner.Text()
	buses := make([][2]int, 0)
	for i, c := range strings.Split(line, ",") {
		s := string(c)
		if s == "x" {
			continue
		}
		bus, _ := strconv.Atoi(s)
		buses = append(buses, [2]int{i, bus})
	}
	fmt.Println(buses)

	// sort buses in descending order so we can take bigger steps
	busesOnly := make([]int, 0)
	offsetByBus := make(map[int]int)
	for _, t := range buses {
		busesOnly = append(busesOnly, t[1])
		offsetByBus[t[1]] = t[0]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(busesOnly)))
	busesAndOffsets := make([][2]int, 0)
	for _, bus := range busesOnly {
		offset := offsetByBus[bus]
		busesAndOffsets = append(busesAndOffsets, [2]int{offset, bus})
	}
	fmt.Println(busesAndOffsets)

	// find T such that T+Oi is divisible by Bi for all i
	// 1068781 + 4 = 59 * 18115
outer:
	for c := 100000000000000 / 601; c > 0; c++ {
		time := busesAndOffsets[0][1]*c - busesAndOffsets[0][0]
		if c%100000000 == 0 {
			fmt.Println("checking", time)
		}

		for i := 1; i < len(busesAndOffsets); i++ {
			offset := busesAndOffsets[i][0]
			bus := busesAndOffsets[i][1]
			if (time+offset)%bus != 0 {
				//fmt.Printf("%d/%d doesn't work\n", bus, offset)
				continue outer
			}
		}
		fmt.Printf("%d works\n", time)
		return
	}
}
