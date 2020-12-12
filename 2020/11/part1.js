"use strict";

const fs = require("fs");

Array.prototype.accumulate = function (callback, operator, initialValue) {
    return this.reduce(
        (accumulator, element, index, array) =>
            operator(accumulator, callback?.(element, index, array) ?? element),
        initialValue
    );
};

Array.prototype.filterMap = function (callback) {
    return this.reduce((newArray, currentValue, index, array) => {
        const value = callback(currentValue, index, array);
        if (value !== undefined) newArray.push(value);
        return newArray;
    }, []);
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
    .matchMap(/[.L]+/g, match => [...match[0]])
    .pipe(a => {
        const kernel = [
            [-1, -1],
            [-1, 0],
            [-1, 1],
            [0, -1],
            [0, 1],
            [1, -1],
            [1, 0],
            [1, 1],
        ];
        while (true) {
            const changes = a.flatMap((row, i) =>
                row.filterMap((c, j) => {
                    const adjacent = kernel.sum(
                        ([n, m]) => a[i + n]?.[j + m] === "#"
                    );
                    if (c === "L" && adjacent === 0) return [i, j, "#"];
                    if (c === "#" && adjacent >= 4) return [i, j, "L"];
                })
            );
            changes.forEach(([i, j, c]) => (a[i][j] = c));
            if (!changes.length) break;
        }
        return a;
    })
    .flat()
    .sum(seat => seat === "#")
    .pipe(console.log);
