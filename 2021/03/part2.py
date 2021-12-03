import re
from functools import reduce


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


def bits_to_int(iterator):
    return reduce(lambda a, b: (a << 1) + b, iterator)


def rating(data, criteria):
    for i, _ in enumerate(data[0]):
        bit = criteria(sum(bits[i] for bits in data), len(data) / 2)
        data = [n for n in data if n[i] == bit]
        if len(data) == 1:
            return data[0]


data = parse("input.txt", r"([01]+)", lambda a: [int(b) for b in a])

oxygen = bits_to_int(rating(data, lambda a, b: a >= b))
co2 = bits_to_int(rating(data, lambda a, b: a < b))

print(oxygen * co2)
