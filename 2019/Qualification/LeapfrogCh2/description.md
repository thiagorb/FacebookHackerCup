# Leapfrog: Ch. 2
15 points

This problem statement differs from that of Leapfrog Ch. 2 in only one spot, highlighted in bold below.

A colony of frogs peacefully resides in a pond. The colony is led by a single Alpha Frog, and also includes 0 or more Beta Frogs. In order to be a good leader, the Alpha Frog diligently studies the high art of fractions every day.

There are N lilypads in a row on the pond's surface, numbered 1 to N from left to right, each of which is large enough to fit at most one frog at a time. Today, the Alpha Frog finds itself on the leftmost lilypad, and must leap its way to the rightmost lilypad before it can begin its fractions practice.

The initial state of each lilypad i is described by a character Li, which is one of the following:

- "A": Occupied by the Alpha Frog (it's guaranteed that Li = "A" if and only if i = 1)
- "B": Occupied by a Beta Frog
- ".": Unoccupied

At each point in time, one of the following things may occur:

1) The Alpha Frog may leap over one or more lilypads immediately to either its left or right which are occupied by Beta Frogs, and land on the next unoccupied lilypad past them, if such a lilypad exists. The Alpha Frog must leap over at least one Beta Frog; it may not just leap to an adjacent lilypad. **Note that, unlike in Leapfrog Ch. 1, the Alpha Frog may leap to either its left or right**.

2) Any Beta Frog may leap to the next lilypad to either its left or right, if such a lilypad exists and is unoccupied.

Assuming the frogs all cooperate, determine whether or not it's possible for the Alpha Frog to ever reach the rightmost lilypad and begin its daily fractions practice.

## Input
Input begins with an integer T, the number of days on which the Alpha Frog studies fractions. For each day, there is a single line containing the length-N string L.

## Output
For the ith day, print a line containing "Case #i: " followed by a single character: "Y" if the Alpha Frog can reach the rightmost lilypad, or "N" otherwise.

## Constraints
- 1 ≤ T ≤ 500
- 2 ≤ N ≤ 5,000

## Explanation of Sample

In the first case, the Alpha Frog can't leap anywhere.

In the second case, the Alpha Frog can leap over the Beta Frog to reach the rightmost lilypad.

In the third case, neither the Alpha Frog nor either of the Beta Frogs can leap anywhere.

In the fourth case, if the first Beta Frog leaps one lilypad to the left, and then the second Beta Frog also leaps one lilypad to the left, then the Alpha Frog can leap over both of them to reach the rightmost lilypad.

## Sample input

    8
    A.
    AB.
    ABB
    A.BB
    A..BB..B
    A.B..BBB.
    AB.........
    A.B..BBBB.BB

## Sample output

    Case #1: N
    Case #2: Y
    Case #3: N
    Case #4: Y
    Case #5: Y
    Case #6: Y
    Case #7: N
    Case #8: Y
