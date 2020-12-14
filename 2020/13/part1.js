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
    .matchMap(/(\d+)\n(.+)/, match => {
        const timestamp = +match[1];
        const ids = match[2].matchMap(/\d+/g, match => +match[0]);
        return [timestamp, ids];
    })
    .pipe(([timestamp, ids]) =>
        ids.map(id => ({
            id: id,
            wait: id * (1 + Math.floor(timestamp / id)) - timestamp,
        }))
    )
    .reduce((a, b) => (a.wait < b.wait ? a : b))
    .pipe(bus => bus.id * bus.wait)
    .pipe(console.log);
