var whilePromise = require('./while-promise');
var reader = require('./line-reader')(process.stdin);

var current = 0;

var processCase = weights => {
    var trips = 0;
    weights.sort((a, b) => a - b);

    while (weights.length) {
        var heavier = weights.pop();
        var count = 1;

        while (weights.length && heavier * count < 50) {
            weights.shift();
            count++;
        }

        if (heavier * count >= 50) {
            trips++;
        }
    }

    console.log(`Case #${current}: ${trips}`);
};

reader.next().then(T =>
    whilePromise(
        () => current++ < T,

        () => {
            var weights = [];
            var i = 0;

            return reader.next()
                .then(N =>
                    whilePromise(
                        () => i++ < N,
                        () => reader.next().then(W => weights.push(parseInt(W)))
                    )
                    .then(() => processCase(weights))
                )
        }
    )
);