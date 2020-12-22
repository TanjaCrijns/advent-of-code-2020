package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func i_hate_go_copy(input []int) []int {
	copy_list := make([]int, len(input))
	for i, _ := range input {
		copy_list[i] = input[i]
	}
	return copy_list
}

func get_decks(input_lines []string) ([]int, []int) {
	deck_p1, deck_p2 := make([]int, 0), make([]int, 0)
	var player_2 bool = false
	for _, line := range input_lines {
		if line == "Player 2:" {
			player_2 = true
		}
		if line == "" {
			continue
		}
		if !player_2 && string(line[0]) != "P" {
			line_int, _ := strconv.Atoi(line)
			deck_p1 = append(deck_p1, line_int)
		} else if player_2 && string(line[0]) != "P" {
			line_int, _ := strconv.Atoi(line)
			deck_p2 = append(deck_p2, line_int)
		}
	}
	return deck_p1, deck_p2
}

func get_player_score(deck []int) int {
	result := 0
	for i, value := range deck {
		result += (value * (len(deck) - i))
	}
	return result
}

func combat(deck_p1, deck_p2 []int) int {
	var winning_deck []int
	var no_winner bool = true
	for no_winner {
		player_1, player_2 := deck_p1, deck_p2
		var top_card_1, top_card_2 int
		top_card_1, player_1 = player_1[0], player_1[1:]
		top_card_2, player_2 = player_2[0], player_2[1:]
		if top_card_1 > top_card_2 {
			player_1 = append(player_1, top_card_1, top_card_2)
		} else {
			player_2 = append(player_2, top_card_2, top_card_1)
		}
		deck_p1 = player_1
		deck_p2 = player_2

		if len(player_1) == 0 {
			no_winner = false
			winning_deck = deck_p2
		} else if len(player_2) == 0 {
			no_winner = false
			winning_deck = deck_p1
		}
	}
	return get_player_score(winning_deck)
}

func compare_lists(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func recursive_play(deck_p1, deck_p2 []int, game int, last_game int) (winner int, winning_deck []int) {
	round := 0
	previous_games := [][][]int{}

	for {
		round++
		player_1, player_2 := deck_p1, deck_p2

		// Check if cards are the same as in the previous rounds to prevent loops
		for _, deck_pair := range previous_games {
			previous_player_1, previous_player_2 := deck_pair[0], deck_pair[1]
			if compare_lists(previous_player_1, player_1) || compare_lists(previous_player_2, player_2) {
				return 1, player_1
			}
		}

		current_game := [][]int{i_hate_go_copy(deck_p1), i_hate_go_copy(deck_p2)}
		previous_games = append(previous_games, current_game)

		var top_card_1, top_card_2 int
		top_card_1, player_1 = player_1[0], player_1[1:]
		top_card_2, player_2 = player_2[0], player_2[1:]

		// New game of recursive combat
		if len(player_1) >= top_card_1 && len(player_2) >= top_card_2 {
			var recursive_p1, recursive_p2 []int
			for i := 0; i < top_card_1; i++ {
				recursive_p1 = append(recursive_p1, player_1[i])
			}
			for i := 0; i < top_card_2; i++ {
				recursive_p2 = append(recursive_p2, player_2[i])
			}
			last_game++
			winner, _ := recursive_play(recursive_p1, recursive_p2, game+1, last_game)

			if winner == 1 {
				player_1 = append(player_1, top_card_1, top_card_2)
			} else if winner == 2 {
				player_2 = append(player_2, top_card_2, top_card_1)
			}
			// Normal play
		} else if top_card_1 > top_card_2 {
			player_1 = append(player_1, top_card_1, top_card_2)
		} else if top_card_2 > top_card_1 {
			player_2 = append(player_2, top_card_2, top_card_1)
		}

		deck_p1 = i_hate_go_copy(player_1)
		deck_p2 = i_hate_go_copy(player_2)

		if len(player_1) == 0 {
			winner = 2
			return winner, player_2
		} else if len(player_2) == 0 {
			winner = 1
			return winner, player_1
		}
	}
}

func recursive_combat(deck_p1, deck_p2 []int) int {
	var game, last_game int = 1, 1
	_, winning_deck := recursive_play(deck_p1, deck_p2, game, last_game)
	return get_player_score(winning_deck)
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	input_lines := strings.Split(string(input), "\n")

	deck_p1, deck_p2 := get_decks(input_lines)

	fmt.Println("The answer to puzzle 1 is: ", combat(deck_p1, deck_p2))
	fmt.Println("The answer to puzzle 2 is: ", recursive_combat(deck_p1, deck_p2))
}
