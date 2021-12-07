import re


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = parse("input.txt", r"(\d+)", int)
costs = [sum(abs(n - i) for n in data) for i in range(min(data), max(data) + 1)]
print(min(costs))
