package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file_content, _ := ioutil.ReadFile("input.txt")
	input_list := strings.Split(string(file_content), "\n\n")

	var sum_counts int
	for _, input := range input_list {
		var char_count int
		var characters [26]bool
		for _, character := range input {
			if character != 10 && !characters[character-97] {
				characters[character-97] = true
				char_count++
			}
		}
		sum_counts += char_count
	}
	fmt.Println("The answer is: ", sum_counts)
}
