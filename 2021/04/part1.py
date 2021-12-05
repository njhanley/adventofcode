import re
from itertools import chain


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


data = read("input.txt").split("\n\n")
drawings = match(data[0], r"(\d+)", int)
boards = [
    [match(line, r"(\d+)", int) for line in entry.splitlines()] for entry in data[1:]
]


def is_winner(board):
    for row in board:
        if all(n is None for n in row):
            return True
    for column in zip(*board):
        if all(n is None for n in column):
            return True
    return False


def mark(board, number):
    for row in board:
        for i, n in enumerate(row):
            if n == number:
                row[i] = None


def score(board, number):
    return number * sum(n for n in chain(*board) if n is not None)


def first_winning_score(boards, drawings):
    for n in drawings:
        for board in boards:
            mark(board, n)
            if is_winner(board):
                return score(board, n)


print(first_winning_score(boards, drawings))
