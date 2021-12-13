import re


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = parse("input.txt", r"(\w+)-(\w+)", lambda a, b: (a, b))

graph = {}
for a, b in data:
    graph.setdefault(a, []).append(b)
    graph.setdefault(b, []).append(a)


def paths(graph, a, b, visited=None):
    if visited is None:
        visited = set()

    if a == b:
        return 1
    if a in visited:
        return 0
    if a.islower():
        visited.add(a)

    return sum(paths(graph, a, b, visited.copy()) for a in graph[a])


print(paths(graph, "start", "end"))
