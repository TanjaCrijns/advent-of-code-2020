package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file_content, _ := ioutil.ReadFile("input.txt")
	input_instructions := strings.Split(string(file_content), "\n")

	visited := make([]bool, len(input_instructions))
	var loop_found bool = false
	var i, acc int
	for loop_found == false {
		if !visited[i] {
			visited[i] = true
		} else {
			fmt.Println("The answer is: ", acc)
			loop_found = true
		}

		instruction_descr := input_instructions[i]
		instruction := instruction_descr[0:3]
		size_argument := len(strings.Split(string(instruction_descr), " ")[1])
		argument := instruction_descr[len(instruction_descr)-size_argument:]
		arg_int, _ := strconv.Atoi(argument)

		if instruction == "acc" {
			acc += arg_int
			i++
		}
		if instruction == "jmp" {
			i += arg_int
		}
		if instruction == "nop" {
			i++
		}
	}
}
