const fs = require("fs");

function parse(str, regexp, callback) {
    return Array.from(str.matchAll(regexp), callback);
}

function sum(array, callback) {
    return array.reduce((total, value) => total + callback(value), 0);
}

const input = fs.readFileSync("input.txt", "utf-8");
const passports = parse(input, /(\S+\s)+(\n|$)/g, match =>
    Object.fromEntries(
        parse(match[0], /(\w+):(\S+)/g, match => [match[1], match[2]])
    )
);

const required = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"];

const valid = sum(passports, passport =>
    required.every(field => field in passport)
);
console.log(valid);
