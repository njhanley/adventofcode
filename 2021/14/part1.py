import re
from collections import Counter


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = read("input.txt")
polymer = data.splitlines()[0]
rules = dict(match(data, r"(\w{2}) -> (\w)", lambda a, b: (a, b)))

for _ in range(10):
    polymer = polymer[0] + "".join(
        rules[a + b] + b for a, b in zip(polymer, polymer[1:])
    )

most_common = Counter(polymer).most_common()
print(most_common[0][1] - most_common[-1][1])
