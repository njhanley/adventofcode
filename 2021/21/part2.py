import re
from collections import Counter
from itertools import product


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = parse("input.txt", r": (\d+)", int)

dirac_dice = list(Counter(sum(dice) for dice in product([1, 2, 3], repeat=3)).items())

# (p1 pos, p2 pos, p1 score, p2 score, turn, multiplicity)
states = [(*data, 0, 0, 0, 1)]
wins = [0, 0]
while states:
    state = states.pop()
    for roll, multiple in dirac_dice:
        pos = list(state[0:2])
        score = list(state[2:4])
        turn = state[4]
        multiplicity = state[5] * multiple
        pos[turn] = (pos[turn] - 1 + roll) % 10 + 1
        score[turn] += pos[turn]
        if score[turn] >= 21:
            wins[turn] += multiplicity
            continue
        states.append((*pos, *score, int(not turn), multiplicity))
print(max(wins))
