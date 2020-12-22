
def get_sides(tile):
    sides = {}
    left, right = "", ""
    split_tile = tile.split("\n")
    for i, line in enumerate(split_tile):
        left += line[0]
        right += line[-1]
        if i == 0:
            sides["top"] = line
        if i == len(split_tile)-1:
            sides["bottom"] = line
    sides['left'] = left
    sides['right'] = right

    # Add flipped/rotated sides
    sides['left_flipped'] = left[len(left)::-1] 
    sides['right_flipped'] = right[len(right)::-1] 
    sides['top_rotated'] = sides["top"][len(sides["top"])::-1] 
    sides['bottom_rotated'] = sides["bottom"][len(sides["bottom"])::-1] 
    return sides

def get_side_matches(cur_id, cur_tile, sides_list, tile_dict):
    match_dict = {}
    for id, tile in tile_dict.items():
        if id != cur_id:
            for side_cur in sides_list[:4]:
                for side_new in sides_list:
                    if cur_tile['sides'][side_cur] == tile['sides'][side_new]:
                        match_dict[side_cur] = (id, side_new)

    return match_dict

if __name__ == "__main__":
    with open("input.txt") as f:
        tiles = f.read().split('\n\n')  

    tile_dict = {}
    for tile in tiles:
        id = tile[5:9]
        tile_dict[id] = {"sides": get_sides(tile[11:])}

    sum_corners = 1
    id_list = []
    sides_list = ["left", "top", "right", "bottom", "left_flipped", "right_flipped", "top_rotated", "bottom_rotated"]
    for id, tile in tile_dict.items():
        id_list.append(id)
        matches = get_side_matches(id, tile, sides_list, tile_dict)
        tile_dict[id]['sides_match'] = matches
        if len(matches) == 2:
            sum_corners *= int(id)
    print("The answer to puzzle 1 is: ", sum_corners)
