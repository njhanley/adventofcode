const fs = require("fs");

Array.prototype.sum = function (callback) {
    return this.reduce(
        (accumulator, element, index) => accumulator + callback(element, index),
        0
    );
};

const input = fs.readFileSync("input.txt", "utf-8");
const groups = input
    .split("\n\n")
    .map(record => new Set(record.replace(/\n/g, "")));
console.log(groups.sum(group => group.size));
