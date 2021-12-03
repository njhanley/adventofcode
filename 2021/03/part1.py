import re


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = parse("input.txt", r"([01]+)", lambda a: a)

bits = len(data[0])
mcb = "".join(
    "1" if sum(s[i] == "1" for s in data) > len(data) / 2 else "0" for i in range(bits)
)
gamma, epsilon = int(mcb, 2), int("".join({"0": "1", "1": "0"}[b] for b in mcb), 2)
print(gamma * epsilon)
