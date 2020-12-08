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

Object.prototype.chain = function (callback) {
    return callback(this);
};

String.prototype.lines = function () {
    return parse(this, /(.*)\n|(.+)$/g, match => match[1] ?? match[2]);
};

function parse(str, regexp, callback) {
    return Array.from(str.matchAll(regexp), callback);
}

fs.readFileSync("input.txt", "utf-8")
    .lines()
    .reduce((rules, line) => {
        const [color] = line.match(/\w+ \w+/);
        const contents = parse(line, /(\d) (\w+ \w+)/g, match => [
            match[2],
            +match[1],
        ]);
        return rules.set(color, contents);
    }, new Map())
    .chain(rules => {
        function contains(color) {
            return rules
                .get(color)
                .chain(
                    contents =>
                        1 +
                        contents.sum(
                            ([color, quantity]) => quantity * contains(color)
                        )
                );
        }
        return contains("shiny gold") - 1;
    })
    .log();
