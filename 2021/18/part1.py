import re
import functools


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


def snailfish(s):
    return [int(c) if c.isdigit() else c for c in s]


def unsnailfish(s):
    return "".join(map(str, s))


def add(a, b):
    return reduce(["[", *a, ",", *b, "]"])


def reduce(s):
    while explode(s) or split(s):
        pass
    return s


def isnum(x):
    return isinstance(x, int)


def explode(s):
    depth = 0
    for i, c in enumerate(s):
        if c == "[":
            depth += 1
            continue
        if c == "]":
            depth -= 1
            continue
        if c == "," and depth == 5 and isnum(s[i - 1]) and isnum(s[i + 1]):
            for j in range(i - 2, -1, -1):
                if isnum(s[j]):
                    s[j] += s[i - 1]
                    break
            for j in range(i + 2, len(s)):
                if isnum(s[j]):
                    s[j] += s[i + 1]
                    break
            s[i - 2 : i + 3] = [0]
            return True
    return False


def split(s):
    for i, c in enumerate(s):
        if isnum(c) and c > 9:
            s[i : i + 1] = ["[", c // 2, ",", (c + 1) // 2, "]"]
            return True
    return False


def magnitude(s):
    if isnum(s):
        return s
    a, b = s
    return 3 * magnitude(a) + 2 * magnitude(b)


data = read("input.txt").splitlines()
snailfishes = [snailfish(line) for line in data]
total = functools.reduce(add, snailfishes)
print(magnitude(eval(unsnailfish(total))))
