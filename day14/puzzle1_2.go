package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func apply_mask(mask string, input int) int {
	binary_input := strconv.FormatInt(int64(input), 2)
	binary_input_string := fmt.Sprintf("%0*v", len(mask), binary_input)
	for i, item := range mask {
		if string(item) == "0" {
			binary_input_string = binary_input_string[:i] + "0" + binary_input_string[i+1:]
		} else if string(item) == "1" {
			binary_input_string = binary_input_string[:i] + "1" + binary_input_string[i+1:]
		}
	}
	output, _ := strconv.ParseInt(binary_input_string, 2, 64)
	return int(output)
}

func get_values(line string) (int, int) {
	split_line := strings.Split(line, " = ")
	memory_index, _ := strconv.Atoi(split_line[0][4 : len(split_line[0])-1])
	set_value, _ := strconv.Atoi(split_line[1])
	return memory_index, set_value
}

func sum_memory_dict(memory_dict map[int]int) int {
	var answer int = 0
	for _, memory_location := range memory_dict {
		answer += memory_location
	}
	return answer
}

func puzzle1(input []string) int {
	memory_dict := make(map[int]int)
	mask := ""
	for _, line := range input {
		if line[:7] == "mask = " {
			mask = line[7:]
			continue
		} else {
			memory_index, set_value := get_values(line)
			memory_dict[memory_index] = apply_mask(mask, set_value)
		}
	}
	return sum_memory_dict(memory_dict)
}

func apply_mask2(mask string, input int) string {
	binary_input := strconv.FormatInt(int64(input), 2)
	binary_input_string := fmt.Sprintf("%0*v", len(mask), binary_input)
	for i, item := range mask {
		if string(item) == "1" {
			binary_input_string = binary_input_string[:i] + "1" + binary_input_string[i+1:]
		} else if string(item) == "X" {
			binary_input_string = binary_input_string[:i] + "X" + binary_input_string[i+1:]
		}
	}
	return binary_input_string
}

func set_values(mask_list []string, set_value int, memory_dict map[int]int) map[int]int {
	for _, mask := range mask_list {
		output, _ := strconv.ParseInt(mask, 2, 64)
		memory_dict[int(output)] = set_value
	}
	return memory_dict
}

func get_possible_masks(mask string, mask_list []string) []string {
	for i, item := range mask {
		if string(item) == "X" {
			alternative_mask_1 := mask[:i] + "0" + mask[i+1:]
			alternative_mask_2 := mask[:i] + "1" + mask[i+1:]
			return append(get_possible_masks(alternative_mask_1, mask_list), get_possible_masks(alternative_mask_2, mask_list)...)
		}
	}
	mask_list = append(mask_list, mask)
	return mask_list
}

func puzzle2(input []string) int {
	memory_dict := make(map[int]int)
	mask := ""
	for _, line := range input {
		if line[:7] == "mask = " {
			mask = line[7:]
			continue
		} else {
			var mask_list []string
			memory_index, set_value := get_values(line)
			new_mask := apply_mask2(mask, memory_index)
			mask_list = get_possible_masks(new_mask, mask_list)
			memory_dict = set_values(mask_list, set_value, memory_dict)
		}
	}
	return sum_memory_dict(memory_dict)
}

func main() {
	file_content, _ := ioutil.ReadFile("input.txt")
	input := strings.Split(string(file_content), "\n")

	fmt.Println("The answer for puzzle 1 is: ", puzzle1(input))
	fmt.Println("The answer for puzzle 2 is: ", puzzle2(input))

}
