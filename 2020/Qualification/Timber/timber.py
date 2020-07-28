#!/usr/bin/env python3

cases = int(input())

for case in range(1, cases + 1):
    trees_length = int(input())
    trees = sorted([tuple(map(int, input().split())) for i in range(trees_length)])
    best_overall = 0
    bests = dict()

    for tree in trees:
        right_best = (bests.get(tree[0]) or 0) + tree[1]
        left_best = (bests.get(tree[0] - tree[1]) or 0) + tree[1]

        bests[tree[0] + tree[1]] = max(bests.get(tree[0] + tree[1]) or 0, right_best)
        bests[tree[0]] = max(bests.get(tree[0]) or 0, left_best)

        best_overall = max(best_overall, right_best, left_best)

    print('Case #%d: %s' % (case, best_overall))
