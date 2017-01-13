console.log(500);

var randint = (min, max) => (((max + 1 - min) * Math.random()) | 0) + min;

for (var i = 0; i < 500; i++) {
    console.log(100);

    for (var j = 0; j < 100; j++) {
        console.log(randint(1, 100));
    }
}