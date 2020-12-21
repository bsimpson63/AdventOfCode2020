#!/usr/bin/python3

from itertools import product

simple_rules = """
0: 1 2
1: "a"
2: 1 3 | 3 1
3: "b"
"""

simple_rules2 = """
0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"
"""

def flatten_match(match):
    """
    match can be a list of rules [4, 1, 5]
    or a list of rules with an embedded simple rule [4, [[2, 3]], 5]
    -> [[4, 2, 3, 5]]
    where the 1 has been replaced with [[2, 3]]
    or a list of rules with embedded OR [4 [[2, 3], [3, 3]], 5]
    where the 1 has been replaced with another match: [[2, 3], [3, 3]]
    -> [[4, 2, 3, 5], [4, 3, 3, 5]]

    """
    if not any(isinstance(word, list) for word in match):
        return [match[:]]

    flattened = []
    options = []
    for word in match:
        if isinstance(word, list):
            options.append(word)

    for vals in product(*options):
        vals = list(vals)
        m = []
        for word in match:
            if isinstance(word, list):
                val = vals.pop(0)
                m.extend(val)
            else:
                m.append(word)
        flattened.append(m)

    return flattened


def part1():
    with open("./input.txt") as f:
        txt = f.read()

    final_rules, tests = txt.split("\n\n")

    tests = tests.split("\n")

    #final_rules = simple_rules2

    rules = {}
    for line in final_rules.split("\n"):
        if not line:
            continue

        # strip out quotes--we don't need them
        line = line.replace("\"", "")
        number, rest = line.split(": ")
        parts = rest.split(" | ")

        # each rule is a list of matches
        # each match is a list of integers (rule numbers)
        # or a or b
        rule = []
        for part in parts:
            match = []
            for word in part.split():
                if word == "a" or word == "b":
                    match.append(word)
                else:
                    match.append(int(word))
            rule.append(match)

        rules[int(number)] = rule

    counter = 0
    already_replaced = {}
    while True:
        print("iteration %s" % counter)
        counter += 1

        for number in list(rules.keys()):
            if number in already_replaced:
                continue

            has_digit = False
            for match in rules[number]:
                if any(c not in ("a", "b") for c in match):
                    has_digit = True
                    break

            if has_digit:
                continue

            # print("replacing %s" % (number))
            already_replaced[number] = True

            for j in sorted(list(rules.keys())):
                if j == number or j in already_replaced:
                    continue

                rule = rules[j]
                new_rule = []
                for match in rule:
                    new_match = [c if c != number else rules[number] for c in match]
                    new_match = flatten_match(new_match)
                    new_rule.extend(new_match) 
                rules[j] = new_rule      

        if len(already_replaced) == len(rules):
            break
        elif counter > len(rules):
            print("too many loops")
            break

    zero = rules[0]
    zero = ["".join(match) for match in zero]
    """
    for i, match in enumerate(zero):
        print("%s: %s" % (i, match))
    """

    matches = sum(1 for test in tests if test in zero)
    print("%s matches" % matches)


if __name__ == "__main__":
    part1()
    """
    print(flatten_match([4, 1, 5]))
    print(flatten_match([4, [[2, 3]], 5]))
    print(flatten_match([4, [[2, 3], [3, 3]], 5]))
    """