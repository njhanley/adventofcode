const fs = require("fs");

Array.prototype.sum = function (callback) {
    return this.reduce(
        (accumulator, element, index) => accumulator + callback(element, index),
        0
    );
};

Set.prototype.intersect = function (set) {
    return new Set([...this].filter(element => set.has(element)));
};

function parse(str, regexp, callback) {
    return Array.from(str.matchAll(regexp), callback);
}

const input = fs.readFileSync("input.txt", "utf-8");
const groups = input
    .split("\n\n")
    .map(record =>
        parse(record, /([a-z]+)(\n|$)/g, match => new Set(match[1]))
    );
console.log(groups.sum(group => group.reduce((a, b) => a.intersect(b)).size));
