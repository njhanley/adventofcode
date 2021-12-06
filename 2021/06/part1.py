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

for _ in range(80):
    new = Counter()
    for timer, count in fish.items():
        if timer == 0:
            new[6] += count
            new[8] += count
        else:
            new[timer - 1] += count
    fish = new

print(sum(fish.values()))
