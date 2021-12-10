import re


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = parse(
    "input.txt", r"([a-z ]+) \| ([a-z ]+)", lambda a, b: (a.split(), b.split())
)

print(
    sum(sum(1 for digit in digits if len(digit) in [2, 3, 4, 7]) for _, digits in data)
)
