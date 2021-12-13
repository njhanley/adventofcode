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


def paths(graph, node, visited=None, can_revisit=True):
    if visited is None:
        visited = set()

    if node == "end":
        return 1
    if node in visited:
        if node == "start" or not can_revisit:
            return 0
        can_revisit = False
    if node.islower():
        visited.add(node)

    return sum(paths(graph, n, visited.copy(), can_revisit) for n in graph[node])


print(paths(graph, "start"))
