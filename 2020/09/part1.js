"use strict";

const fs = require("fs");

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
        return data.slice(preambleLength).find((number, index) =>
            data
                .slice(index, index + preambleLength)
                .flatMap((x, i, a) => a.slice(i + 1).map(y => x + y))
                .pipe(a => new Set(a))
                .has(number)
                .pipe(valid => !valid)
        );
    })
    .pipe(console.log);
