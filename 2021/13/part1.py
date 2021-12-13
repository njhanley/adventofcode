import re


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = read("input.txt")
dots = set(match(data, r"(\d+),(\d+)", lambda a, b: (int(a), int(b))))
folds = match(data, r"([xy])=(\d+)", lambda a, b: (a, int(b)))


def print_dots():
    w, h = max(x for x, y in dots) + 1, max(y for x, y in dots) + 1
    print(
        "\n".join(
            "".join("#" if (x, y) in dots else "." for x in range(w)) for y in range(h)
        )
    )


for axis, n in folds[:1]:
    if axis == "x":
        dots = {(2 * n - x if x > n else x, y) for x, y in dots}
    if axis == "y":
        dots = {(x, 2 * n - y if y > n else y) for x, y in dots}

print(len(dots))
