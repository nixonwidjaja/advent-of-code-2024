import os
from collections import defaultdict


def read_input():
    path = os.path.dirname(os.path.realpath(__file__))
    data = open(f"{path}/input", "r").read().strip().split("\n")
    return data


data = read_input()
idx = data.index("")
first, second = data[0:idx], data[idx + 1 :]

graph = defaultdict(set)
for line in first:
    a, b = map(int, line.split("|"))
    graph[b].add(a)


def check_order(arr):
    n = len(arr)
    for i in range(n):
        for j in range(i + 1, n):
            if arr[j] in graph[arr[i]]:
                return 0
    return arr[n // 2]


ans = 0
for line in second:
    arr = list(map(int, line.split(",")))
    print(arr, check_order(arr))
    ans += check_order(arr)

print(ans)
