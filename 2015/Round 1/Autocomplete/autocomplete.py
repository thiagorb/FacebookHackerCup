def solve(words):
    d = {}
    total = 0
    for w in words:
        dn = d
        min_chars = None
        chars = 0
        for c in w:
            if not c in dn:
                dn[c] = {}
            dn = dn[c]

            chars += 1
            if len(dn) == 0 and min_chars == None:
                min_chars = chars
        dn[0] = {}
        total += min_chars if min_chars != None else len(w)
    return total

t = int(raw_input())
for i in xrange(t):
    n = int(raw_input())
    words = [raw_input() for j in xrange(n)]
    print("Case %i: %s" % (i + 1, solve(words)))
