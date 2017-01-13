var step = (condition, iteration, resolve, reject) =>
    Promise.resolve(condition()).then(
        shouldContinue => {
            if (!shouldContinue) {
                return resolve();
            }

            Promise.resolve(iteration()).then(
                () => step(condition, iteration, resolve, reject),
                reject
            );
        },
        reject
    );

module.exports = (condition, iteration) => new Promise((resolve, reject) =>
    step(
        condition,
        iteration,
        resolve,
        reject
    )
);