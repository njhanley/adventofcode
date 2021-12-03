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


data = parse("input.txt", r"([01]+)", lambda a: [int(b) for b in a])

bit_count = [sum(bits) for bits in zip(*data)]
gamma = bits_to_int(n >= len(data) / 2 for n in bit_count)
epsilon = bits_to_int(n < len(data) / 2 for n in bit_count)

print(gamma * epsilon)
