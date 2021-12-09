import re


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = parse("input.txt", r"(\d+)", lambda s: [int(c) for c in s])
width, height = len(data[0]), len(data)


def height_at(x, y):
    if 0 <= x < width and 0 <= y < height:
        return data[x][y]
    return 9


def adjacent(x, y):
    return [(x + dx, y + dy) for dx, dy in [(1, 0), (0, 1), (-1, 0), (0, -1)]]


risk = sum(
    1 + n
    for x, row in enumerate(data)
    for y, n in enumerate(row)
    if all(n < height_at(x, y) for x, y in adjacent(x, y))
)

print(risk)
