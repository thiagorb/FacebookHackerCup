#!/usr/bin/env python3

from collections import Counter

cases = int(input())

for case in range(1, cases + 1):
    shards_length = int(input())
    shards = input()

    counter = Counter()
    for shard in shards:
        counter[shard] += 1

    result = 'Y' if abs(counter['A'] - counter['B']) <= 2 else 'N'
    print('Case #%d: %s' % (case, result))
