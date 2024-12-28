from collections import defaultdict, deque
from functools import cache
from itertools import product

numeric_keypad = [["7", "8", "9"], ["4", "5", "6"], ["1", "2", "3"], [None, "0", "A"]]
directional_keypad = [
    [None, "^", "A"],
    ["<", "v", ">"],
]
chx = [0, 0, 1, -1]
chy = [1, -1, 0, 0]
dir = [">", "<", "v", "^"]


def find_start(src, graph):
    n, m = len(graph), len(graph[0])
    for x in range(n):
        for y in range(m):
            if graph[x][y] == src:
                return x, y
    return -1, -1


def bfs(src, dst, graph):
    n, m = len(graph), len(graph[0])
    startx, starty = find_start(src, graph)
    visited = defaultdict(lambda: 1e9)
    q = deque([(startx, starty, "")])
    visited[(startx, starty)] = 0
    best_step = 1e9
    ans = []
    while q:
        x, y, path = q.popleft()
        if graph[x][y] == dst:
            if len(path) <= best_step:
                ans.append(path + "A")
                best_step = len(path)
            else:
                break
        for i in range(4):
            newx, newy = x + chx[i], y + chy[i]
            if (
                0 <= newx < n
                and 0 <= newy < m
                and graph[newx][newy] is not None
                and visited[(newx, newy)] >= len(path) + 1
            ):
                visited[(newx, newy)] = len(path) + 1
                q.append((newx, newy, path + dir[i]))
    return ans


@cache
def f(src, dst, depth):
    options = bfs(src, dst, directional_keypad)
    if depth == 1:
        return len(options[0])
    ans = float("inf")
    for opt in options:
        opt = "A" + opt
        length = 0
        for i in range(len(opt) - 1):
            length += f(opt[i], opt[i + 1], depth - 1)
        ans = min(ans, length)
    return ans


def min_directional_instructions(lines, depth):
    ans = float("inf")
    for line in lines:
        line = "A" + line
        length = 0
        for i in range(len(line) - 1):
            length += f(line[i], line[i + 1], depth)
        ans = min(ans, length)
    return ans


def numeric_instructions(line):
    line = "A" + line
    ans = []
    for i in range(len(line) - 1):
        ans.append(bfs(line[i], line[i + 1], numeric_keypad))
    return ["".join(i) for i in product(*ans)]


def part_one():
    with open("./input", "r") as f:
        lines = f.read().split("\n")
    ans = 0
    for line in lines:
        numeric = numeric_instructions(line)
        ans += int(line[:-1]) * min_directional_instructions(numeric, 2)
    print(ans)


def part_two():
    with open("./input", "r") as f:
        lines = f.read().split("\n")
    ans = 0
    for line in lines:
        numeric = numeric_instructions(line)
        ans += int(line[:-1]) * min_directional_instructions(numeric, 25)
    print(ans)


part_one()
part_two()
