package main

import (
	"fmt"
	"io/ioutil"
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

func sum(list []int) int {
	result := 0
	for _, v := range list {
		result += v
	}
	return result
}

func min_max(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func check_validity(number int, previous_numbers []int) bool {
	for _, key := range previous_numbers {
		for _, key1 := range previous_numbers {
			if key+key1 == number && key != key1 {
				return true
			}
		}
	}
	return false
}

func find_weakness(number int, input_numbers []int) int {
	for i, key := range input_numbers {
		list := []int{key}
		for _, key1 := range input_numbers[i+1:] {
			list = append(list, key1)
			if sum(list) == number {
				min, max := min_max(list)
				return min + max
			}
		}
	}
	return 0
}

func main() {
	file_content, _ := ioutil.ReadFile("input.txt")
	input_numbers := list_string_to_int(strings.Split(string(file_content), "\n"))
	const preamble_size int = 25
	var previous_numbers []int
	for i, number := range input_numbers {
		if i <= preamble_size {
			continue
		} else {
			previous_numbers = input_numbers[i-preamble_size : i]
		}
		if !check_validity(number, previous_numbers) {
			fmt.Println("The answer to puzzle 1 is :", number)
			fmt.Println("The answer to puzzle 2 is :", find_weakness(number, input_numbers))
			break
		}
	}
}
