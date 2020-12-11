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
    .sort((a, b) => a - b)
    .pipe(a => [...a, a[a.length - 1] + 3])
    .map((n, i, a) => n - (a[i - 1] ?? 0))
    .pipe(a => [1, 2, 3].map(i => a.sum(n => n === i)))
    .pipe(a => console.log(a[0] * a[2]));
