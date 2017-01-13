def solve(gp, gc, gf, foods):
    memory = set()

    def sub_solve(gp, gc, gf, fi):
        if ((gp, gc, gf) in memory) or fi >= len(foods):
            return False

        if gp == gc == gf == 0:
            return True

        for i in xrange(fi, len(foods)):
            f = foods[i]
            qty_max = min(gp / f[0], gc / f[1], gf / f[2])
            while qty_max > 0:
                if sub_solve(gp - f[0] * qty_max, gc - f[1] * qty_max, gf - f[2] * qty_max, fi + 1):
                    return True
                qty_max -= 1

                memory.add((gp, gc, gf))
        return False

    return sub_solve(gp, gc, gf, 0)

t = int(raw_input())
for x in xrange(t):
    gp, gc, gf = [int(i) for i in raw_input().split(" ")]
    n = int(raw_input())
    foods = [
        [int(i) for i in raw_input().split(" ")] for j in xrange(n)
    ]

    print("Case #%i: %s" % (x + 1, "yes" if solve(gp, gc, gf, foods) else "no"))
