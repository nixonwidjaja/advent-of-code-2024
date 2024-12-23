from collections import deque


def numeric_keypad(src, dst):
    graph = [["7", "8", "9"], ["4", "5", "6"], ["1", "2", "3"], [None, "0", "A"]]
    n, m = len(graph), len(graph[0])
    q = deque()
    visited = set()
    startx, starty, endx, endy = 0, 0, 0, 0
    for x in range(n):
        for y in range(m):
            if graph[x][y] == src:
                startx, starty = x, y
            if graph[x][y] == dst:
                endx, endy = x, y
    q.append((startx, starty, ""))
    visited.add((startx, starty))
    chx = [0, 0, 1, -1]
    chy = [1, -1, 0, 0]
    code = [">", "<", "v", "^"]
    min_step = 1e9
    ans = []
    while q:
        x, y, path = q.popleft()
        if (x, y) == (endx, endy):
            if len(path) <= min_step:
                min_step = len(path)
                ans.append(path)
            else:
                break
        for i in range(4):
            newx, newy = x + chx[i], y + chy[i]
            if (
                0 <= newx < n
                and 0 <= newy < m
                and graph[newx][newy] is not None
                and (newx, newy) not in visited
            ):
                visited.add((newx, newy))
                q.append((newx, newy, path + code[i]))
    return ans


def robot_type(code):
    now = "A"
    ans = ""
    for c in code:
        ans += numeric_keypad(now, c) + "A"
        now = c
    return ans


def directional_keypad(src, dst):
    graph = [[None, "^", "A"], ["<", "v", ">"]]
    n, m = len(graph), len(graph[0])
    q = deque()
    visited = set()
    startx, starty, endx, endy = 0, 0, 0, 0
    for x in range(n):
        for y in range(m):
            if graph[x][y] == src:
                startx, starty = x, y
            if graph[x][y] == dst:
                endx, endy = x, y
    q.append((startx, starty, ""))
    visited.add((startx, starty))
    chx = [0, 0, 1, -1]
    chy = [1, -1, 0, 0]
    code = [">", "<", "v", "^"]
    while q:
        x, y, path = q.popleft()
        if (x, y) == (endx, endy):
            return path
        for i in range(4):
            newx, newy = x + chx[i], y + chy[i]
            if (
                0 <= newx < n
                and 0 <= newy < m
                and graph[newx][newy] is not None
                and (newx, newy) not in visited
            ):
                visited.add((newx, newy))
                q.append((newx, newy, path + code[i]))


def type_direction(code):
    now = "A"
    ans = ""
    for c in code:
        ans += directional_keypad(now, c) + "A"
        now = c
    return ans


def part_one():
    with open("./test", "r") as f:
        lines = f.read().split("\n")
    ans = 0
    print(robot_type(lines[0]))
    for line in lines:
        direction = type_direction(type_direction(robot_type(line)))
        print(len(direction), int(line[:-1]))
        ans += len(direction) * int(line[:-1])
    print(ans)


part_one()
