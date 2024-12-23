from collections import defaultdict


def sort_set(*a):
    return ",".join(sorted(list(a)))


def part_one():
    with open("./input", "r") as f:
        lines = f.read().split("\n")
        edges = [line.split("-") for line in lines]
    graph = defaultdict(set)
    for a, b in edges:
        graph[a].add(b)
        graph[b].add(a)
    connected = set()
    for a in graph:
        if a[0] != "t":
            continue
        for b in graph[a]:
            for c in graph[b]:
                if a in graph[c]:
                    connected.add(sort_set(a, b, c))
    print(len(connected))


def check_if_connected(v, graph):
    n = len(v)
    for i in range(n):
        for j in range(i + 1, n):
            if v[j] not in graph[v[i]]:
                return False
    return True


def all_combinations(v, graph):
    n = len(v)
    ans = []
    for i in range(1 << n):
        selected = []
        for j in range(n):
            if i & (1 << j) != 0:
                selected.append(v[j])
        if check_if_connected(selected, graph) and len(selected) > 3:
            ans.append((len(selected), sort_set(*selected)))
    return ans


def part_two():
    with open("./input", "r") as f:
        lines = f.read().split("\n")
        edges = [line.split("-") for line in lines]
    graph = defaultdict(set)
    for a, b in edges:
        graph[a].add(b)
        graph[b].add(a)
    all_connections = []
    for a in list(graph):
        v = list(graph[a]) + [a]
        all_connections.extend(all_combinations(v, graph))
    print(max(all_connections, key=lambda x: x[0]))


part_one()
part_two()
