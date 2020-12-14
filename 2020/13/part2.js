"use strict";

const fs = require("fs");

Object.prototype.pipe = function (callback) {
    return callback(this);
};

String.prototype.matchMap = function (regexp, callback) {
    return regexp.global
        ? Array.from(this.matchAll(regexp), callback)
        : callback(this.match(regexp));
};

fs.readFileSync("input.txt", "utf-8")
    .matchMap(/(\n|,)([x\d]+)/g, match => +match[2])
    .reduce((a, n, i) => (n ? [...a, { cycle: n, offset: i }] : a), [])
    .reduce((a, b) => {
        let offset = a.offset;
        while ((offset + b.offset) % b.cycle) offset += a.cycle;
        return { cycle: a.cycle * b.cycle, offset };
    })
    .pipe(({ offset }) => offset)
    .pipe(console.log);
