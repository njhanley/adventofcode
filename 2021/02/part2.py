import re


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = parse("input.txt", r"(\w+) (\d+)", lambda a, b: (a, int(b)))

position, depth, aim = 0, 0, 0
for direction, value in data:
    if direction == "forward":
        position += value
        depth += value * aim
        continue
    if direction == "down":
        aim += value
        continue
    if direction == "up":
        aim -= value
        continue

print(position * depth)
