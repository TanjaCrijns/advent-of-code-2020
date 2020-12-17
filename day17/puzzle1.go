package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func i_hate_go_copy(input [][][]string, grid_size int) [][][]string {
	var copy_input = make([][][]string, grid_size)
	for x := 0; x < grid_size; x++ {
		copy_input[x] = make([][]string, grid_size)
		for y := 0; y < grid_size; y++ {
			copy_input[x][y] = make([]string, grid_size)
			for z := 0; z < grid_size; z++ {
				copy_input[x][y][z] = input[x][y][z]
			}
		}
	}
	return copy_input
}

func middle_coordinates(x, y, z, width, height int, grid_size int) bool {
	var middle int = grid_size/2 - 1
	if z == middle && x >= middle && x < middle+width && y >= middle && y < middle+height {
		return true
	}
	return false
}

func make_3d_grid(grid_size int, input string, input_width int, input_heigth int) [][][]string {
	var i_input int
	var grid_3d = make([][][]string, grid_size)
	for x := 0; x < grid_size; x++ {
		grid_3d[x] = make([][]string, grid_size)
		for y := 0; y < grid_size; y++ {
			grid_3d[x][y] = make([]string, grid_size)
			for z := 0; z < grid_size; z++ {
				if middle_coordinates(x, y, z, input_width, input_heigth, grid_size) {
					grid_3d[x][y][z] = string(input[i_input])
					i_input++
				} else {
					grid_3d[x][y][z] = "."
				}
			}
		}
	}
	return grid_3d
}

func count_cubes(grid [][][]string, grid_size int) int {
	cube_count := 0
	for x := 0; x < grid_size; x++ {
		for y := 0; y < grid_size; y++ {
			for z := 0; z < grid_size; z++ {
				if grid[x][y][z] == "#" {
					cube_count++

				}
			}
		}
	}
	return cube_count
}

func valid_position(x, y, z int, grid [][][]string, grid_size int) bool {
	if x > 0 && x < grid_size && y > 0 && y < grid_size && z > 0 && z < grid_size {
		return true
	}
	return false
}

func neighbour_active_count(x, y, z int, grid [][][]string, grid_size int) int {
	var neighbour_count int
	var coord_list = []int{-1, 0, 1}
	for _, coordx := range coord_list {
		for _, coordy := range coord_list {
			for _, coordz := range coord_list {
				if coordx == 0 && coordy == 0 && coordz == 0 {
					continue
				} else {
					if valid_position(x+coordx, y+coordy, z+coordz, grid, grid_size) {
						if grid[x+coordx][y+coordy][z+coordz] == "#" {
							neighbour_count++
						}
					}
				}
			}
		}
	}
	return neighbour_count
}

func apply_cycles(grid [][][]string, n_cycles int, grid_size int) int {
	var current_state = i_hate_go_copy(grid, grid_size)
	for i := 0; i < n_cycles; i++ {
		copy_input := i_hate_go_copy(current_state, grid_size)
		for x := 0; x < grid_size; x++ {
			for y := 0; y < grid_size; y++ {
				for z := 0; z < grid_size; z++ {
					count := neighbour_active_count(x, y, z, current_state, grid_size)
					if current_state[x][y][z] == "#" {
						if count == 2 || count == 3 {
							copy_input[x][y][z] = "#"
						} else {
							copy_input[x][y][z] = "."
						}
					} else if current_state[x][y][z] == "." {
						if count == 3 {
							copy_input[x][y][z] = "#"
						} else {
							copy_input[x][y][z] = "."
						}
					}
				}
			}
		}
		current_state = copy_input
	}
	return count_cubes(current_state, grid_size)
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	input_strings_list := strings.Split(string(input), "\n")
	input_string := strings.Join(input_strings_list, "")
	input_width, input_heigth := len(input_strings_list[0]), len(input_strings_list)

	n_cycles := 6
	// make the grid big enough
	grid_size := input_width + n_cycles*3
	grid := make_3d_grid(grid_size, input_string, input_width, input_heigth)

	fmt.Println("The answer to puzzle 1 is: ", apply_cycles(grid, n_cycles, grid_size))
}
