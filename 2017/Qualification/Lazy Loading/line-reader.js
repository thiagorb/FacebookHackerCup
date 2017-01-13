module.exports = inputStream => {
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
            if (waiting) {
                throw 'a promise is already pending';
            }

            var promise = new Promise((resolve, reject) => {
                waiting = resolve;
            });

            serve();

            return promise;
        }
    };
};