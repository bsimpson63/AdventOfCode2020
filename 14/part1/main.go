package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
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
	matcher := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
	memory := make(map[int]int)
	mask := make(map[int]string)

	for scanner.Scan() {
		line := scanner.Text()
		match := matcher.FindStringSubmatch(line)
		if match == nil {
			// we got a new mask
			parts := strings.Split(line, " = ")
			maskStr := parts[1]

			mask = make(map[int]string)
			for i, c := range maskStr {
				switch v := string(c); v {
				case "X":
					continue
				case "0", "1":
					mask[i] = v
				}

			}
			//fmt.Println("mask:", mask)
			continue
		}

		addrStr := match[1]
		addr, err := strconv.Atoi(addrStr)
		if err != nil {
			fmt.Println("oops")
			return
		}
		valueStr := match[2]
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			fmt.Println("oops")
			return
		}
		//fmt.Println(addr, value)
		valueBits := fmt.Sprintf("%036s\n", strconv.FormatInt(int64(value), 2))
		//fmt.Println(valueBits)
		masked := 0
		for i, c := range valueBits {
			v := string(c)

			adjustment, isMasked := mask[i]
			if isMasked {
				v = adjustment
			}

			if v == "1" {
				masked += int(math.Pow(float64(2), float64(35-i)))
			}
		}
		memory[addr] = masked
		//fmt.Println("adjusted to", masked)
	}
	sum := 0
	for _, value := range memory {
		sum += value
	}
	fmt.Println(sum)

}
