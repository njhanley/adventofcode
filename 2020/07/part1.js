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

Set.prototype.union = function (set) {
    return new Set([...this, ...set]);
};

String.prototype.lines = function () {
    return parse(this, /(.*)\n|(.+)$/g, match => match[1] ?? match[2]);
};

function parse(str, regexp, callback) {
    return Array.from(str.matchAll(regexp), callback);
}

fs.readFileSync("input.txt", "utf-8")
    .lines()
    .map(line => {
        const [color] = line.match(/\w+ \w+/);
        const contents = new Map(
            parse(line, /(\d) (\w+ \w+)/g, match => [match[2], +match[1]])
        );
        return { color, contents };
    })
    .chain(rules => {
        function contains(color) {
            return rules
                .filter(rule => rule.contents.has(color))
                .reduce(
                    (bags, rule) =>
                        bags.add(rule.color).union(contains(rule.color)),
                    new Set()
                );
        }
        return contains("shiny gold");
    })
    .size.log();
