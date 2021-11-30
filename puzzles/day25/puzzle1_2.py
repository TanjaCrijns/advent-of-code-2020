if __name__ == "__main__":
    public_key_door = 11404017
    public_key_card = 13768789

    loop_size = 0
    value = 1
    for i in range(1, 20201227):
        value = (value * 7) % 20201227
        if value == public_key_door: 
            loop_size = i
            print(f'Loop size is: {loop_size}')
            break

    encryption_key = 1
    for _ in range(0, loop_size):
        encryption_key = (encryption_key * public_key_card) % 20201227

    print('The answer to puzzle 1 is: ', encryption_key)
