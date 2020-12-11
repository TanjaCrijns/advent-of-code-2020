package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func list_string_to_int(list []string) []int {
	var new_list = []int{}
	for _, item := range list {
		item_int, _ := strconv.Atoi(item)
		new_list = append(new_list, item_int)
	}
	return new_list
}

func count_one_three_jolts(adapters []int) int {
	var one_jolt_count, three_jolt_count int = 0, 1
	for i, adapter := range adapters {
		if i == 0 {
			if adapters[i]-1 == 0 {
				one_jolt_count++
			} else {
				three_jolt_count++
			}
		} else {
			if adapter-adapters[i-1] == 1 {
				one_jolt_count++
			} else {
				three_jolt_count++
			}

		}
	}
	return one_jolt_count * three_jolt_count
}

func count_variants(adapters []int) int {
	// Only one way to arrange the starting outlet
	variants := map[int]int{0: 1}
	for _, adapter := range adapters {
		variants[adapter] = variants[adapter-1] + variants[adapter-2] + variants[adapter-3]
	}
	// Only one way way to arrange the end outlet, the difference is always three, which will try to get:
	// variants[adapter-1] = 0 (doesn't exist)
	// variants[adapter-2] = 0 (doesn't exist)
	// variants[adapter-3] = last result, so we just return the last result:
	return variants[adapters[len(adapters)-1]]
}

func main() {
	file_content, _ := ioutil.ReadFile("input.txt")
	input_numbers := list_string_to_int(strings.Split(string(file_content), "\n"))
	sort.Ints(input_numbers)
	fmt.Println("The answer to puzzle 1 is: ", count_one_three_jolts(input_numbers))
	fmt.Println("The answer to puzzle 2 is: ", count_variants(input_numbers))
}
