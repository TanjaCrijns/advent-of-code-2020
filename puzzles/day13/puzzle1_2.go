package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func i_hate_go_modulo(x int, y int) int {
	if x%y < 0 {
		return y + (x % y)
	}
	return x % y
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func puzzle1(departure_time int, bus_schedule []string) int {
	var earliest_bus, time_earliest_bus int
	for _, bus := range bus_schedule {
		if bus != "x" {
			bus_int, _ := strconv.Atoi(bus)
			next_bus_time := departure_time - (i_hate_go_modulo(departure_time, bus_int)) + bus_int
			if time_earliest_bus == 0 || next_bus_time < time_earliest_bus {
				time_earliest_bus = next_bus_time
				earliest_bus = bus_int
			}
		}
	}
	return earliest_bus * (time_earliest_bus - departure_time)
}

func puzzle2(departure_time int, bus_schedule []string) int {
	timestamp, step := 0, 1
	for i, bus := range bus_schedule {
		if bus != "x" {
			bus_int, _ := strconv.Atoi(bus)
			for i_hate_go_modulo((timestamp+i), bus_int) != 0 {
				timestamp += step
			}
			step = LCM(step, bus_int)
		}
	}
	return timestamp
}

// # Too slow brute force approach
// #
// func puzzle2(departure_time int, bus_schedule []string) int {
// 	var bus_list []int
// 	bus_dict := map[int]int{}
// 	for i, bus := range bus_schedule {
// 		if bus != "x" {
// 			bus_int, _ := strconv.Atoi(bus)
// 			bus_dict[bus_int] = i
// 			bus_list = append(bus_list, bus_int)
// 		}
// 	}

// 	found := false
// 	current_timestamp := bus_list[0]

// 	for !found {
// 		n_matched_busses := 0
// 		for _, bus := range bus_list[1:] {
// 			if i_hate_go_modulo((current_timestamp+bus_dict[bus]), bus) == 0 {
// 				n_matched_busses++
// 			} else {
// 				break
// 			}
// 		}
// 		if n_matched_busses == len(bus_dict)-1 {
// 			found = true
// 		} else {
// 			current_timestamp += bus_list[0]
// 		}
// 	}
// 	return current_timestamp
// }

func main() {
	file_content, _ := ioutil.ReadFile("input.txt")
	input := strings.Split(string(file_content), "\n")

	departure_time, _ := strconv.Atoi(input[0])
	bus_schedule := strings.Split(string(input[1]), ",")

	fmt.Println("The answer for puzzle 1 i: ", puzzle1(departure_time, bus_schedule))
	fmt.Println("The answer for puzzle 2 i: ", puzzle2(departure_time, bus_schedule))

}
