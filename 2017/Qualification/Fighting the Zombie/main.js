var whilePromise = require('./while-promise');
var reader = require('./line-reader')(process.stdin);

var current = 0;

var p = (n, f, d) => {
    if (d <= 0) {
        return 1;
    }

    var m = [[1]];
    for (var j = 1; j <= d; j++) {
        m[0][j] = 0;
    }

    for (var i = 1; i <= n; i++) {
        m[i] = [1];
        for (var j = 1; j <= d; j++) {
            m[i][j] = 0;
            for (var k = 1; k <= f; k++) {
                var ki = i - 1;
                var kj = j - k;
                if (kj < 0) {
                    m[i][j] += 1;
                } else {
                    m[i][j] += m[ki][kj];
                }
            }
            m[i][j] /= f;
        }
    }
    return m[n][d];
};

var parse = s => {
    var m = /([0-9]+)d([0-9]+)([+\-]?[0-9]+)?/.exec(s);
    return {
        n: parseInt(m[1]),
        f: parseInt(m[2]),
        s: parseInt(m[3] || 0)
    };
};

var processCase = (d, ss) => {
    var best = 0;
    ss.map(parse).forEach(s => best = Math.max(best, p(s.n, s.f, d - s.s)))
    console.log(`Case #${current}: ${best.toFixed(6)}`);
};

reader.next().then(T =>
    whilePromise(
        () => current++ < T,

        () => {
            var d, ss;

            return reader.next()
                .then(line => d = line.split(' ')[0])
                .then(() => reader.next())
                .then(line => ss = line.split(' '))
                .then(() => processCase(d, ss))
        }
    )
);