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

Array.prototype.product = function (callback) {
    return this.accumulate(callback, (a, b) => a * b, 1);
};

Array.prototype.sum = function (callback) {
    return this.accumulate(callback, (a, b) => a + b, 0);
};

Function.prototype.memoize = function () {
    const results = new Map();
    return (...args) => {
        const key = JSON.stringify(args);
        if (results.has(key)) return results.get(key);
        const result = this(...args);
        results.set(key, result);
        return result;
    };
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
    .pipe(a => [0, ...a, a[a.length - 1] + 3])
    .reverse()
    .map((n, i, a) => [
        n,
        [1, 2, 3].filterMap(j =>
            a[i + j]?.pipe(m => {
                if (m + 3 >= n) return m;
            })
        ),
    ])
    .pipe(a => new Map(a))
    .pipe(graph => {
        const outlet = [...graph.keys()].reduce((a, b) => Math.max(a, b));
        const fn = (vertex =>
            graph.get(vertex).sum(edge => fn(edge)) || 1).memoize();
        return fn(outlet);
    })
    .pipe(console.log);
