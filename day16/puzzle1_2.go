package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func make_range(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func intersection(a, b []int) (c []int) {
	m := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func parse_input(input string) ([]int, [][]int, []int, [][]int) {
	var departure_indices, my_ticket []int
	var rules, nearby_tickets [][]int
	// parse_index 0 = rules, 1 = my ticket, 2 = nearby tickets
	var parse_index int
	for i, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			parse_index++
		}

		// parse rules
		// rules are represented as lists of their respective ranges
		// 1-3 or 5-7 > [1,2,3,5,6,7]
		if parse_index == 0 {
			// store indices of "departure" rules for puzzle 2
			if line[:9] == "departure" {
				departure_indices = append(departure_indices, i)
			}
			rule := make([]int, 0)
			values := strings.Split(strings.Split(line, ": ")[1], " or ")
			for _, value := range values {
				value1, _ := strconv.Atoi(strings.Split(value, "-")[0])
				value2, _ := strconv.Atoi(strings.Split(value, "-")[1])
				values := make_range(value1, value2)
				rule = append(rule, values...)
			}
			rules = append(rules, rule)
		} else {
			// parse tickets into lists of integers
			ticket := make([]int, 0)
			if !strings.Contains(line, ":") {
				for _, character := range strings.Split(line, ",") {
					int_character, _ := strconv.Atoi(character)
					ticket = append(ticket, int_character)
				}
				if parse_index == 1 {
					my_ticket = ticket
				} else {
					nearby_tickets = append(nearby_tickets, ticket)
				}
			}
		}

	}
	return departure_indices, rules, my_ticket, nearby_tickets
}

func check_tickets(nearby_tickets [][]int, rules [][]int) (int, [][]int) {
	// count invalid values for puzzle 1 and simultaneously store tickets
	// that are valid for puzzle 2
	invalid_values := make([]int, 0)
	valid_tickets := make([][]int, 0)
	for _, ticket := range nearby_tickets {
		contains_invalid := false
		for _, value := range ticket {
			value_found := false
			for _, rule := range rules {
				if contains(rule, value) {
					value_found = true
				}
			}
			if !value_found {
				invalid_values = append(invalid_values, value)
				contains_invalid = true
			}
		}
		if !contains_invalid {
			valid_tickets = append(valid_tickets, ticket)
		}
	}
	return sum(invalid_values), valid_tickets
}

func my_ticket_sum(rules [][]int, my_ticket []int, valid_tickets [][]int, departure_indices []int) int {
	// map that contains all possible positions for each rule
	possible_rules := map[int][]int{}
	for _, ticket := range valid_tickets {
		for i_value, value := range ticket {
			var temp_rules []int
			for i_rule, rule := range rules {
				if contains(rule, value) {
					temp_rules = append(temp_rules, i_rule)
				}
			}
			if len(possible_rules[i_value]) > 0 {
				possible_rules[i_value] = intersection(possible_rules[i_value], temp_rules)
			} else {
				possible_rules[i_value] = temp_rules
			}
		}
	}

	// map that contains the the indices of rules, sorted by number of possible positions
	sorted_rules := map[int]int{}
	for i := 0; i < len(possible_rules); i++ {
		sorted_rules[len(possible_rules[i])] = i
	}

	// starting with the rule that only has one position, we go through the sorted rules
	// map and remove positions that we have seen before. We store this in a new map
	// that looks like {actual_position_on_ticket: starting_rule_index}
	var checked []int
	definitive_rule_sorting := map[int]int{}
	for i := 1; i < len(sorted_rules); i++ {
		if i == 1 {
			definitive_rule_sorting[sorted_rules[i]] = possible_rules[sorted_rules[i]][0]
			checked = append(checked, possible_rules[sorted_rules[i]][0])
		} else {
			for _, value := range possible_rules[sorted_rules[i]] {
				if !contains(checked, value) {
					definitive_rule_sorting[value] = sorted_rules[i]
					checked = append(checked, value)
				}
			}
		}
	}

	// combine values from my ticket that match the departure indices we calculated before
	var answer int = 1
	for _, val := range departure_indices {
		answer *= my_ticket[definitive_rule_sorting[val]]
	}
	return answer
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	departure_indices, rules, my_ticket, nearby_tickets := parse_input(string(input))
	sum_invalid_values, valid_tickets := check_tickets(nearby_tickets, rules)

	fmt.Println("The answer to puzzle 1 is: ", sum_invalid_values)
	fmt.Println("The answer to puzzle 1 is: ", my_ticket_sum(rules, my_ticket, valid_tickets, departure_indices))
}
