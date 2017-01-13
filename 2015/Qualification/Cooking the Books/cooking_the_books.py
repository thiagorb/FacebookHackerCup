def swap0(string, index):
    if index == 0:
        return string

    return string[index] + string[1: index] + string[0] + string[index + 1:]

t = int(raw_input())

for j in xrange(t):
    n = raw_input()
    lowest = 0
    highest = 0
    for i in xrange(1, len(n)):
        if n[i] != '0' and n[i] < n[lowest]:
            lowest = i

        if n[i] > n[highest]:
            highest = i

    print("Case #%i: %s %s" % (j + 1, swap0(n, lowest), swap0(n, highest)))
