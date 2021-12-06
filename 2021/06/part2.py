import re
from collections import Counter


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = parse("input.txt", r"(\d+)", int)
fish = Counter(data)

for day in range(256):
    new = Counter()
    for n, m in fish.items():
        if n == 0:
            new[6] += m
            new[8] += m
        else:
            new[n - 1] += m
    fish = new

print(sum(fish.values()))
