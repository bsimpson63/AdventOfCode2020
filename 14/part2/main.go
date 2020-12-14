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

func makeMasks(maskStr string) []map[int]string {
	masks := make([]map[int]string, 0)
	floatingBits := make([]int, 0)
	baseMask := make(map[int]string)
	for i, c := range maskStr {
		switch v := string(c); v {
		case "X":
			// floating
			floatingBits = append(floatingBits, i)
		case "0":
			// don't change value
			continue
		case "1":
			// overwrite with 1
			baseMask[i] = v
		}
	}
	numMasks := int(math.Pow(float64(2), float64(len(floatingBits))))
	masks = make([]map[int]string, numMasks)
	for i := 0; i < numMasks; i++ {
		masks[i] = make(map[int]string)
		for j, val := range baseMask {
			masks[i][j] = val
		}

	}

	// generate all permutations of the floating bits
	fmtString := fmt.Sprintf("%%0%ds", len(floatingBits))
	for i := 0; i < len(floatingBits); i++ {
		for j := 0; j < numMasks; j++ {
			// set half i values to 0, half to 1
			// check the ith bit of j
			valueBits := fmt.Sprintf(fmtString, strconv.FormatInt(int64(j), 2))
			positionToFlip := floatingBits[i]
			if string(valueBits[i]) == "0" {
				masks[j][positionToFlip] = "0"
			} else if string(valueBits[i]) == "1" {
				masks[j][positionToFlip] = "1"
			} else {
				fmt.Println("WHAT")
			}
		}
	}
	//fmt.Println("got", maskStr)
	//fmt.Println("base", baseMask)
	//fmt.Println("returning", masks)
	return masks
}

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
	masks := make([]map[int]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		match := matcher.FindStringSubmatch(line)

		if match == nil {
			// we got a new mask
			parts := strings.Split(line, " = ")
			maskStr := parts[1]
			masks = makeMasks(maskStr)
			continue
		}

		// not a mask
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

		for _, mask := range masks {
			addrBits := fmt.Sprintf("%036s\n", strconv.FormatInt(int64(addr), 2))
			masked := 0
			for i, c := range addrBits {
				v := string(c)

				adjustment, isMasked := mask[i]
				if isMasked {
					v = adjustment
				}

				if v == "1" {
					masked += int(math.Pow(float64(2), float64(35-i)))
				}
			}
			memory[masked] = value
			//fmt.Println(addr, "adjusted to", masked)
		}
	}
	sum := 0
	for _, value := range memory {
		sum += value
	}
	fmt.Println(sum)
}
