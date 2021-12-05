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


def sign(x):
    return 1 if x > 0 else 0 if x == 0 else -1


points = Counter()
for a, b in data:
    d = (sign(b[0] - a[0]), sign(b[1] - a[1]))
    if d[0] != 0 and d[1] != 0:
        continue
    while True:
        points[a] += 1
        if a == b:
            break
        a = (a[0] + d[0], a[1] + d[1])

print(sum(n > 1 for n in points.values()))
