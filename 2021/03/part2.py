import re
from operator import ge, lt


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = parse("input.txt", r"([01]+)", lambda a: [b == "1" for b in a])


def find(ns, op):
    for i, _ in enumerate(ns[0]):
        bit = op(sum(n[i] for n in ns), len(ns) / 2)
        ns = [n for n in ns if n[i] == bit]
        if len(ns) == 1:
            return ns[0]


oxy = int("".join("1" if b else "0" for b in find(data, ge)), 2)
co2 = int("".join("1" if b else "0" for b in find(data, lt)), 2)
print(oxy * co2)
