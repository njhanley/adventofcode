"use strict";

const fs = require("fs");

Array.prototype.filterMap = function (callback) {
    return this.reduce((newArray, currentValue, index, array) => {
        const value = callback(currentValue, index, array);
        if (value !== undefined) newArray.push(value);
        return newArray;
    }, []);
};

Array.prototype.sum = function (callback) {
    return this.reduce(
        (accumulator, element, index, array) =>
            accumulator + (callback?.(element, index, array) ?? element),
        0
    );
};

Array.prototype.zip = function (...rest) {
    return this.map((element, index) => [
        element,
        ...rest.map(array => array[index]),
    ]);
};

Object.prototype.pipe = function (callback) {
    return callback(this);
};

Set.prototype.difference = function (set) {
    return new Set([...this].filter(element => !set.has(element)));
};

Set.prototype.union = function (set) {
    return new Set([...this, ...set]);
};

String.prototype.matchMap = function (regexp, callback) {
    return regexp.global
        ? Array.from(this.matchAll(regexp), callback)
        : callback(this.match(regexp));
};

function range(start, stop, step = 1) {
    return Array.from(
        { length: (stop - start) / step + 1 },
        (_, index) => start + index * step
    );
}

fs.readFileSync("input.txt", "utf-8")
    .pipe(input => ({
        rules: input.matchMap(/([\w ]+): (\d+)-(\d+) or (\d+)-(\d+)/g, $ => [
            $[1], new Set(range(+$[2], +$[3]).concat(range(+$[4], +$[5]))),
        ]),
        your: input.matchMap(/your ticket:\n([\d,]+)/, $ =>
            $[1].split(",").map(x => +x)
        ),
        nearby: input.matchMap(/nearby tickets:\n([\d,\n]+)/, $ =>
            $[1].matchMap(/([\d,]+)/g, $ => $[1].split(",").map(x => +x))
        ),
    }))
    .pipe(({ rules, your, nearby }) => {
        const valid = rules.map(([, valid]) => valid).reduce((a, b) => a.union(b));
        return {
            rules,
            your,
            nearby: nearby.filter(ticket => ticket.every(value => valid.has(value))),
        };
    })
    .pipe(({rules, your, nearby}) => {
        const transpose = a => a[0].map((_, i) => a.map(x => x[i]));
        const a = [your, ...nearby].pipe(transpose);
        const possible = rules.map(([name, valid]) => [name, new Set([...a.entries()].filterMap(([index, field]) => { if (field.every(value => valid.has(value))) return index; }))]);
        const fields = possible.map(([name, indexes]) => [name, [indexes, ...possible.filterMap(([_name, _indexes]) => { if (name !== _name) return _indexes; })].reduce((a, b) => a.difference(b))]);
        return fields;
    })
    .pipe(console.log);
