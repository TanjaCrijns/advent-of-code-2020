package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

func check_validity(contents string, check_elements []string) bool {
	var check_count int = 0
	for _, element := range check_elements {
		if strings.Contains(string(contents), string(element)) {
			check_count++
		}
	}
	if check_count >= len(check_elements) {
		return true
	} else {
		return false
	}
}

func split_entry(contents string) []string {
	split_contents_enter := strings.Split(contents, "\n")
	var element_list []string
	for _, key := range split_contents_enter {
		split_contents_space := strings.Split(key, " ")
		element_list = append(element_list, split_contents_space...)
	}
	return element_list
}

func ContainsNumbers(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func ContainsLettersAndNumbers(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func main() {

	file_content, _ := ioutil.ReadFile("input.txt")
	string_content := string(file_content)
	input := strings.Split(string_content, "\n\n")

	check_elements := []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"}
	var count_passports int = 0
	for _, key := range input {
		if check_validity(key, check_elements) {
			split_content := split_entry(key)
			var checks_passed int = 0
			for _, element := range check_elements {
				for _, item := range split_content {
					if element == "byr:" && strings.Contains(string(item), string(element)) {
						value := strings.Split(item, ":")[1]
						int_value, _ := strconv.Atoi(value)
						if int_value >= 1920 && int_value <= 2002 {
							checks_passed++
						}
					}
					if element == "iyr:" && strings.Contains(string(item), string(element)) {
						value := strings.Split(item, ":")[1]
						int_value, _ := strconv.Atoi(value)
						if int_value >= 2010 && int_value <= 2020 {
							checks_passed++
						}
					}
					if element == "eyr:" && strings.Contains(string(item), string(element)) {
						value := strings.Split(item, ":")[1]
						int_value, _ := strconv.Atoi(value)
						if int_value >= 2020 && int_value <= 2030 {
							checks_passed++
						}
					}
					if element == "hgt:" && strings.Contains(string(item), string(element)) {
						value := strings.Split(item, ":")[1]
						end := value[len(value)-2:]
						begin := value[0 : len(value)-2]
						if end == "cm" {
							int_value, _ := strconv.Atoi(begin)
							if int_value >= 150 && int_value <= 193 {
								checks_passed++
							}
						}
						if end == "in" {
							int_value, _ := strconv.Atoi(begin)
							if int_value >= 59 && int_value <= 76 {
								checks_passed++

							}
						}
					}
					if element == "hcl:" && strings.Contains(string(item), string(element)) {
						value := strings.Split(item, ":")[1]
						if string(value[0]) == "#" && len(value) == 7 {
							end := value[1:6]
							if ContainsLettersAndNumbers(end) {
								checks_passed++
							}
						}
					}
					if element == "ecl:" && strings.Contains(string(item), string(element)) {
						value := strings.Split(item, ":")[1]
						if value == "amb" || value == "blu" || value == "brn" || value == "gry" || value == "grn" || value == "hzl" || value == "oth" {
							checks_passed++
						}
					}
					if element == "pid:" && strings.Contains(string(item), string(element)) {
						value := strings.Split(item, ":")[1]
						if len(value) == 9 && ContainsNumbers(value) {
							checks_passed++
						}
					}
				}
			}
			if checks_passed == 7 {
				count_passports++
			}
		}
	}
	fmt.Println("The answer is: ", count_passports)
}
