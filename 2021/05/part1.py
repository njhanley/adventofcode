import re
from collections import Counter


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = parse(
    "input.txt",
    r"(\d+),(\d+) -> (\d+),(\d+)",
    lambda a, b, c, d: ((int(a), int(b)), (int(c), int(d))),
)

points = Counter()
for a, b in data:
    if a[0] == b[0]:
        y1, y2 = min(a[1], b[1]), max(a[1], b[1])
        for y in range(y1, y2 + 1):
            points[a[0], y] += 1
    if a[1] == b[1]:
        x1, x2 = min(a[0], b[0]), max(a[0], b[0])
        for x in range(x1, x2 + 1):
            points[x, a[1]] += 1

print(sum(crossings > 1 for crossings in points.values()))
