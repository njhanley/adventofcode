import re
from functools import reduce


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


class Set(frozenset):
    def __repr__(self):
        return "".join(sorted(self))


int_to_wiring = {
    0: Set("abcefg"),
    1: Set("cf"),
    2: Set("acdeg"),
    3: Set("acdfg"),
    4: Set("bcdf"),
    5: Set("abdfg"),
    6: Set("abdefg"),
    7: Set("acf"),
    8: Set("abcdefg"),
    9: Set("abcdfg"),
}
wiring_to_int = {v: k for k, v in int_to_wiring.items()}

data = parse(
    "input.txt",
    r"([a-z ]+) \| ([a-z ]+)",
    lambda a, b: (
        [Set(x) for x in a.split()],
        [Set(x) for x in b.split()],
    ),
)

answer = 0
for patterns, digits in data:
    unknown = {
        pattern: [
            wiring for wiring in int_to_wiring.values() if len(wiring) == len(pattern)
        ]
        for pattern in patterns
    }

    known = {}
    while unknown:
        for wiring, candidates in unknown.copy().items():
            for w1, w2 in known.items():
                candidates = [
                    candidate
                    for candidate in candidates
                    if wiring.issubset(w1) == candidate.issubset(w2)
                    and wiring.issuperset(w1) == candidate.issuperset(w2)
                ]
            if len(candidates) == 1:
                del unknown[wiring]
                known[wiring] = candidates[0]
            else:
                unknown[wiring] = candidates

    mapping = {w1: wiring_to_int[w2] for w1, w2 in known.items()}
    answer += reduce(lambda a, b: a * 10 + b, (mapping[digit] for digit in digits))

print(answer)
