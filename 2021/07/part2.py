import re
from math import factorial


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


def triangle(n):
    return n * (n + 1) // 2


data = parse("input.txt", r"(\d+)", int)
costs = [
    sum(triangle(abs(n - i)) for n in data) for i in range(min(data), max(data) + 1)
]
print(min(costs))
