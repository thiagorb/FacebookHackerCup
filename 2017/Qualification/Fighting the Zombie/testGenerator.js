console.log(1000);

var randint = (min, max) => (((max + 1 - min) * Math.random()) | 0) + min;
var dices = [4, 6, 8, 10, 12, 20];

for (var i = 0; i < 1000; i++) {
    console.log('1000 10');
    var s = [];
    for (var j = 0; j < 10; j++) {
        var X = randint(1, 20);
        var Y = dices[randint(0, dices.length - 1)];
        var Z = randint(880, 890);
        s.push(`${X}d${Y}${Z >= 0? '+' : ''}${Z}`);
    }
    console.log(s.join(' '));
}