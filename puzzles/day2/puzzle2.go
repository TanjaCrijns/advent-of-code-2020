package main

import (
	"fmt"
    "io/ioutil"
	"log"
	"strings"
	"strconv"
)

func main() {

	file_content, err := ioutil.ReadFile("input.txt")
	
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(file_content), "\n")

	var count int
	for _, key := range input{
		split_key := strings.Split(key, " ")
		bounds := split_key[0]
		lowest, _ := strconv.Atoi(strings.Split(bounds, "-")[0])
		highest, _ := strconv.Atoi(strings.Split(bounds, "-")[1])
		letter := string(split_key[1][0])
		password := split_key[2]

		var lowest_bool bool = string(password[lowest-1]) == letter
		var highest_bool bool = string(password[highest-1]) == letter 
		if lowest_bool != highest_bool{
			count += 1
		}
	}
	
	fmt.Println("The answer is: ", count)
}