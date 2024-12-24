package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	sum_of_multiples := 0
	inputfile, err := os.Open("../input")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(inputfile)
	for scanner.Scan() { //read a line
		re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
		re2 := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		matches := re.FindAllString(scanner.Text(), -1)
		for _, match := range matches {
			matches2 := re2.FindStringSubmatch(match)
			number1, _ := strconv.Atoi(matches2[1])
			number2, _ := strconv.Atoi(matches2[2])
			sum_of_multiples += multiply(number1, number2)
		}
	}
	fmt.Println("The sum of multiples is:", sum_of_multiples)
}

func multiply(i, j int) int {
	return i * j
}
