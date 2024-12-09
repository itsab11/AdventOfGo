package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var valid_reports int
	inputfile, err := os.Open("../input")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(inputfile)
	for scanner.Scan() { //read a line
		values := []int{}
		values_as_string := strings.Split(scanner.Text(), " ")
		for _, v := range values_as_string {
			value_as_int, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
			}
			values = append(values, value_as_int)
		}

		var increasing, invalid bool
		t := values[0] - values[1]
		if t < 0 {
			increasing = true
		} else if t > 0 {
			increasing = false
		} else if t == 0 {
			continue
		}

		for i := 1; i < len(values); i++ {
			switch increasing {
			case true:
				{
					if values[i]-values[i-1] <= 3 && values[i]-values[i-1] > 0 {
						continue
					}
				}
			case false:
				{
					if values[i-1]-values[i] <= 3 && values[i-1]-values[i] > 0 {
						continue
					}
				}
			}
			invalid = true
			break
		}
		if !invalid {
			valid_reports++
		}

	}
	fmt.Printf("Valid reports: %v \n", valid_reports)
}
