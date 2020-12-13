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
        ([waypoint, ship], [action, value]) => {
            const modulo = (a, n) => ((a % n) + n) % n;
            const rotate = ([x, y], angle) =>
                ({
                    0: [x, y],
                    90: [-y, x],
                    180: [-x, -y],
                    270: [y, -x],
                }[modulo(angle, 360)]);
            switch (action) {
                case "N":
                    waypoint[1] += value;
                    break;
                case "S":
                    waypoint[1] -= value;
                    break;
                case "E":
                    waypoint[0] += value;
                    break;
                case "W":
                    waypoint[0] -= value;
                    break;
                case "L":
                    waypoint = rotate(waypoint, value);
                    break;
                case "R":
                    waypoint = rotate(waypoint, -value);
                    break;
                case "F":
                    ship[0] += value * waypoint[0];
                    ship[1] += value * waypoint[1];
                    break;
            }
            return [waypoint, ship];
        },
        [
            [10, 1],
            [0, 0],
        ]
    )
    .pipe(([_, [x, y]]) => [x, y].sum(n => Math.abs(n)))
    .pipe(console.log);
