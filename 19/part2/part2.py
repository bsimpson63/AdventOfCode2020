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


simple = """42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31 | 42 11 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42 | 42 8
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1
"""

simple_tests = """abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba
"""

def get_rules(rules_string):
    rules = {}
    for line in rules_string.split("\n"):
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

        if len(already_replaced) == len(rules) - 2:
            # 2 rules have loops (8, 11) so we can't solve those
            break
        elif counter > len(rules):
            print("too many loops")
            break
    return rules


def part2_simple():
    rules = get_rules(simple)
    """
    0: 8 11

    8: 42 | 42 8
    11: 42 31 | 42 11 31

    31, 42 are stable

    8 loops like:
    42, 42 42, 42 42 42, 42 42 42, etc

    11 loops like:
    42 31, 42 42 31 31, 42 42 42 31 31 31, etc

    so 8 11 looks like
    42 42 42 ... 42 31 31 31 .. 31

    if there are N 31's there must be at least N+1 42's

    """

    print("42")
    forty_two = []
    for i, match in enumerate(set(tuple(match) for match in rules[42])):
        match = "".join(match)
        print("%s: %s" % (i, match))
        forty_two.append(match)
    
    assert all(len(m) == len(forty_two[0]) for m in forty_two)

    print("31")
    thirty_one = []
    for i, match in enumerate(set(tuple(match) for match in rules[31])):
        match = "".join(match)
        print("%s: %s" % (i, match))
        thirty_one.append(match)

    assert all(len(m) == len(thirty_one[0]) for m in thirty_one)

    def count_starting_42s(test):
        i = 0
        count = 0
        while i < len(test):
            if not test[i:i + len(forty_two[0])] in forty_two:
                return count
            count += 1
            i += len(forty_two[0])
        return count


    def count_ending_31s(test):
        i = len(test)
        count = 0
        while i > 0:
            if not test[i - len(thirty_one[0]):i] in thirty_one:
                return count
            count += 1
            i -= len(thirty_one[0])
        return count

    # check that the string starts with 42 42 and ends with 31!
    match_count = 0
    for test in simple_tests.split("\n"):
        if not test:
            continue

        forty_two_count = count_starting_42s(test)
        thirty_one_count = count_ending_31s(test)
        length = len(forty_two[0]) * forty_two_count + len(thirty_one[0]) * thirty_one_count
        print(forty_two_count, thirty_one_count, length, len(test))

        i = 0
        if not (test[:len(forty_two[0])] in forty_two and test[len(forty_two[0]):2*len(forty_two[0])] in forty_two):
            print("%s NO MATCH (doesn't start with 2 42)" % test)
            continue

        i += 2 * len(forty_two[0])

        checking_forty_two = True

        while i < len(test):
            if checking_forty_two:
                if test[i:i+len(forty_two[0])] in forty_two:
                    i += len(forty_two[0])
                else:
                    checking_forty_two = False
            else:
                if test[i:i+len(thirty_one[0])] in thirty_one:
                    i += len(thirty_one[0])
                else:
                    print("%s NO MATCH (doesn't end with 31)" % test)
                    break
        else:
            if checking_forty_two:
                # never got to check for 31, not a match!
                print("%s NO MATCH (only 42)" % test)
            else:
                if forty_two_count <= thirty_one_count:
                    print("%s NO MATCH (42 count <= 31 count)" % test)
                else:
                    print("%s MATCH" % test)
                    match_count += 1
    print("%s matches" % match_count)

def part2():
    with open("./input.txt") as f:
        txt = f.read()

    final_rules, tests = txt.split("\n\n")

    rules = get_rules(final_rules)

    """
    0: 8 11

    8: 42 | 42 8
    11: 42 31 | 42 11 31

    31, 42 are stable

    8 loops like:
    42, 42 42, 42 42 42, 42 42 42, etc

    11 loops like:
    42 31, 42 42 31 31, 42 42 42 31 31 31, etc

    so 8 11 looks like
    42 42 42 ... 42 31 31 31 .. 31

    """

    print("42")
    forty_two = []
    for i, match in enumerate(set(tuple(match) for match in rules[42])):
        match = "".join(match)
        print("%s: %s" % (i, match))
        forty_two.append(match)
    
    assert all(len(m) == len(forty_two[0]) for m in forty_two)

    print("31")
    thirty_one = []
    for i, match in enumerate(set(tuple(match) for match in rules[31])):
        match = "".join(match)
        print("%s: %s" % (i, match))
        thirty_one.append(match)

    assert all(len(m) == len(thirty_one[0]) for m in thirty_one)

    def count_starting_42s(test):
        i = 0
        count = 0
        while i < len(test):
            if not test[i:i + len(forty_two[0])] in forty_two:
                return count
            count += 1
            i += len(forty_two[0])
        return count


    def count_ending_31s(test):
        i = len(test)
        count = 0
        while i > 0:
            if not test[i - len(thirty_one[0]):i] in thirty_one:
                return count
            count += 1
            i -= len(thirty_one[0])
        return count

    # check that the string starts with 42 42 and ends with 31!
    match_count = 0
    for test in tests.split("\n"):
        if not test:
            continue

        forty_two_count = count_starting_42s(test)
        thirty_one_count = count_ending_31s(test)
        length = len(forty_two[0]) * forty_two_count + len(thirty_one[0]) * thirty_one_count
        print(forty_two_count, thirty_one_count, length, len(test))

        i = 0
        if not test[i:i + len(forty_two[0])] in forty_two:
            print("%s NO MATCH (doesn't start with 1 42)" % test)
            continue

        i += len(forty_two[0])

        if not test[i:i + len(forty_two[0])] in forty_two:
            print("%s NO MATCH (doesn't start with 2 42)" % test)
            continue

        i += len(forty_two[0])

        checking_forty_two = True

        while i < len(test):
            if checking_forty_two:
                if test[i:i+len(forty_two[0])] in forty_two:
                    i += len(forty_two[0])
                else:
                    checking_forty_two = False
            else:
                if test[i:i+len(thirty_one[0])] in thirty_one:
                    i += len(thirty_one[0])
                else:
                    print("%s NO MATCH (doesn't end with 31)" % test)
                    break
        else:
            if checking_forty_two:
                # never got to check for 31, not a match!
                print("%s NO MATCH (only 42)" % test)
            else:
                if forty_two_count <= thirty_one_count:
                    print("%s NO MATCH (42 count <= 31 count)" % test)
                else:
                    print("%s MATCH" % test)
                    match_count += 1
    print("%s matches" % match_count)
    # 151 too low
    # 302 too low
    # 377 too low
    # 410
    # 393



if __name__ == "__main__":
    #part2_simple()
    part2()
    """
    print(flatten_match([4, 1, 5]))
    print(flatten_match([4, [[2, 3]], 5]))
    print(flatten_match([4, [[2, 3], [3, 3]], 5]))
    """