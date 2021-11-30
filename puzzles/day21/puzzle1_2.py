if __name__ == "__main__":
    with open("input.txt") as f:
        foods = f.read().split('\n')  

    allergen_dict = {}  
    ingredients_dict = {}
    for i, food in enumerate(foods):
        ingredients, allergens = food.split(' (contains ')
        ingredients = set(ingredients.split())
        for ingredient in ingredients:
            if ingredient in ingredients_dict:
                ingredients_dict[ingredient] += 1
            else:
                ingredients_dict[ingredient] = 1

        allergens = allergens[:-1].split(', ')
        for allergen in allergens:
            if allergen in allergen_dict:
                allergen_dict[allergen] &= ingredients
            else:
                allergen_dict[allergen] = ingredients.copy()

    for values in allergen_dict.values():
        for value in values:
            if value in ingredients_dict:
                ingredients_dict.pop(value)

    print("The answer to puzzle 1 is: ", sum(ingredients_dict.values()))
    
    canon_dang_list = []
    while len(canon_dang_list) != len(allergen_dict):
        to_remove = ""
        for allergen, ingredients in allergen_dict.items():
            if len(ingredients) == 1:
                to_remove = ingredients.pop()
                canon_dang_list.append((allergen, to_remove))
                break
        for allergen, ingredients in allergen_dict.items():
            if to_remove in ingredients:
                ingredients.remove(to_remove)
    canon_dang_list.sort()
    
    print("The answer to puzzle 2 is: ", ",".join([ingredient for (allergen, ingredient) in canon_dang_list]))