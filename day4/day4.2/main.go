package main

import (
	"bufio"
	"fmt"
	"os"
)

type coords struct {
	x, y int
}

var linelength int

func main() {
	result_count := 0
	inputfile, err := os.Open("../input")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(inputfile)
	data := make(map[coords]string)
	linecount := 1

	for scanner.Scan() { //read a line
		linelength = len(scanner.Text())
		for k, v := range scanner.Text() {
			data[coords{k + 1, linecount}] = string(v)
		}
		linecount++
	}
	fmt.Println(linelength)
	for i := 2; i <= linecount-1; i++ {

		result_count += checkForXMAS(data, i)

	}

	fmt.Println("the result is", result_count)
}

func checkForXMAS(data map[coords]string, linecount int) int {
	count_on_this_line := 0
	for i := 2; i <= linelength-1; i++ {
		if data[coords{i, linecount}] != "A" {
			continue
		}
		// do evaluation onl< if the first letter X is found

		if checkXMAS(data, i, linecount) {
			count_on_this_line++
		}

	}
	return count_on_this_line
}

func checkXMAS(data map[coords]string, x, y int) bool {
	var i, j bool
	if data[coords{x + 1, y - 1}] == "M" && data[coords{x - 1, y + 1}] == "S" {
		i = true
	}
	if data[coords{x + 1, y - 1}] == "S" && data[coords{x - 1, y + 1}] == "M" {
		i = true
	}
	if data[coords{x - 1, y - 1}] == "M" && data[coords{x + 1, y + 1}] == "S" {
		j = true
	}
	if data[coords{x - 1, y - 1}] == "S" && data[coords{x + 1, y + 1}] == "M" {
		j = true
	}
	if i && j {
		return true
	}
	return false
}
