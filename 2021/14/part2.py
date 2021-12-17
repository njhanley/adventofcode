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
template = data.splitlines()[0]
polymer = list(zip(template, template[1:]))
rules = dict(
    match(data, r"(\w)(\w) -> (\w)", lambda a, b, c: ((a, b), ((a, c), (c, b))))
)

pairs = Counter(polymer)
for _ in range(40):
    pairs_ = Counter()
    for pair, n in pairs.items():
        for p in rules[pair]:
            pairs_[p] += n
    pairs = pairs_

elements = Counter()
for pair, n in pairs.items():
    elements[pair[1]] += n
elements[template[0]] += 1

most_common = elements.most_common()
print(most_common[0][1] - most_common[-1][1])
