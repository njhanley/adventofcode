import re


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return (function(match) for match in re.finditer(pattern, string))


data = read("input.txt")
depths = list(match(data, r"\d+", lambda m: int(m[0])))

counter = 0
for i, _ in enumerate(depths[1:]):
    if depths[i + 1] > depths[i]:
        counter += 1

print(counter)
