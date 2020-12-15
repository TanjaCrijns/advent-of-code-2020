package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func puzzle(input string, goal int) int {
	last_number := 0
	number_dict := map[int]int{}
	for i, s := range strings.Split(strings.TrimSpace(string(input)), ",") {
		last_number, _ = strconv.Atoi(s)
		number_dict[last_number] = i + 1
	}

	for i := len(number_dict); i < goal; i++ {
		if i == len(number_dict) {
			number_dict[last_number] = i
			last_number = 0
		} else if val, ok := number_dict[last_number]; ok {
			number_dict[last_number] = i
			last_number = i - val
		} else {
			number_dict[last_number] = i
			last_number = 0
		}
	}
	return last_number
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	fmt.Println("The answer to puzzle 1 is: ", puzzle(string(input), 2020))
	fmt.Println("The answer to puzzle 2 is: ", puzzle(string(input), 30000000))
}
