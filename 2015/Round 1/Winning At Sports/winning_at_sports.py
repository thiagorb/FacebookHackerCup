def solve(h, v):
    stressful = [[None] * (v + 1) for i in xrange(h + 1)]
    stressfree = [[None] * (v + 1) for i in xrange(h + 1)]

    for hc in xrange(h, -1, -1):
        for vc in xrange(v, -1, -1):
            if hc > vc and vc < v:
                stressful[hc][vc] = 0
            elif hc == h or vc == v:
                stressful[hc][vc] = 1
            else:
                stressful[hc][vc] = (stressful[hc + 1][vc] + stressful[hc][vc + 1]) % 1000000007

            if hc <= vc and (hc > 0 or vc > 0):
                stressfree[hc][vc] = 0
            elif hc == h or vc == v:
                stressfree[hc][vc] = 1
            else:
                stressfree[hc][vc] = (stressfree[hc + 1][vc] + stressfree[hc][vc + 1]) % 1000000007

    return str(stressfree[0][0]) + " " + str(stressful[0][0])

t = int(raw_input())
for j in xrange(t):
    h, v = [int(i) for i in raw_input().split("-")]
    print("Case %i: %s" % (j + 1, solve(h, v)))
