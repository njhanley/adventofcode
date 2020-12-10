"use strict";

const fs = require("fs");

Array.prototype.sum = function (callback) {
    return this.reduce(
        (accumulator, element, index, array) =>
            accumulator + (callback?.(element, index, array) ?? element),
        0
    );
};

Object.prototype.pipe = function (callback) {
    return callback(this);
};

String.prototype.matchMap = function (regexp, callback) {
    return Array.from(this.matchAll(regexp), callback);
};

fs.readFileSync("input.txt", "utf-8")
    .matchMap(/\d+/g, match => +match[0])
    .pipe(data => {
        const preambleLength = 25;
        const invalid = data.slice(preambleLength).find((number, index) =>
            data
                .slice(index, index + preambleLength)
                .flatMap((x, i, a) => a.slice(i + 1).map(y => x + y))
                .pipe(a => new Set(a))
                .has(number)
                .pipe(valid => !valid)
        );
        let start = 0,
            end = 1;
        while (true) {
            const difference = invalid - data.slice(start, end).sum();
            if (difference === 0) break;
            else if (difference < 0) start++;
            else if (difference > 0) end++;
        }
        return data
            .slice(start, end)
            .pipe(a => [Math.min(...a), Math.max(...a)])
            .sum();
    })
    .pipe(console.log);
