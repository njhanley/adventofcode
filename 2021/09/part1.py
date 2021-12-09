import re


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = parse("input.txt", r"(\d+)", lambda s: [int(c) for c in s])


def height_at(x, y):
    try:
        return data[x][y]
    except IndexError:
        return 9


risk = 0
for x, row in enumerate(data):
    for y, n in enumerate(row):
        if all(
            n < height_at(x + dx, y + dy)
            for dx, dy in [(1, 0), (0, 1), (-1, 0), (0, -1)]
        ):
            risk += 1 + n

print(risk)
