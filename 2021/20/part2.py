import re


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


class Image:
    def __init__(self, data, default):
        self.pixels = {
            (x, y): c
            for y, line in enumerate(data.splitlines())
            for x, c in enumerate(line)
        }
        self.default = default

    def get(self, x, y):
        return self.pixels.get((x, y), self.default)

    def print(self, pad=0, min_p=None, max_p=None):
        x0, y0 = min_p if min_p is not None else min(self.pixels)
        x1, y1 = max_p if max_p is not None else max(self.pixels)
        for y in range(y0 - pad, y1 + 1 + pad):
            for x in range(x0 - pad, x1 + 1 + pad):
                print(self.get(x, y), end="")
            print()

    def index(self, x, y):
        n = 0
        for y_ in range(y - 1, y + 2):
            for x_ in range(x - 1, x + 2):
                n = (n << 1) + (self.get(x_, y_) == "#")
        return n

    def enhance(self, table):
        new_pixels = {}
        x0, y0 = min(self.pixels)
        x1, y1 = max(self.pixels)
        for y in range(y0 - 1, y1 + 2):
            for x in range(x0 - 1, x1 + 2):
                new_pixels[(x, y)] = table[self.index(x, y)]
        self.pixels = new_pixels
        self.default = table[0b000000000 if self.default == "." else 0b111111111]


algorithm, data = read("input.txt").split("\n\n")
image = Image(data, ".")

for _ in range(50):
    image.enhance(algorithm)
print(sum(c == "#" for c in image.pixels.values()))
