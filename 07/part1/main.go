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

	sources := make(map[Component][]Component)
	for outer, inners := range recipes {
		for _, inner := range inners {
			oneInner := Component{1, inner.bag}
			if _, ok := sources[oneInner]; !ok {
				sources[oneInner] = make([]Component, 0)
			}
			sources[oneInner] = append(sources[oneInner], outer)
		}
	}
	/*
		for inner, outers := range sources {
			fmt.Printf("%+v can be stored in %+v\n", inner, outers)
		}
	*/

	allowedContainers := make(map[Component]bool)
	stack := make([]Component, 0)
	stack = append(stack, Component{1, "shiny gold"})

	for i := 0; i < 1000; i++ {
		// fmt.Printf("stack is %+v\n", stack)
		if len(stack) == 0 {
			fmt.Println("done", i)
			break
		}

		inner := stack[0]
		stack = stack[1:]
		// fmt.Printf("now stack is %+v\n", stack)

		for _, outer := range sources[inner] {
			// fmt.Printf("considering %+v\n", outer)
			_, alreadySeen := allowedContainers[outer]
			if !alreadySeen {
				allowedContainers[outer] = true
				stack = append(stack, outer)
				// fmt.Printf("updated stack is %+v\n", stack)
			}
		}
		// fmt.Printf("and now stack is %+v\n", stack)
	}
	fmt.Printf("%d\n", len(allowedContainers))
}
