class Game:
    def __init__(self, m, n, lines):
        self.m = m
        self.n = n
        self.walls = [[None for j in xrange(n)] for i in xrange(m)]
        self.start = None
        self.goal = None
        self.lasers = [[[None for j in xrange(n)] for i in xrange(m)] for k in xrange(4)]

        for i in xrange(m):
            for j in xrange(n):
                if lines[i][j] == "S":
                    self.start = (i, j)
                elif lines[i][j] == "G":
                    self.goal = (i, j)
                elif lines[i][j] == "#":
                    self.walls[i][j] = True
                elif lines[i][j] == ".":
                    pass
                else:
                    self.walls[i][j] = True

        for i in xrange(m):
            for j in xrange(n):
                if lines[i][j] in [">", "v", "<", "^"]:
                    for d in xrange(4):
                        self.write_laser(i, j, [">", "v", "<", "^"].index(lines[i][j]), d)

    def write_laser(self, i, j, d0, iteration):
        di, dj = [(0, 1), (1, 0), (0, -1), (-1, 0)][(d0 + iteration) % 4]
        i += di
        j += dj

        while self.is_free(i, j):
            self.lasers[iteration][i][j] = True
            i += di
            j += dj

    def is_free(self, i, j):
        return 0 <= i < self.m and 0 <= j < self.n and not self.walls[i][j]

    def can_move_to(self, i, j, iteration):
        return self.is_free(i, j) and not self.lasers[iteration % 4][i][j]

    def is_goal(self, position):
        return self.goal == position

    def get_lasers_map(self, iteration):
        for i in xrange(self.m):
            line = ""
            for j in xrange(self.n):
                line += "-" if self.lasers[iteration][i][j] else "."
            yield line

class State:
    def __init__(self, game, i, j, iteration):
        self.game = game
        self.i = i
        self.j = j
        self.iteration = iteration

    def next_states(self):
        neighborhood = [
            (self.i, self.j + 1),
            (self.i + 1, self.j),
            (self.i, self.j - 1),
            (self.i - 1, self.j)
        ]
        return [State(self.game, i, j, self.iteration + 1) for i, j in neighborhood if self.game.can_move_to(i, j, self.iteration + 1)]

    def __eq__(self, other):
        if not isinstance(other, State):
            return False

        return self.game == other.game and self.i == other.i and self.j == other.j and self.iteration % 4 == other.iteration % 4

    def __hash__(self):
        return hash((self.game, self.i, self.j, self.iteration % 4))

def heuristic_cost_estimate(start, goal):
    dx = start[0] - goal[0]
    dy = start[1] - goal[1]
    return (dx * dx + dy * dy) ** 0.5

def solve(game):
    start = State(game, game.start[0], game.start[1], 0)
    closedset = {}
    openset = { start: True }
    came_from = {}

    g_score = {}
    g_score[start] = 0

    f_score = {}
    f_score[start] = g_score[start] + heuristic_cost_estimate((start.i, start.j), game.goal)

    while openset:
        current = min([(f_score[s], s) for s in openset])[1]

        if game.is_goal((current.i, current.j)):
            return current.iteration

        del openset[current]
        closedset[current] = True

        for neighbor in current.next_states():
            if neighbor in closedset:
                continue

            tentative_g_score = g_score[current] + 1

            if not neighbor in openset or tentative_g_score < g_score[neighbor]:
                came_from[neighbor] = current
                g_score[neighbor] = tentative_g_score
                f_score[neighbor] = g_score[neighbor] + heuristic_cost_estimate((neighbor.i, neighbor.j), game.goal)
                if not neighbor in openset:
                    openset[neighbor] = True

    return "impossible"


t = int(raw_input())

for x in xrange(t):
    m, n = [int(i) for i in raw_input().split(" ")]

    lines = [raw_input() for i in xrange(m)]

    print("Case #%i: %s" % (x + 1, solve(Game(m, n, lines))))
