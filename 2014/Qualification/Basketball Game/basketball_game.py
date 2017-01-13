class Player:
    def __init__(self, name, shot_percentage, height):
        self.draft = None
        self.name = name
        self.shot_percentage = shot_percentage
        self.height = height
        self.total_played = 0

class Team:
    def __init__(self, players, p, player_draft_match):
        self.players = [player for player in players if player_draft_match(player.draft)]
        self.playing = self.players[0: p]
        self.bench = self.players[p:]

    def passMinute(self):
        if len(self.bench) == 0:
            return
        for player in self.playing:
            player.total_played += 1
        player_in = self.bench.pop(0)
        player_out = self.playing.pop()
        self.bench.append(player_out)
        self.bench = sorted(self.bench, key = lambda player: (player.total_played, player.draft))
        self.playing.insert(0, player_in)

def parse_player(line):
    data = line.split(" ")
    return Player(data[0], int(data[1]), int(data[2]))

t = int(input())
for i in range(t):
    n, m, p = [int(j) for j in input().split(" ")]
    players = [parse_player(input()) for j in range(n)]
    players.sort(key = lambda player: player.height, reverse = True)
    players.sort(key = lambda player: player.shot_percentage, reverse = True)
    draft = 0
    for player in players:
        player.draft = draft
        draft += 1

    team1 = Team(players, p, lambda draft: 1 - draft % 2)
    team2 = Team(players, p, lambda draft: draft % 2)

    for j in range(m):
        team1.passMinute()
        team2.passMinute()

    print("Case #%i: %s" % (i + 1, " ".join(sorted(map(lambda player: player.name, team1.playing + team2.playing)))))
