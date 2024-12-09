package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	var left, right []int
	inputfile, err := os.Open("../input")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(inputfile)
	scanner.Split(bufio.ScanWords)
	var isLeft = true
	for scanner.Scan() {
		if isLeft {
			value, _ := strconv.Atoi(scanner.Text())
			left = append(left, value)
			isLeft = false // for the next run
		} else {
			value, _ := strconv.Atoi(scanner.Text())
			right = append(right, value)
			isLeft = true // so it alternates between left and right
		}
	}
	slices.Sort(left)
	slices.Sort(right)
	total_difference := 0
	for i := range left {
		diff := right[i] - left[i]
		if diff < 0 {
			diff = diff * -1
		}
		total_difference += diff
	}

	fmt.Printf("Total difference: %v \n", total_difference)
}
