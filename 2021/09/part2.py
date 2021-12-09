import re
from functools import reduce


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


def basin_size(x, y, points=None):
    if points is None:
        points = set()

    if (x, y) in points or height_at(x, y) == 9:
        return 0

    points.add((x, y))

    return 1 + sum(basin_size(x, y, points) for x, y in adjacent(x, y))


low_points = [
    (x, y)
    for x, row in enumerate(data)
    for y, n in enumerate(row)
    if all(n < height_at(x, y) for x, y in adjacent(x, y))
]

basin_sizes = [basin_size(x, y) for x, y in low_points]
basin_sizes.sort()

print(reduce(lambda a, b: a * b, basin_sizes[-3:]))
