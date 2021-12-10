import re
from functools import reduce


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


class Pattern(frozenset):
    def __repr__(self):
        return "".join(sorted(self))


table = {
    Pattern("abcefg"): 0,
    Pattern("cf"): 1,
    Pattern("acdeg"): 2,
    Pattern("acdfg"): 3,
    Pattern("bcdf"): 4,
    Pattern("abdfg"): 5,
    Pattern("abdefg"): 6,
    Pattern("acf"): 7,
    Pattern("abcdefg"): 8,
    Pattern("abcdfg"): 9,
}

data = parse(
    "input.txt",
    r"([a-z ]+) \| ([a-z ]+)",
    lambda a, b: (
        [Pattern(x) for x in a.split()],
        [Pattern(x) for x in b.split()],
    ),
)

values = []
for patterns, digits in data:
    known = {}
    unknown = {pattern: frozenset(table) for pattern in patterns}
    while unknown:
        for pattern, candidates in unknown.copy().items():
            for candidate in candidates:
                candidates = frozenset(
                    candidate
                    for candidate in candidates
                    if len(pattern) == len(candidate)
                    and all(
                        pattern.issubset(a) == candidate.issubset(b)
                        and pattern.issuperset(a) == candidate.issuperset(b)
                        for a, b in known.items()
                    )
                )
            try:
                [known[pattern]] = candidates
                del unknown[pattern]
            except ValueError:
                unknown[pattern] = candidates

    mapping = {a: table[b] for a, b in known.items()}
    values.append(reduce(lambda a, b: a * 10 + b, (mapping[digit] for digit in digits)))

print(sum(values))
