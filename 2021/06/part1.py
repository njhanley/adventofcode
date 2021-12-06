import re


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = parse("input.txt", r"(\d+)", int)

for day in range(80):
    new = []
    for i, n in enumerate(data):
        if n == 0:
            data[i] = 6
            new.append(8)
        else:
            data[i] -= 1
    data += new

print(len(data))
