package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	x      int
	before []int
}

func main() {
	inputfile, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(inputfile)
	ruleset := make(map[int]Rule)
	sum_of_middle_values_valid := 0
	sum_of_middle_values_corrected := 0
	var updates [][]int
	re := regexp.MustCompile(`(\d{1,2})\|(\d{1,2})`)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "|") { //read a line
			matches := re.FindStringSubmatch(scanner.Text())
			i, _ := strconv.Atoi(matches[1])
			j, _ := strconv.Atoi(matches[2])
			ruleset[i] = Rule{
				x:      i,
				before: append(ruleset[i].before, j),
			}
			//}
		}

		if strings.Contains(scanner.Text(), ",") {
			update_as_strings := strings.Split(scanner.Text(), ",")
			var update []int
			for _, item := range update_as_strings {
				i, _ := strconv.Atoi(item)
				update = append(update, i)
			}
			updates = append(updates, update)
		}
	}
	// do the actual evaluation
	for _, update := range updates {
		if isInCorrectOrder(update, ruleset) {
			middlevalue, err := getMiddleValue(update)
			if err != nil {
				log.Println(err)
				continue
			}
			sum_of_middle_values_valid += middlevalue
		} else {
			orderedUpdate := correctOrder(update, ruleset)
			middlevalue, err := getMiddleValue(orderedUpdate)
			if err != nil {
				log.Println(err)
				continue
			}
			sum_of_middle_values_corrected += middlevalue
		}
	}
	fmt.Println("The result for part 1 is", sum_of_middle_values_valid)
	fmt.Println("The result for part 2 is", sum_of_middle_values_corrected)

}

func isInCorrectOrder(update []int, ruleset map[int]Rule) bool {
	for i, value := range update {
		if i > 0 {
			for j := 0; j < i; j++ {
				if slices.Contains(ruleset[value].before, update[j]) {
					//log.Println("Rule violation (wrong order) found on update", update)
					return false
				}
			}
		}
	}
	return true
}

func correctOrder(update []int, ruleset map[int]Rule) []int {
	var isvalid bool
	for !isvalid {
		issuefound := false
		for i, value := range update {
			if i > 0 {
				for j := 0; j < i; j++ {
					if slices.Contains(ruleset[value].before, update[j]) {
						//log.Println("slice before ordering:", update)
						update = slices.Insert(update, j, value)
						update = slices.Delete(update, i+1, i+2)
						//log.Println("slice after ordering:", update)
						issuefound = true
						break
					}
				}
			}

		}
		if !issuefound {
			isvalid = true
		}
	}
	return update
}

func getMiddleValue(update []int) (int, error) {
	length := len(update)
	if length%2 == 0 {
		return 0, fmt.Errorf("middle value cannot be determinded when length of update is even")
	}
	return update[(length-1)/2], nil
}
