import AoCFramework as AoC

def part_1(block=None):
    global Visited
    r, c = start
    count = 1
    Visited2 = {(r, c, (-1, 0))}
    Dirns, dirptr = ((-1, 0), (0, 1), (1, 0), (0, -1)), 0
    while True:
        dr, dc = Dirns[dirptr]
        newr, newc = r + dr, c + dc
        if newr < 0 or newc < 0 or newr >= numlines or newc >= linelen:
            return count
        elif Lines[newr][newc] == '#' or (newr, newc) == block:
            dirptr += 1 if dirptr < 3 else -3
        else:
            if (newr, newc) not in Visited and block is None:
                Visited.add((newr, newc))
                count += 1
            elif block is not None:
                if (newr, newc, Dirns[dirptr]) in Visited2:
                    return 0
                else:
                    Visited2.add((newr, newc, Dirns[dirptr]))
            r, c = newr, newc

def part_2():
    return sum(1 if part_1(pos) == 0 else 0 for pos in Visited)

Lines, numlines, linelen = AoC.Init("days/day6/input/1.txt")
for r in range(numlines):
    for c in range(linelen):
        if Lines[r][c] == '^':
            start = (r, c)
            Visited = {(r, c)}

AoC.verify(4374, 1705)
AoC.run(part_1, part_2)