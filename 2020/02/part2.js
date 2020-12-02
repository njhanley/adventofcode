const fs = require("fs");

const input = fs.readFileSync("input.txt", "utf-8");
const entries = Array.from(
    input.matchAll(/(\d+)-(\d+) ([a-z]): ([a-z]+)/g),
    match => ({
        positions: [match[1] - 1, match[2] - 1],
        letter: match[3],
        password: match[4],
    })
);

function sum(a, f) {
    return a.reduce((n, c) => n + f(c), 0);
}

const valid = sum(
    entries,
    entry => sum(entry.positions, i => entry.password[i] === entry.letter) === 1
);
console.log(valid);
