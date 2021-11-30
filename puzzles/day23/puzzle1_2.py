from tqdm import tqdm

# Crab cups with lists
# def crab_cups(cup_list, n_moves):
# 	current_cup_index = 0
# 	cup_modulo = len(cup_list)
# 	for i in tqdm(range(n_moves)):
# 		current_cup = cup_list[current_cup_index]
# 		pick_up = [cup_list[(current_cup_index+1)%cup_modulo], cup_list[(current_cup_index+2)%cup_modulo], cup_list[(current_cup_index+3)%cup_modulo]]
# 		new_list = [x for x in cup_list if x not in pick_up]
# 		destination_found = False
# 		for i in range(1, (current_cup-min(new_list)) + 1):
# 			if current_cup-i in new_list:
# 				destination = current_cup-i
# 				destination_found = True
# 				break
# 		if not destination_found:
# 			destination = max(new_list)
# 		destination_index = new_list.index(destination)
# 		new_list = new_list[:destination_index+1] + pick_up + new_list[destination_index+1:]
# 		current_cup_index = new_list.index(current_cup)
# 		current_cup_index = (current_cup_index+1)%cup_modulo
# 		cup_list = new_list
# 	return cup_list 

# Crab cups with dictionary
def crab_cups(cup_list, n_moves):
	cup_dict = {}
	for i, cup in enumerate(cup_list):
		if i < len(cup_list)-1:
			cup_dict[cup] = cup_list[i+1]
		else:
			cup_dict[cup] = cup_list[0]

	cup_modulo = len(cup_list) + 1
	current_cup = cup_list[0]
	for i in tqdm(range(n_moves)):
		pick_up = []
		for _ in range(3):
			pick_up.append(cup_dict[current_cup])
			cup_dict[current_cup] = cup_dict[cup_dict[current_cup]]
		destination_found = False
		i = 1
		while not destination_found:
			if current_cup-i == 0:
				i += 1
			if (current_cup-i)%cup_modulo in pick_up:
				i += 1
			else:
				destination_found = True
				cup_dict[pick_up[2]] = cup_dict[(current_cup-i)%cup_modulo]
				cup_dict[(current_cup-i)%cup_modulo] = pick_up[0]
		current_cup = cup_dict[current_cup]
	return cup_dict


if __name__ == "__main__":
	input = [7,9,2,8,4,5,1,3,6]	

	n_moves = 100
	cup_dict_puzzle_1 = crab_cups(input, n_moves)
	answer_puzzle_1 = ""
	current = 1
	for i in range(len(cup_dict_puzzle_1)-1):
		answer_puzzle_1 += str(cup_dict_puzzle_1[current])
		current = cup_dict_puzzle_1[current]
	
	print("The answer to puzzle 1 is: ", answer_puzzle_1)

	input = input + list(range(max(input)+1, 1000001))
	n_moves = 10000000
	cup_dict_puzzle_2 = crab_cups(input, n_moves)
	current = 1
	first = cup_dict_puzzle_2[current]
	second = cup_dict_puzzle_2[first]
	answer_puzzle_2 = first * second
	print("The answer to puzzle 2 is: ", answer_puzzle_2)
