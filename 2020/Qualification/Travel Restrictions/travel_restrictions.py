#!/usr/bin/env python3

from collections import deque

cases = int(input())

for case in range(1, cases + 1):
    countries = int(input())
    incoming = [i == 'Y' for i in input()]
    outgoing = [i == 'Y' for i in input()]
    trip_possible = [{i} for i in range(countries)]
    reverse_trip_possible = [{i} for i in range(countries)]

    def new_route(departure, arrival):
        trip_possible[departure].add(arrival)
        reverse_trip_possible[arrival].add(departure)

        for indirect_departure in reverse_trip_possible[departure]:
            if not arrival in trip_possible[indirect_departure]:
                new_route(indirect_departure, arrival)

        for indirect_arrival in trip_possible[arrival]:
            if not indirect_arrival in trip_possible[departure]:
                new_route(departure, indirect_arrival)

    def check_new_route(departure, arrival):
        if arrival in trip_possible[departure]:
            return

        if not outgoing[departure]:
            return

        if not incoming[arrival]:
            return

        new_route(departure, arrival)

    for current_index in range(countries - 1):
        next_index = current_index + 1

        check_new_route(current_index, next_index)
        check_new_route(next_index, current_index)

    print('Case #%d:' % (case))
    for i in range(countries):
        print(''.join(['Y' if j in trip_possible[i] else 'N' for j in range(countries)]))
