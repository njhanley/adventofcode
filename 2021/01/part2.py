import re


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = parse("input.txt", r"(\d+)", int)
sums = [a + b + c for a, b, c in zip(data, data[1:], data[2:])]
count = sum(a < b for a, b in zip(sums, sums[1:]))
print(count)
