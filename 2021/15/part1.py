import re
from math import inf


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = parse("input.txt", r"(\d+)", lambda a: [int(c) for c in a])
graph = {(x, y): n for y, row in enumerate(data) for x, n in enumerate(row)}

initial = (0, 0)
destination = max(graph)
distance = {node: 0 if node == initial else inf for node in graph}

unvisited = set(graph)
while unvisited:
    current = min(unvisited, key=lambda node: distance[node])
    if current == destination:
        break

    unvisited.remove(current)

    x, y = current
    adjacent = {
        node
        for dx, dy in [(1, 0), (0, 1), (-1, 0), (0, -1)]
        if (node := (x + dx, y + dy)) in unvisited
    }
    for node in adjacent:
        distance[node] = min(distance[current] + graph[node], distance[node])

print(distance[destination])
