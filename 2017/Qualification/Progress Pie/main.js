var lineReader = inputStream => {
    var buffer = [];
    var waiting = null;
    var finished = false;
    var incomplete = '';

    inputStream.on('readable', () => {
        var chunk = inputStream.read();

        if (!chunk) {
            return;
        }

        buffer = buffer.concat((incomplete + chunk.toString()).split('\n'));
        incomplete = buffer.pop();

        serve();
    });

    var serve = () => {
        if (!waiting) {
            return;
        }

        var value;
        if (buffer.length) {
            value = buffer.shift();
        } else if (!finished) {
            return;
        }
        waiting(value);
        waiting = null;
    };

    inputStream.on('end', () => {
        finished = true;

        if (incomplete) {
            buffer.push(incomplete);
            incomplete = null;
        }

        serve();
    });

    return {
        next: () => {
            var promise = new Promise((resolve, reject) => {
                waiting = resolve;
            });

            serve();

            return promise;
        }
    };
};

var While = iteration => iteration().then(stop => stop? null : While(iteration));

var current = 1;
var reader = lineReader(process.stdin);

var normalizeAngle = a => (a + 2 * Math.PI) % (2 * Math.PI);

reader.next().then(T => While(() =>
    reader.next().then(line => {
        if (!line) {
            return true;
        }

        var [P, X, Y] = line.split(' ');

        var dx = X - 50;
        var dy = Y - 50;
        var d2 = dx * dx + dy * dy;
        var a = normalizeAngle(Math.PI / 2 - Math.atan2(dy, dx));
        var b = Math.PI * P / 50;

        var x = d2 < 50 * 50 && a <= b;

        console.log(`Case #${current}: ${x ? 'black' : 'white'}`);

        return current++ == T;
    })
));