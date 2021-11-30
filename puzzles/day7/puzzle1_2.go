package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type bag struct {
	color    string
	contents []string
	amounts  []int
}

func format_input(input []byte) map[string]bag {
	input_strings := strings.Split(string(input), ".\n")
	bags := make(map[string]bag)
	for _, rule := range input_strings {
		n_contents := strings.Count(rule, ",") + 1
		split_string := strings.Split(rule, " ")
		main_bag := split_string[0] + " " + split_string[1]

		var content_bags []string
		var content_amounts []int
		var current_amount_pos, current_color_1, current_color_2 int = 4, 5, 6
		for i := 0; i < n_contents; i++ {
			color := split_string[current_color_1] + " " + split_string[current_color_2]
			content_bags = append(content_bags, color)
			amount, _ := strconv.Atoi(split_string[current_amount_pos])
			content_amounts = append(content_amounts, amount)
			current_amount_pos += 4
			current_color_1 += 4
			current_color_2 += 4
		}
		temp_bag := bag{color: main_bag, contents: content_bags, amounts: content_amounts}
		bags[temp_bag.color] = temp_bag
	}
	return bags
}

func check_goal(goal string, bag bag, bags map[string]bag) bool {
	if len(bag.contents) == 0 {
		return false
	} else {
		temp_bool := false
		for _, temp_bag := range bag.contents {
			if temp_bag == goal {
				if !temp_bool {
					temp_bool = true
				}
			} else {
				if !temp_bool {
					temp_bool = check_goal(goal, bags[temp_bag], bags)
				}
			}
		}
		return temp_bool
	}
}

func get_bag_in_bag_count(goal string, bags map[string]bag) int {
	total_count := 0
	for i, bag := range bags[goal].contents {
		total_count += bags[goal].amounts[i]
		total_count += bags[goal].amounts[i] * get_bag_in_bag_count(bag, bags)
	}
	return total_count
}
func main() {
	file_content, _ := ioutil.ReadFile("input.txt")
	goal := "shiny gold"

	// Puzzle 1
	bags := format_input(file_content)
	var bag_count int
	for _, bag := range bags {
		if check_goal(goal, bag, bags) {
			bag_count++
		}
	}
	fmt.Println("The answer for puzzle 1 is: ", bag_count)

	// Puzzle 2
	bag_in_bags_count := get_bag_in_bag_count(goal, bags)
	fmt.Println("The answer for puzzle 2 is: ", bag_in_bags_count)
}
