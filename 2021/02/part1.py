import re


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [
        function(match.group(), *match.groups())
        for match in re.finditer(pattern, string)
    ]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = parse("input.txt", r"(\w+) (\d+)", lambda _, a, b: (a, int(b)))

position, depth = 0, 0
for direction, value in data:
    if direction == "forward":
        position += value
        continue
    if direction == "down":
        depth += value
        continue
    if direction == "up":
        depth -= value
        continue

print(position * depth)
