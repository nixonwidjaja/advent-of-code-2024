from collections import defaultdict


def part_one():
    with open("./input", "r") as f:
        lines = f.read().split("\n")
        avails = lines[0].split(", ")
        towels = lines[2:]

    memo = defaultdict(bool)

    def dfs(towel, i):
        if i == len(towel):
            return True
        if (towel, i) in memo:
            return memo[(towel, i)]
        ans = False
        for avail in avails:
            if towel[i : i + len(avail)] == avail:
                ans = ans or dfs(towel, i + len(avail))
        memo[(towel, i)] = ans
        return ans

    count = 0
    for towel in towels:
        if dfs(towel, 0):
            count += 1
    print(count)


def part_two():
    with open("./input", "r") as f:
        lines = f.read().split("\n")
        avails = lines[0].split(", ")
        towels = lines[2:]

    memo = defaultdict(int)

    def dfs(towel, i):
        if i == len(towel):
            return 1
        if (towel, i) in memo:
            return memo[(towel, i)]
        ans = 0
        for avail in avails:
            if towel[i : i + len(avail)] == avail:
                ans += dfs(towel, i + len(avail))
        memo[(towel, i)] = ans
        return ans

    count = 0
    for towel in towels:
        count += dfs(towel, 0)
    print(count)


part_one()
part_two()
