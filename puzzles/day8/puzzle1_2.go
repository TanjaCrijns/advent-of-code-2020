package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func find_loops(input_instructions []string) (bool, int) {
	var acc, i int
	visited := make([]bool, len(input_instructions))

	for true {
		if visited[len(input_instructions)-1] == true {
			return false, acc
		}
		if visited[i] {
			return true, acc
		} else {
			visited[i] = true
		}
		instruction_descr := input_instructions[i]
		instruction := instruction_descr[0:3]
		size_argument := len(strings.Split(string(instruction_descr), " ")[1])
		argument := instruction_descr[len(instruction_descr)-size_argument:]
		arg_int, _ := strconv.Atoi(argument)
		if instruction == "acc" {
			i++
			acc += arg_int
		}
		if instruction == "nop" {
			i++
		}
		if instruction == "jmp" {
			i += arg_int
		}
	}
	return true, acc
}

func main() {
	file_content, _ := ioutil.ReadFile("input.txt")
	input_instructions := strings.Split(string(file_content), "\n")

	_, solution1 := find_loops(input_instructions)
	fmt.Println("The answer to puzzle 1 is: ", solution1)

	var variations [][]string
	for i, instruction_descr := range input_instructions {
		copy_input := make([]string, len(input_instructions))
		copy(copy_input, input_instructions)

		instruction := instruction_descr[0:3]
		size_argument := len(strings.Split(string(instruction_descr), " ")[1])
		argument := instruction_descr[len(instruction_descr)-size_argument:]
		if instruction == "nop" {
			copy_input[i] = "jmp " + argument
		}
		if instruction == "jmp" {
			copy_input[i] = "nop " + argument
		}
		variations = append(variations, copy_input)
	}

	for _, variation := range variations {
		loop, solution2 := find_loops(variation)
		if loop == false {
			fmt.Println("The answer to puzzle 1 is: ", solution2)
			break
		}
	}
}
