import re
from heapq import heappush, heappop
from math import inf


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


class PriorityQueue:
    def __init__(self):
        self.heap = []

    def __len__(self):
        return len(self.heap)

    def push(self, item, priority):
        heappush(self.heap, (priority, item))

    def pop(self):
        priority, item = heappop(self.heap)
        return item


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

queue = PriorityQueue()
queue.push(initial, distance[initial])
while queue:
    current = queue.pop()
    if current == destination:
        break

    x, y = current
    neighbors = [
        node
        for dx, dy in [(1, 0), (0, 1), (-1, 0), (0, -1)]
        if (node := (x + dx, y + dy)) in distance
    ]
    for node in neighbors:
        length = distance[current] + graph[node]
        if length < distance[node]:
            distance[node] = length
            queue.push(node, length)

    del distance[current]

print(distance[destination])
