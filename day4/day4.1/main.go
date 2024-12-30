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
	for i := 1; i <= linecount; i++ {
		switch {
		case i <= 3:
			{
				result_count += checkDown(data, i)
			}
		case i > 3 && i <= linecount-3:
			{
				result_count += checkDownAndUp(data, i)
			}
		case i > linecount-3:
			{
				result_count += checkUp(data, i)
			}
		}
	}
	fmt.Println("the result is", result_count)
}

func checkDown(data map[coords]string, linecount int) int {
	count_on_this_line := 0
	for i := 1; i <= linelength; i++ {
		if data[coords{i, linecount}] != "X" {
			continue
		}
		// do evaluation onl< if the first letter X is found
		switch {
		case i <= 3:
			{
				if checkDirection(data, 90, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 135, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 180, i, linecount) {
					count_on_this_line++
				}
			}
		case i > 3 && i <= linelength-3:
			{
				if checkDirection(data, 90, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 135, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 180, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 225, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 270, i, linecount) {
					count_on_this_line++
				}
			}
		case i > linelength-3:
			{
				if checkDirection(data, 180, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 225, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 270, i, linecount) {
					count_on_this_line++
				}
			}
		}
	}
	return count_on_this_line
}

func checkDownAndUp(data map[coords]string, linecount int) int {
	count_on_this_line := 0
	for i := 1; i <= linelength; i++ {
		if data[coords{i, linecount}] != "X" {
			continue
		}
		// do evaluation onl< if the first letter X is found
		switch {
		case i <= 3:
			{
				if checkDirection(data, 0, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 45, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 90, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 135, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 180, i, linecount) {
					count_on_this_line++
				}
			}
		case i > 3 && i <= linelength-3:
			{
				if checkDirection(data, 0, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 45, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 90, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 135, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 180, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 225, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 270, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 315, i, linecount) {
					count_on_this_line++
				}
			}
		case i > linelength-3:
			{
				if checkDirection(data, 0, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 180, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 225, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 270, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 315, i, linecount) {
					count_on_this_line++
				}
			}
		}
	}
	return count_on_this_line
}

func checkUp(data map[coords]string, linecount int) int {
	count_on_this_line := 0
	for i := 1; i <= linelength; i++ {
		if data[coords{i, linecount}] != "X" {
			continue
		}
		// do evaluation onl< if the first letter X is found
		switch {
		case i <= 3:
			{
				if checkDirection(data, 0, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 45, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 90, i, linecount) {
					count_on_this_line++
				}
			}
		case i > 3 && i <= linelength-3:
			{
				if checkDirection(data, 0, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 45, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 90, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 270, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 315, i, linecount) {
					count_on_this_line++
				}
			}
		case i > linelength-3:
			{
				if checkDirection(data, 0, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 270, i, linecount) {
					count_on_this_line++
				}
				if checkDirection(data, 315, i, linecount) {
					count_on_this_line++
				}
			}
		}
	}
	return count_on_this_line
}

func checkDirection(data map[coords]string, direction, x, y int) bool {
	switch {
	case direction == 0:
		{
			if data[coords{x, y - 1}] == "M" && data[coords{x, y - 2}] == "A" && data[coords{x, y - 3}] == "S" {
				return true
			}
		}
	case direction == 45:
		{
			if data[coords{x + 1, y - 1}] == "M" && data[coords{x + 2, y - 2}] == "A" && data[coords{x + 3, y - 3}] == "S" {
				return true
			}
		}
	case direction == 90:
		{
			if data[coords{x + 1, y}] == "M" && data[coords{x + 2, y}] == "A" && data[coords{x + 3, y}] == "S" {
				return true
			}
		}
	case direction == 135:
		{
			if data[coords{x + 1, y + 1}] == "M" && data[coords{x + 2, y + 2}] == "A" && data[coords{x + 3, y + 3}] == "S" {
				return true
			}
		}
	case direction == 180:
		{
			if data[coords{x, y + 1}] == "M" && data[coords{x, y + 2}] == "A" && data[coords{x, y + 3}] == "S" {
				return true
			}
		}
	case direction == 225:
		{
			if data[coords{x - 1, y + 1}] == "M" && data[coords{x - 2, y + 2}] == "A" && data[coords{x - 3, y + 3}] == "S" {
				return true
			}
		}

	case direction == 270:
		{
			if data[coords{x - 1, y}] == "M" && data[coords{x - 2, y}] == "A" && data[coords{x - 3, y}] == "S" {
				return true
			}
		}
	case direction == 315:
		{
			if data[coords{x - 1, y - 1}] == "M" && data[coords{x - 2, y - 2}] == "A" && data[coords{x - 3, y - 3}] == "S" {
				return true
			}
		}
	}
	return false
}
