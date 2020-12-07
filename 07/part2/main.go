package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Component struct {
	quantity int
	bag      string
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	r := regexp.MustCompile(`(\d?) ?([\s\w]*) bags?\.?`)
	scanner := bufio.NewScanner(file)
	recipes := make(map[Component][]Component)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " contain ")
		outerStr := parts[0]
		// outer doesn't include a quantity, assume 1
		outerBag := Component{1, r.FindStringSubmatch(outerStr)[2]}
		innerMultipleStr := strings.Split(parts[1], ", ")
		innerBags := make([]Component, 0)
		for _, innerStr := range innerMultipleStr {
			m := r.FindStringSubmatch(innerStr)
			innerQuantity, _ := strconv.Atoi(m[1])
			innerBag := m[2]
			innerBags = append(innerBags, Component{innerQuantity, innerBag})
		}
		recipes[outerBag] = innerBags
	}

	// How many bags within the shiny gold bag?
	stack := make([]Component, 0)
	stack = append(stack, Component{1, "shiny gold"})
	contents := make([]Component, 0)

	for i := 0; i < 10000; i++ {
		// fmt.Printf("stack is %+v\n", stack)
		if len(stack) == 0 {
			fmt.Println("done", i)
			break
		}

		outer := stack[0]
		stack = stack[1:]

		for _, inner := range recipes[outer] {
			for i := 0; i < inner.quantity; i++ {
				stack = append(stack, Component{1, inner.bag})
				contents = append(contents, Component{1, inner.bag})
			}

		}
	}
	fmt.Println(len(contents))

}
