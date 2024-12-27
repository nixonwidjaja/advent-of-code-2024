from collections import defaultdict


def part_one():
    with open("./input", "r") as f:
        lines = f.read().split("\n")
        lines = [line.split(" ") for line in lines]
    d = defaultdict(int)
    instructions = defaultdict(list)
    for line in lines:
        if len(line) == 2:
            d[line[0][:-1]] = int(line[1])
        if len(line) == 5:
            instructions[line[4]] = tuple(line[:3])

    def eval(key):
        if key in d:
            return d[key]
        a, op, b = instructions[key]
        evalA, evalB = eval(a), eval(b)
        if op == "AND":
            ans = evalA & evalB
        elif op == "OR":
            ans = evalA | evalB
        else:
            ans = evalA ^ evalB
        d[key] = ans
        return ans

    for i in instructions:
        eval(i)
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


part_one()
