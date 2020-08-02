console.log(1000);

var randint = (min, max) => (((max + 1 - min) * Math.random()) | 0) + min;

for (var i = 0; i < 1000; i++) {
    var P = randint(0, 100);
    var X = randint(0, 100);
    var Y = randint(0, 100);
    console.log(P, X, Y);
}