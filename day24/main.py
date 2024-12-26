from collections import defaultdict


def part_one():
    with open("./input", "r") as f:
        lines = f.read().split("\n")
        lines = [line.split(" ") for line in lines]
    d = defaultdict(int)
    graph = defaultdict(list)
    instructions = defaultdict(list)
    for line in lines:
        if len(line) == 2:
            d[line[0][:-1]] = int(line[1])
        if len(line) == 5:
            graph[line[0]].append(line[4])
            graph[line[2]].append(line[4])
            instructions[line[4]] = tuple(line[:3])
    order = []
    visited = set()

    def topo(now):
        visited.add(now)
        for i in graph[now]:
            if i not in visited:
                topo(i)
        order.append(now)

    for i in list(graph):
        if i not in visited:
            topo(i)
    for i in reversed(order):
        if i in d:
            continue
        a, op, b = instructions[i]
        if a not in d or b not in d:
            print(a, b, d)
        else:
            if op == "AND":
                d[i] = d[a] & d[b]
            elif op == "OR":
                d[i] = d[a] | d[b]
            else:
                d[i] = d[a] ^ d[b]
    print(convert_to_int(d, "z"))


def convert_to_int(d, prefix):
    bits = []
    for i in d:
        if i[0] == prefix:
            bits.append((i, d[i]))
    bits.sort(reverse=True)
    res = "".join([str(i[1]) for i in bits])
    return int(res, 2)


def convert_to_binary(d, prefix):
    bits = []
    for i in d:
        if i[0] == prefix:
            bits.append((i, d[i]))
    bits.sort(reverse=True)
    return "".join([str(i[1]) for i in bits])


def solve(instructions, d):
    graph = defaultdict(list)


def part_two():
    with open("./input", "r") as f:
        lines = f.read().split("\n")
        lines = [line.split(" ") for line in lines]
    d = defaultdict(int)
    graph = defaultdict(list)
    instructions = defaultdict(list)
    for line in lines:
        if len(line) == 2:
            d[line[0][:-1]] = int(line[1])
        if len(line) == 5:
            graph[line[0]].append(line[4])
            graph[line[2]].append(line[4])
            instructions[line[4]] = tuple(line[:3])
    order = []
    visited = set()
    x = convert_to_int(d, "x")
    y = convert_to_int(d, "y")
    supposed = x + y

    def topo(now):
        visited.add(now)
        for i in graph[now]:
            if i not in visited:
                topo(i)
        order.append(now)

    for i in list(graph):
        if i not in visited:
            topo(i)
    for i in reversed(order):
        if i in d:
            continue
        a, op, b = instructions[i]
        if a not in d or b not in d:
            print(a, b, d)
        else:
            if op == "AND":
                d[i] = d[a] & d[b]
            elif op == "OR":
                d[i] = d[a] | d[b]
            else:
                d[i] = d[a] ^ d[b]
    z = convert_to_binary(d, "z")
    print(str(bin(supposed))[2:])
    print(z)
    n = len(order)
    for i in range(n):
        a = order[i]
        if a not in instructions:
            continue
        for j in range(i + 1, n):
            b = order[j]
            if b not in instructions or d[a] == d[b]:
                continue
            d_copy = d.copy()


part_two()
