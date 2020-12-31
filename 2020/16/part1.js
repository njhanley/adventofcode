"use strict";

const fs = require("fs");

Array.prototype.sum = function (callback) {
    return this.reduce(
        (accumulator, element, index, array) =>
            accumulator + (callback?.(element, index, array) ?? element),
        0
    );
};

Array.prototype.zip = function (...rest) {
    return this.map((element, index) => [
        element,
        ...rest.map(array => array[index]),
    ]);
};

Object.prototype.pipe = function (callback) {
    return callback(this);
};

String.prototype.matchMap = function (regexp, callback) {
    return regexp.global
        ? Array.from(this.matchAll(regexp), callback)
        : callback(this.match(regexp));
};

function range(start, stop, step = 1) {
    return Array.from(
        { length: (stop - start) / step + 1 },
        (_, index) => start + index * step
    );
}

fs.readFileSync("input.txt", "utf-8")
    .pipe(input => ({
        rules: input.matchMap(/([\w ]+): (\d+)-(\d+) or (\d+)-(\d+)/g, $ => [
            $[1],
            range(+$[2], +$[3]),
            range(+$[4], +$[5]),
        ]),
        your: input.matchMap(/your ticket:\n([\d,]+)/, $ =>
            $[1].split(",").map(x => +x)
        ),
        nearby: input.matchMap(/nearby tickets:\n([\d,\n]+)/, $ =>
            $[1].matchMap(/([\d,]+)/g, $ => $[1].split(",").map(x => +x))
        ),
    }))
    .pipe(({ rules, nearby }) => ({
        valid: new Set(rules.flatMap(([, a, b]) => a.concat(b))),
        nearby,
    }))
    .pipe(({ valid, nearby }) =>
        nearby.sum(ticket =>
            ticket.sum(value => (valid.has(value) ? 0 : value))
        )
    )
    .pipe(console.log);
