from collections import deque


def part_one():
    with open("./input", "r") as f:
        lines = f.read().split("\n")
        lines = [list(map(int, line.split(","))) for line in lines]
    n, m = 71, 71
    graph = [["." for _ in range(m)] for _ in range(n)]
    for y, x in lines[:1024]:
        graph[x][y] = "#"
    q = deque()
    visited = [[False for _ in range(m)] for _ in range(n)]
    q.append((0, 0, 0))
    visited[0][0] = True
    chx = [0, 0, 1, -1]
    chy = [1, -1, 0, 0]
    while q:
        x, y, step = q.popleft()
        if (x, y) == (n - 1, m - 1):
            print(step)
            break
        for i in range(4):
            newx, newy = x + chx[i], y + chy[i]
            if (
                0 <= newx < n
                and 0 <= newy < m
                and graph[newx][newy] != "#"
                and not visited[newx][newy]
            ):
                q.append((newx, newy, step + 1))
                visited[newx][newy] = True


def part_two():
    with open("./input", "r") as f:
        lines = f.read().split("\n")
        lines = [list(map(int, line.split(","))) for line in lines]
    n, m = 71, 71
    for k in range(1025, len(lines)):
        passed = False
        graph = [["." for _ in range(m)] for _ in range(n)]
        for y, x in lines[:k]:
            graph[x][y] = "#"
        q = deque()
        visited = [[False for _ in range(m)] for _ in range(n)]
        q.append((0, 0, 0))
        visited[0][0] = True
        chx = [0, 0, 1, -1]
        chy = [1, -1, 0, 0]
        while q:
            x, y, step = q.popleft()
            if (x, y) == (n - 1, m - 1):
                passed = True
                break
            for i in range(4):
                newx, newy = x + chx[i], y + chy[i]
                if (
                    0 <= newx < n
                    and 0 <= newy < m
                    and graph[newx][newy] != "#"
                    and not visited[newx][newy]
                ):
                    q.append((newx, newy, step + 1))
                    visited[newx][newy] = True
        if not passed:
            print(lines[k - 1])
            break


part_one()
part_two()
