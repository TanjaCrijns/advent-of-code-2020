package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	file_content, _ := ioutil.ReadFile("input.txt")
	string_content := string(file_content)
	input := strings.Split(string_content, "\n")
	grid_width := len(input[0])
	grid_height := len(input)

	grid := make([][]string, grid_height)
	for i := 0; i < grid_height; i++ {
		grid[i] = make([]string, grid_width+1)
	}

	var width int = 0
	var height int = 0
	for _, char := range string_content {
		if string(char) == string('\n') {
			height += 1
			width = 0
		} else {
			grid[height][width] = string(char)
			width += 1
		}
	}

	var right int = 3
	var down int = 1
	var trees int = 0
	var currentpositionx int = 0
	var currentpositiony int = 0
	var maxwidth int = len(grid[0]) - 1
	i := 0
	for i < len(grid)-1 {
		if (currentpositiony + right) < maxwidth {
			currentpositiony = currentpositiony + right
		} else {
			currentpositiony = right - (maxwidth - currentpositiony)
		}
		currentpositionx = currentpositionx + down
		possible_tree := grid[currentpositionx][currentpositiony]
		if string(possible_tree) == string("#") {
			trees++
		}
		i++
	}

	fmt.Println("The answer is: ", trees)

}
