package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func split_numbers(r rune) bool {
	return r == '+' || r == '*'
}

func split_operators(r rune) bool {
	return unicode.IsDigit(r)
}

func strip_chars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}

func pairwise_eval(equation string) string {
	numbers := strings.FieldsFunc(equation, split_numbers)
	operators := strings.FieldsFunc(equation, split_operators)
	result, _ := strconv.Atoi(numbers[0])
	for i := 0; i < len(operators); i++ {
		current_op := (operators[i])
		number, _ := strconv.Atoi((numbers[i+1]))
		if current_op == "+" {
			result += number
		} else if current_op == "*" {
			result *= number
		}
	}
	return strconv.Itoa(result)
}

func advanced_math_eval(equation string) string {
	var pattern = `\d*\+\d*`
	r, _ := regexp.Compile(pattern)
	var additions_left = true
	for additions_left {
		var additions = r.FindStringIndex(equation)
		if len(additions) == 0 {
			additions_left = false
		} else {
			equation = equation[:additions[0]] + pairwise_eval(equation[additions[0]:additions[1]]) + equation[additions[1]:]
		}
	}
	return pairwise_eval(equation)
}

func evaluate(equation string, puzzle int) string {
	var pattern = `\([^\(\)]+\)`
	r, _ := regexp.Compile(pattern)
	var parenthesis_left = true
	for parenthesis_left {
		var indices = r.FindStringIndex(equation)
		if len(indices) == 0 {
			parenthesis_left = false
		} else {
			if puzzle == 1 {
				equation = equation[:indices[0]] + pairwise_eval(equation[indices[0]+1:indices[1]-1]) + equation[indices[1]:]
			} else if puzzle == 2 {
				equation = equation[:indices[0]] + advanced_math_eval(equation[indices[0]+1:indices[1]-1]) + equation[indices[1]:]
			}
		}
	}
	if puzzle == 1 {
		return pairwise_eval(equation)
	} else {
		return advanced_math_eval(equation)
	}
}

func calculate_input(input []string, puzzle int) int {
	var sum int
	for _, equation := range input {
		result, _ := strconv.Atoi(evaluate(equation, puzzle))
		sum += result
	}
	return sum
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	input_lines := strings.Split(strip_chars(string(input), " "), "\n")
	fmt.Println("The answer to puzzle 1 is: ", calculate_input(input_lines, 1))
	fmt.Println("The answer to puzzle 2 is: ", calculate_input(input_lines, 2))
}
