import re


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


def sign(x):
    if x < 0:
        return -1
    if x > 0:
        return 1
    return 0


def triangle(n):
    return n * (n + 1) // 2


def x_lower_bound(x_min):
    i = 0
    while True:
        if triangle(i) >= x_min:
            return i
        i += 1


data = parse("input.txt", r"(-?\d+)..(-?\d+)", lambda a, b: (int(a), int(b)))
x_min, x_max = sorted(data[0])
y_min, y_max = sorted(data[1])

hits = []
for vel_x in range(x_lower_bound(x_min), x_max + 1):
    for vel_y in range(y_min, -y_min):
        x, y = 0, 0
        v_x, v_y = vel_x, vel_y
        while True:
            x += v_x
            y += v_y
            v_x -= sign(v_x)
            v_y -= 1
            if x_min <= x <= x_max and y_min <= y <= y_max:
                hits.append((vel_x, vel_y))
                break
            if x > x_max or y < y_min:
                break

print(len(hits))
