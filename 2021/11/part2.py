import re


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


def adjacent(x, y):
    return [
        (x + dx, y + dy)
        for dx, dy in [
            (1, 0),
            (1, 1),
            (0, 1),
            (-1, 1),
            (-1, 0),
            (-1, -1),
            (0, -1),
            (1, -1),
        ]
    ]


data = parse("input.txt", r"(\d+)", list)

grid = {(x, y): int(n) for y, row in enumerate(data) for x, n in enumerate(row)}

step = 0
while True:
    step += 1

    for point in grid:
        grid[point] += 1

    flashed = set()
    while True:
        flashing = [
            point
            for point, energy in grid.items()
            if point not in flashed and energy > 9
        ]
        if not flashing:
            break
        for point in flashing:
            flashed.add(point)
            for point in adjacent(*point):
                try:
                    grid[point] += 1
                except KeyError:
                    pass

    for point in flashed:
        grid[point] = 0

    if len(flashed) == len(grid):
        break

print(step)
