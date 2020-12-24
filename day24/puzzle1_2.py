from tqdm import tqdm

def get_directions(direction_string):
    directions = []
    skip = False
    for i, character in enumerate(direction_string):
        if skip:
            skip = False
            continue
        if character == "e" or character == "w":
            directions.append(character)
        else: 
            directions.append(character + direction_string[i+1])
            skip = True
    return directions

def tile_renovation(lines):
    start_coords = (0,0)
    black_tiles = []
    for direction_string in lines:
        current_coords = start_coords
        directions = get_directions(direction_string)
        for direction in directions:
            if direction == "e":
                current_coords = (current_coords[0]+ 1, current_coords[1])
            elif direction == "se":
                current_coords = (current_coords[0], current_coords[1]+1)
            elif direction == "sw":
                current_coords = (current_coords[0]-1, current_coords[1]+1)
            elif direction == "w":
                current_coords = (current_coords[0]-1, current_coords[1])
            elif direction == "nw":
                current_coords = (current_coords[0], current_coords[1]-1)
            elif direction == "ne":
                current_coords = (current_coords[0]+1, current_coords[1]-1)
        if current_coords in black_tiles:
            black_tiles.remove(current_coords)
        else:
            black_tiles.append(current_coords)   
    return black_tiles

def get_neighbor_coords(coords):
    neighbors = [
        (coords[0]+1, coords[1]), # east
        (coords[0], coords[1]+1), # south east
        (coords[0]-1, coords[1]+1), # south west
        (coords[0]-1, coords[1]), # west
        (coords[0], coords[1]-1), # north west
        (coords[0]+1, coords[1]-1), # north east
        ]
    return neighbors

def get_neighbor_info(neighbors, black_tiles):
    black_count = 0
    white_tiles = set()
    for neighbor in neighbors:
        if neighbor in black_tiles:
            black_count += 1
        else:
            white_tiles.add(neighbor)
    return black_count, white_tiles

def living_art(black_tiles, n_days):
    for _ in tqdm(range(n_days)):
        black_to_remove = set()
        white_to_add = set()
        white_tiles = set()

        for tile in black_tiles:
            neighbors = get_neighbor_coords(tile)
            neighbor_black_count, white_neighbors = get_neighbor_info(neighbors, black_tiles)
            white_tiles = white_tiles | white_neighbors
            if neighbor_black_count == 0 or neighbor_black_count > 2:
                black_to_remove.add(tile)
    
        for tile in white_tiles:
            neighbors = get_neighbor_coords(tile)
            neighbor_black_count, _ = get_neighbor_info(neighbors, black_tiles)
            if neighbor_black_count == 2:
                white_to_add.add(tile)

        for tile in black_to_remove:
            black_tiles.remove(tile)
        for tile in white_to_add:
            black_tiles.add(tile)

    return len(black_tiles)

if __name__ == "__main__":
    with open("input.txt") as f:
        lines = f.read().split('\n')  
    
    n_days = 100
    black_tiles = set(tile_renovation(lines))
    print("The answer to puzzle 1 is: ", len(black_tiles))
    print("The answer to puzzle 2 is: ", living_art(black_tiles, n_days))