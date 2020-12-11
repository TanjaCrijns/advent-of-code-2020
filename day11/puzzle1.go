package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func i_hate_go_copy(input [][]string, len int, width int) [][]string {
	copy_input := make([][]string, len)
	for i := 0; i < len; i++ {
		copy_input[i] = make([]string, width)
	}
	for i, _ := range input {
		for j, _ := range input {
			copy_input[i][j] = input[i][j]
		}
	}
	return copy_input
}

func make_grid(input []string, string_content string) [][]string {
	grid_width := len(input[0])
	grid_height := len(input)

	grid := make([][]string, grid_height)
	for i := 0; i < grid_height; i++ {
		grid[i] = make([]string, grid_width)
	}

	var width int = 0
	var height int = 0
	for _, char := range string(string_content) {
		if string(char) == string('\n') {
			height += 1
			width = 0
		} else {
			grid[height][width] = string(char)
			width += 1
		}
	}
	return grid
}

func valid_position(i int, j int, len int, width int) bool {
	if i >= 0 && j >= 0 && i <= width-1 && j <= len-1 {
		return true
	}
	return false
}

func get_n_neighbours(i int, j int, grid [][]string) int {
	var neighbour_count int
	var coord_list = []int{-1, 0, +1}
	for _, coordx := range coord_list {
		for _, coordy := range coord_list {
			if coordx == 0 && coordy == 0 {
				continue
			} else {
				if valid_position(i+coordx, j+coordy, len(grid), len(grid[0])) {
					if grid[i+coordx][j+coordy] == "#" {
						neighbour_count++

					}
				}

			}
		}
	}
	return neighbour_count
}

func count_seats(grid [][]string) int {
	var seat_count int
	for _, row := range grid {
		for _, character := range row {
			if character == "#" {
				seat_count++
			}
		}
	}
	return seat_count
}

func main() {
	file_content, _ := ioutil.ReadFile("input.txt")
	input := strings.Split(string(file_content), "\n")
	grid := make_grid(input, string(file_content))

	var change bool = true
	for change {
		copy_grid := i_hate_go_copy(grid, len(input), len(input[0]))

		change = false
		for i, _ := range grid {
			for j, _ := range grid {
				if grid[i][j] == "L" && get_n_neighbours(i, j, grid) == 0 {
					copy_grid[i][j] = "#"
					change = true
				}
				if grid[i][j] == "#" && get_n_neighbours(i, j, grid) >= 4 {
					copy_grid[i][j] = "L"
					change = true
				}
			}
		}
		grid = copy_grid
	}
	fmt.Println("The answer for puzzle 1 is: ", count_seats(grid))
}
