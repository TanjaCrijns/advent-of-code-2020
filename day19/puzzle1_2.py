import re
import numpy as np
from tqdm import tqdm

def is_number(s):
    try:
        float(s)
        return True
    except ValueError:
        return False

def split_rule(all_rules, rule_id, recursive):
    rule = all_rules[rule_id]
    if rule_id == "8" and recursive:
        return ['('] + ['42'] + [')+']
    elif rule_id == "11" and recursive:
        return ['('] + ['42', '31'] + ['|'] + ['('] + ['['] + ['42'] + [']'] + ['{'] + ['31'] + ['}'] + [')'] + [')']
    elif rule == '"a"':
        return ["a"]
    elif rule == '"b"':
        return ["b"]
    elif "|" in rule:
        return ['('] + rule.split(' ') + [')']
    else:
        return rule.split(' ')

def possibility_regex(all_rules, recursive=False):
    done = False
    curr_rule = split_rule(all_rules, '0', recursive)
    while not done:
        next_rule = ""
        for i, poss_rule in enumerate(curr_rule):
            if is_number(poss_rule):
                next_rule = curr_rule[:i]
                next_rule += split_rule(all_rules, poss_rule, recursive)
                next_rule += curr_rule[i+1:]
                break
        if next_rule == "":
            done = True
        else:
            curr_rule = next_rule
    return ''.join(curr_rule)

def is_match(line, poss):
    if re.fullmatch(poss, line):
        return True
    return False

if __name__ == "__main__":
    with open('input.txt') as f:
        input = f.read().splitlines()

    split_index = input.index('')
    rules = input[:split_index]
    lines = input[split_index+1:]
    max_len = np.max([len(x) for x in lines])
    all_rules = {}
    rule_ids = [x.split(':')[0] for x in rules]
    split_rules = [x.split(':')[1].strip() for x in rules]
    for i, id in enumerate(rule_ids):
        all_rules[id] = split_rules[i]
    poss_regex = possibility_regex(all_rules)
    matches = [x for x in tqdm(lines) if is_match(x, poss_regex)]
    print("The answer to puzzle 1 is: ",  len(matches))

    all_rules['8'] = "42 | 42 8"
    all_rules['11'] = "42 31 | 42 11 31"
    poss_regex = possibility_regex(all_rules, True)
    first_poss_regex_left = poss_regex.index('[')
    first_poss_regex_right = poss_regex.index(']')
    second_poss_regex_left = poss_regex.index('{')
    second_poss_regex_right = poss_regex.index('}')
    temp_regex = poss_regex[:first_poss_regex_left]
    for i in range(1, max_len):
        temp_regex += '(' + poss_regex[first_poss_regex_left+1:first_poss_regex_right] + '{' + f'{i}' + '}' + poss_regex[second_poss_regex_left+1:second_poss_regex_right] + '{' + f'{i}' + '}' + ')' + '|'
    temp_regex = temp_regex[:-1]
    temp_regex += poss_regex[second_poss_regex_right+1:]
    matches = [x for x in tqdm(lines) if is_match(x, temp_regex)]
    print("The answer to puzzle 2 is: ",  len(matches))