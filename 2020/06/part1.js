"use strict";

const fs = require("fs");

Array.prototype.sum = function (callback) {
    return this.reduce(
        (accumulator, element, index) => accumulator + callback(element, index),
        0
    );
};

Object.prototype.log = function () {
    console.log(this, ...arguments);
};

fs.readFileSync("input.txt", "utf-8")
    .split("\n\n")
    .map(record => new Set(record.replace(/\n/g, "")))
    .sum(group => group.size)
    .log();
