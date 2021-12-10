import re
from functools import reduce


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


left = "([{<"
right = ")]}>"
closes = dict(zip(left, right))

data = read("input.txt").split("\n")

completions = []
for line in data:
    stack = []
    for c in line:
        if c in left:
            stack.append(c)
            continue
        if c != closes[stack.pop()]:
            break
    else:
        completions.append([closes[c] for c in reversed(stack)])

scores = [reduce(lambda a, b: 5 * a + right.index(b) + 1, c, 0) for c in completions]
scores.sort()
print(scores[len(scores) // 2])
