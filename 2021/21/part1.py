import re


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = parse("input.txt", r": (\d+)", int)


def deterministic_die():
    n = 0
    while True:
        n += 1
        yield n
        n %= 100


players = [[0, position] for position in data]
die = deterministic_die()
rolls = 0
cont = True
while cont:
    for player in players:
        for _ in range(3):
            player[1] += next(die)
        rolls += 3
        player[1] = (player[1] - 1) % 10 + 1
        player[0] += player[1]
        if player[0] >= 1000:
            cont = False
            break
loser = min(players)
print(loser[0] * rolls)
