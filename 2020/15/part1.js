"use strict";

const fs = require("fs");

Object.prototype.pipe = function (callback) {
    return callback(this);
};

fs.readFileSync("input.txt", "utf-8")
    .split(",")
    .map(x => +x)
    .pipe(spoken => {
        const memory = new Map(spoken.map((n, i) => [n, [i]]));
        for (let i = spoken.length; i < 2020; i++) {
            const [$1, $2] = memory.get(spoken[i - 1]);
            const n = $2 === undefined ? 0 : $1 - $2;
            spoken.push(n);
            memory.set(n, [i, memory.get(n)?.shift()]);
        }
        return spoken.pop();
    })
    .pipe(console.log);
