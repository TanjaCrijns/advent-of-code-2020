package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func binary_search(number int, seat string, up string, down string) int {
	low := 0
	high := number

	for i, letter := range seat {
		if string(letter) == down {
			high -= int(math.Ceil(float64((high - low)) / float64(2)))
			if i == len(seat)-1 {
				return low
			}
		}
		if string(letter) == up {
			low += int(math.Ceil((float64(high - low)) / float64(2)))
			if i == len(seat)-1 {
				return high
			}
		}
	}
	return 0
}

func get_row_column(seat string) (int, int) {
	row := binary_search(127, seat[0:len(seat)-3], "B", "F")
	column := binary_search(7, seat[len(seat)-3:], "R", "L")
	return row, column
}

func main() {
	file_content, _ := ioutil.ReadFile("input.txt")
	input_list := strings.Split(string(file_content), "\n")

	// Puzzle 1
	var seats [128 * 8]bool
	var min, max int
	for _, input := range input_list {
		row, column := get_row_column(input)
		seat_id := (row * 8) + column
		if seat_id < min || min == 0 {
			min = seat_id
		}
		if seat_id > max {
			max = seat_id
		}
		seats[seat_id] = true
	}
	// Puzzle 2
	var my_id int
	for id, seat_taken := range seats[min+1 : max] {
		if !seat_taken {
			my_id = min + 1 + id
		}
	}
	fmt.Println("Highest ID: ", max)
	fmt.Println("My ID: ", my_id)
}
