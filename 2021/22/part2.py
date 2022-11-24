import re
from itertools import product


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = parse("test.txt", r"(\w+) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)", lambda a, *args: (a == "on", *map(int, args)))

def intersection(a, b):
    x0, x1 = max(a[0], b[0]), min(a[1], b[1])
    y0, y1 = max(a[2], b[2]), min(a[3], b[3])
    z0, z1 = max(a[4], b[4]), min(a[5], b[5])
    if x0 > x1 or y0 > y1 or z0 > z1:
        return None
    return (x0, x1, y0, y1, z0, z1)

def volume(a):
    return (a[1] - a[0] + 1) * (a[3] - a[2] + 1) * (a[5] - a[4] + 1)

def format_tuple(t, width=6):
    if not isinstance(t, tuple):
        return t
    return "(" + ", ".join(f"{e:{width}}" for e in t) + ")" + f" |{volume(t)}|"

#for a, b in product(data, repeat=2):
#    if a == b:
#        continue
#    a, b = a[1:], b[1:]
#    c = intersection(a, b)
#    if c is not None:
#        print(f"   {format_tuple(a)}\n&  {format_tuple(b)}\n-> {format_tuple(c)}\n")

cuboids = data
i = 0
while i < len(cuboids):
    j = i + 1
    while j < len(cuboids):
        intersection(cuboids[i][1:], cuboids[j][1:])
