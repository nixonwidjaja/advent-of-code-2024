import os
from collections import defaultdict


def read_input():
    path = os.path.dirname(os.path.realpath(__file__))
    data = open(f"{path}/input", "r").read().strip().split("\n")
    return data


data = read_input()
idx = data.index("")
first, second = data[0:idx], data[idx + 1 :]

all = defaultdict(list)
nodes = []
for line in first:
    a, b = map(int, line.split("|"))
    all[a].append(b)


def topo(graph):
    visited = set()
    order = []

    def dfs(now):
        visited.add(now)
        for i in graph[now]:
            if i not in visited:
                dfs(i)
        order.append(now)

    for now in graph:
        if now not in visited:
            dfs(now)
    return list(reversed(order))


ans = 0
for line in second:
    arr = list(map(int, line.split(",")))
    graph = {k: [v for v in all[k] if v in arr] for k in arr}
    correct = topo(graph)
    if correct != arr:
        ans += correct[len(correct) // 2]
print(ans)
