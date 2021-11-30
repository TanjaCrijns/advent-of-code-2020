package main

import (
	"fmt"
	"bufio"
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

	scanner := bufio.NewScanner(strings.NewReader(string(file_content)))
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, x)
	}

	var answer int
	for _, keyi := range result{ 
		for _, keyj := range result {
			for _, keyk := range result {
				if (keyi + keyj + keyk) == 2020 {
					answer = (keyi * keyj * keyk)

				}
			}
		}
	  }
	
	fmt.Println("The answer is: ", answer)
}