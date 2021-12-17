import re
from functools import partial, reduce
from operator import add, eq, gt, lt, mul


def read(filename):
    with open(filename) as file:
        return file.read()


def match(string, pattern, function):
    return [function(*match.groups()) for match in re.finditer(pattern, string)]


def parse(filename, pattern, function):
    return match(read(filename), pattern, function)


int = partial(int, base=2)

data = "".join(parse("input.txt", r"([0-9A-F])", lambda x: f"{int(x, base=16):04b}"))


def parse_packet(data):
    version, type_id, data = int(data[:3]), int(data[3:6]), data[6:]
    if type_id == 4:
        content, data = parse_literal(data)
    else:
        content, data = parse_operator(data)
    return (type_id, content), data


def parse_literal(data):
    bits = ""
    continue_ = 1
    while continue_:
        continue_, nibble, data = int(data[0]), data[1:5], data[5:]
        bits += nibble
    return int(bits), data


def parse_operator(data):
    packets = []
    length_type_id, data = int(data[0]), data[1:]
    if length_type_id == 0:
        length, data = int(data[:15]), data[15:]
        rest, data = data[:length], data[length:]
        while rest:
            packet, rest = parse_packet(rest)
            packets.append(packet)
    else:
        subpackets, data = int(data[:11]), data[11:]
        for _ in range(subpackets):
            packet, data = parse_packet(data)
            packets.append(packet)
    return packets, data


def evaluate(packet):
    type_id, content = packet
    if type_id == 4:
        return content
    op = {
        0: add,
        1: mul,
        2: min,
        3: max,
        5: gt,
        6: lt,
        7: eq,
    }[type_id]
    return reduce(op, [evaluate(packet) for packet in content])


packet, _ = parse_packet(data)
print(evaluate(packet))
