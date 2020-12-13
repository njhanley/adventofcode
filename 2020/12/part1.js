"use strict";

const fs = require("fs");

Array.prototype.accumulate = function (callback, operator, initialValue) {
    return this.reduce(
        (accumulator, element, index, array) =>
            operator(accumulator, callback?.(element, index, array) ?? element),
        initialValue
    );
};

Array.prototype.sum = function (callback) {
    return this.accumulate(callback, (a, b) => a + b, 0);
};

Object.prototype.pipe = function (callback) {
    return callback(this);
};

String.prototype.matchMap = function (regexp, callback) {
    return Array.from(this.matchAll(regexp), callback);
};

fs.readFileSync("input.txt", "utf-8")
    .matchMap(/(\w)(\d+)/g, match => [match[1], +match[2]])
    .reduce(
        (ship, [action, value]) => {
            const modulo = (x, n) => ((x % n) + n) % n;
            const direction = modulo(ship[2], 360);
            if (action === "F")
                action = { 0: "E", 90: "N", 180: "W", 270: "S" }[direction];
            const unit = {
                N: [0, 1, 0],
                S: [0, -1, 0],
                E: [1, 0, 0],
                W: [-1, 0, 0],
                L: [0, 0, 1],
                R: [0, 0, -1],
            }[action];
            return ship.map((x, i) => x + value * unit[i]);
        },
        [0, 0, 0]
    )
    .pipe(([x, y]) => [x, y].sum(n => Math.abs(n)))
    .pipe(console.log);
