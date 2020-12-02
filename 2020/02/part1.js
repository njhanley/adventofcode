const fs = require("fs");

const input = fs.readFileSync("input.txt", "utf-8");
const entries = Array.from(
    input.matchAll(/(\d+)-(\d+) ([a-z]): ([a-z]+)/g),
    match => ({
        min: +match[1],
        max: +match[2],
        letter: match[3],
        password: match[4],
    })
);

function sum(a, f) {
    return a.reduce((n, c) => n + f(c), 0);
}

const valid = sum(entries, entry => {
    const count = sum(Array.from(entry.password), c => c === entry.letter);
    return entry.min <= count && count <= entry.max;
});
console.log(valid);
