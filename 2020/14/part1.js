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

fs.readFileSync("input.txt", "utf-8")
    .matchMap(/(\w+)(?:\[(\d+)\])? = ([\dX]+)/g, ([, $1, $2, $3]) => {
        if ($1 === "mask") return [$1, $3.split("")];
        if ($1 === "mem")
            return [$1, +$2, (+$3).toString(2).padStart(36, "0").split("")];
    })
    .reduce(
        (
            $ = {
                state: { mask: [], mem: [] },
                mask(value) {
                    this.state.mask = value;
                    return this;
                },
                mem(address, value) {
                    this.state.mem[address] = value
                        .zip(this.state.mask)
                        .map(([value, mask]) => (mask === "X" ? value : mask))
                        .pipe(s => parseInt(s.join(""), 2));
                    return this;
                },
            },
            [op, ...args]
        ) => $[op](...args),
        undefined
    )
    .pipe(({ state: { mem } }) => mem.sum())
    .pipe(console.log);
