const fs = require("fs");

Array.prototype.sum = function (callback) {
    return this.reduce(
        (accumulator, element, index) => accumulator + callback(element, index),
        0
    );
};

function parse(str, regexp, callback) {
    return Array.from(str.matchAll(regexp), callback);
}

function parseInteger(str, symbols) {
    return Array.prototype.map
        .call(str, c => symbols.indexOf(c))
        .reverse()
        .sum((digit, position) => digit * symbols.length ** position);
}

const input = fs.readFileSync("input.txt", "utf-8");
const passes = parse(input, /([FB]{7})([LR]{3})/g, match => {
    const row = parseInteger(match[1], "FB"),
        column = parseInteger(match[2], "LR"),
        id = row * 8 + column;
    return { row, column, id };
});

const seats = [...Array(2 ** 10).keys()],
    claimed = new Set(passes.map(pass => pass.id));

console.log(
    seats
        .filter(id => !claimed.has(id))
        .find(id => claimed.has(id + 1) && claimed.has(id - 1))
);
