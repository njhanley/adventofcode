import re
from itertools import product


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = parse("input.txt", r"(\w+) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)", lambda a, *args: (a == "on", *map(int, args)))

irange = lambda a, b: range(a, b + 1)

cubes = set()
for on, x0, x1, y0, y1, z0, z1 in data:
    if not all(-50 <= n <= 50 for n in [x0, x1, y0, y1, z0, z1]):
        continue
    cuboid = set(product(irange(x0, x1), irange(y0, y1), irange(z0, z1)))
    if on:
        cubes |= cuboid
    else:
        cubes -= cuboid

print(len(cubes))
