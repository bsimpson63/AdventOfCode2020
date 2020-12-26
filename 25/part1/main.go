package main

import "fmt"

func main() {

	//fmt.Printf("loop size is %d\n", findLoopSize(7, 5764801))
	//fmt.Printf("loop size is %d\n", findLoopSize(7, 17807724))
	loopSize1 := findLoopSize(7, 10212254)
	fmt.Printf("loop size is %d\n", loopSize1)
	loopSize2 := findLoopSize(7, 12577395)
	fmt.Printf("loop size is %d\n", loopSize2)

	res1 := transform(10212254, loopSize2)
	res2 := transform(12577395, loopSize1)
	fmt.Printf("%d, %d\n", res1, res2)
}

func findLoopSize(subject int, target int) int {
	loopSize := 1
	value := 1
	for {
		value *= subject
		value = value % 20201227
		if value == target {
			break
		}
		loopSize++
	}
	return loopSize
}

func transform(subject int, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value *= subject
		value = value % 20201227
	}
	return value
}
