from collections import deque


def bfs(graph, startx, starty, endx, endy):
    chx = [0, 0, 1, -1]
    chy = [1, -1, 0, 0]
    n, m = len(graph), len(graph[0])
    visited = [[False for _ in range(m)] for _ in range(n)]
    q = deque()
    q.append((startx, starty, 0))
    visited[startx][starty] = True
    while q:
        x, y, step = q.popleft()
        if (x, y) == (endx, endy):
            return step
        for i in range(4):
            newx, newy = x + chx[i], y + chy[i]
            if graph[newx][newy] != "#" and not visited[newx][newy]:
                visited[newx][newy] = True
                q.append((newx, newy, step + 1))


def part_one():
    with open("./input", "r") as f:
        lines = f.read().split("\n")
        graph = [list(line) for line in lines]
    n, m = len(graph), len(graph[0])
    startx, starty, endx, endy = 0, 0, 0, 0
    for i in range(n):
        for j in range(m):
            if graph[i][j] == "S":
                startx, starty = i, j
            if graph[i][j] == "E":
                endx, endy = i, j
    benchmark = bfs(graph, startx, starty, endx, endy)
    count = 0
    for i in range(1, n - 1):
        for j in range(1, m - 1):
            if graph[i][j] == "#":
                graph[i][j] = "."
                step = bfs(graph, startx, starty, endx, endy)
                if benchmark - step >= 100:
                    count += 1
                graph[i][j] = "#"
    print(count)


def check_bounds(n, m, x, y):
    return 1 <= x < n - 1 and 1 <= y < m - 1


def new_bfs(graph, startx, starty, endx, endy):
    chx = [0, 0, 1, -1]
    chy = [1, -1, 0, 0]
    n, m = len(graph), len(graph[0])
    visited = [[-1 for _ in range(m)] for _ in range(n)]
    q = deque()
    q.append((startx, starty, 0))
    visited[startx][starty] = 0
    while q:
        x, y, step = q.popleft()
        if (x, y) == (endx, endy):
            continue
        for i in range(4):
            newx, newy = x + chx[i], y + chy[i]
            if graph[newx][newy] != "#" and visited[newx][newy] == -1:
                visited[newx][newy] = step + 1
                q.append((newx, newy, step + 1))
    count = 0
    for x in range(n):
        for y in range(m):
            for r in range(2, 21):
                for dx in range(r + 1):
                    dy = r - dx
                    for newx, newy in {
                        (x + dx, y + dy),
                        (x + dx, y - dy),
                        (x - dx, y + dy),
                        (x - dx, y - dy),
                    }:
                        if 0 <= newx < n and 0 <= newy < m and graph[newx][newy] != "#":
                            if visited[x][y] - visited[newx][newy] >= 100 + r:
                                count += 1
    print(count)


def part_two():
    with open("./input", "r") as f:
        lines = f.read().split("\n")
        graph = [list(line) for line in lines]
    n, m = len(graph), len(graph[0])
    startx, starty, endx, endy = 0, 0, 0, 0
    for i in range(n):
        for j in range(m):
            if graph[i][j] == "S":
                startx, starty = i, j
            if graph[i][j] == "E":
                endx, endy = i, j
    benchmark = bfs(graph, startx, starty, endx, endy)
    new_bfs(graph, startx, starty, endx, endy)


part_two()
