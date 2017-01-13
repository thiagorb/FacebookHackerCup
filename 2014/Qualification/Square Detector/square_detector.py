def width(line):
    first = line.find("#")
    last = line.rfind("#")
    w = last - first + 1
    if line[first:last + 1] != "#" * w:
        return None
    return w

def solve_case():
    n = int(input())
    white = "." * n
    square_line = None
    square_start = None
    square_end = None
    state = 0
    for j in range(n):
        line = input()
        if state == 0:
            if line != white:
                state = 1
                square_line = line
                square_start = j
                square_width = width(line)
                if square_width == None:
                    state = 5
        elif state == 1:
            if line != square_line:
                if line == white:
                    state = 2
                    square_end = j
                else:
                    state = 3
        elif state == 2:
            if line != white:
                state = 3
    if state == 3:
        return False
    if state == 1:
        square_end = n
    return square_end - square_start == square_width

def solve():
    t = int(input())
    output = open("out.txt", "wb")
    for i in range(t):
        output.write(bytes("Case #%i: %s\n" % (i + 1, "YES" if solve_case() else "NO"), "UTF-8"))

solve()
