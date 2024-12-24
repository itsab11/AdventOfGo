package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sum_of_multiples := 0
	content, err := os.ReadFile("../input")
	str_content := string(content)
	if err != nil {
		fmt.Println(err)
	}
	parts := strings.Split(str_content, "do()")
	for _, part := range parts {
		valid, _, _ := strings.Cut(part, "don't()")
		re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
		re2 := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		matches := re.FindAllString(valid, -1)
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
