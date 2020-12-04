package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	file_content, _ := ioutil.ReadFile("input.txt")
	string_content := string(file_content)
	input := strings.Split(string_content, "\n\n")

	check_elements := []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"}
	var count_passports int = 0
	for _, key1 := range input {
		var check_count int = 0
		for _, element := range check_elements {
			if strings.Contains(string(key1), string(element)) {
				check_count++
			}
		}
		if check_count >= len(check_elements) {
			fmt.Println(key1, "\n")
			count_passports++
		}
	}

	fmt.Println("The answer is: ", count_passports)

}
