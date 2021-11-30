package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func i_hate_go_modulo(x int, y int) int {
	if x%y < 0 {
		return y + (x % y)
	}
	return x % y
}

func abs_int(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func puzzle_1(input []string, direction_index int, index_list []string) map[string]int {
	direction_dict := map[string]int{"N": 0, "E": 0, "S": 0, "W": 0}
	for _, instruction := range input {
		action := string(instruction[0])
		amount, _ := strconv.Atoi(instruction[1:])
		if action == "N" || action == "E" || action == "S" || action == "W" {
			direction_dict[action] += amount
		} else if action == "F" {
			direction_dict[index_list[direction_index]] += amount
		} else if action == "L" {
			direction_index -= amount / 90
			direction_index = i_hate_go_modulo(direction_index, 4)
		} else if action == "R" {
			direction_index += amount / 90
			direction_index = i_hate_go_modulo(direction_index, 4)
		}
	}
	return direction_dict
}

func puzzle_2(input []string, index_list []string) map[string]int {

	boat_dict := map[string]int{"N": 0, "E": 0, "S": 0, "W": 0}
	waypoint_dict := map[string]int{"N": 1, "E": 10, "S": 0, "W": 0}

	for _, instruction := range input {
		action := string(instruction[0])
		amount, _ := strconv.Atoi(instruction[1:])
		if action == "N" || action == "E" || action == "S" || action == "W" {
			waypoint_dict[action] += amount
		} else if action == "F" {
			for _, direction := range index_list {
				boat_dict[direction] += waypoint_dict[direction] * amount
			}
		} else {
			copy_waypoint_dict := map[string]int{"N": 0, "E": 0, "S": 0, "W": 0}
			n_turns := amount / 90
			for i, direction := range index_list {
				if action == "L" {
					turn_index := i_hate_go_modulo(i+n_turns, 4)
					copy_waypoint_dict[direction] = waypoint_dict[index_list[turn_index]]
				} else if action == "R" {
					turn_index := i_hate_go_modulo(i-n_turns, 4)
					copy_waypoint_dict[direction] = waypoint_dict[index_list[turn_index]]
				}
			}
			waypoint_dict = copy_waypoint_dict
		}
	}
	return boat_dict
}

func manhattan(direction_dict map[string]int) int {
	north_south := abs_int(direction_dict["N"] - direction_dict["S"])
	east_west := abs_int(direction_dict["E"] - direction_dict["W"])
	manhattan := north_south + east_west
	return manhattan
}

func main() {
	file_content, _ := ioutil.ReadFile("input.txt")
	input := strings.Split(string(file_content), "\n")

	direction_index := 1
	var index_list = []string{"N", "E", "S", "W"}

	puzzle_1_dict := puzzle_1(input, direction_index, index_list)
	fmt.Println("The answer for puzzle 1 i: ", manhattan(puzzle_1_dict))

	puzzle_2_dict := puzzle_2(input, index_list)
	fmt.Println("The answer for puzzle 2 is: ", manhattan(puzzle_2_dict))
}
