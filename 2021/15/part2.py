import re
from math import inf


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


def manhattan(a, b):
    return abs(b[0] - a[0]) + abs(b[1] - a[1])


data = parse("input.txt", r"(\d+)", lambda a: [int(c) for c in a])
full_map = [
    [(n - 1 + i + j) % 9 + 1 for j in range(5) for n in row]
    for i in range(5)
    for row in data
]
graph = {(x, y): n for y, row in enumerate(full_map) for x, n in enumerate(row)}

initial = (0, 0)
destination = max(graph)
distance = {node: 0 if node == initial else inf for node in graph}

open_ = set()
open_.add(initial)
while open_:
    current = min(
        open_,
        key=lambda node: distance[node] + manhattan(node, destination),
    )
    if current == destination:
        break

    open_.remove(current)

    x, y = current
    adjacent = {
        node
        for dx, dy in [(1, 0), (0, 1), (-1, 0), (0, -1)]
        if (node := (x + dx, y + dy)) in graph
    }
    for node in adjacent:
        score = distance[current] + graph[node]
        if score < distance[node]:
            distance[node] = score
            open_.add(node)

print(distance[destination])
