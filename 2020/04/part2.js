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

const isYear = s => /^\d{4}$/.test(s);
const between = (n, x, y) => x <= n && n <= y;

const rules = {
    byr: s => isYear(s) && between(+s, 1920, 2002),
    iyr: s => isYear(s) && between(+s, 2010, 2020),
    eyr: s => isYear(s) && between(+s, 2020, 2030),
    hgt: s => {
        const [match, height, unit] = /^(\d+)(cm|in)$/.exec(s) || [];
        return match && unit === "cm"
            ? between(+height, 150, 193)
            : unit === "in"
            ? between(+height, 59, 76)
            : false;
    },
    hcl: s => /^#[0-9a-f]{6}$/.test(s),
    ecl: s => ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"].includes(s),
    pid: s => /^\d{9}$/.test(s),
};

const valid = sum(passports, passport =>
    Object.entries(rules).every(
        ([field, validate]) => field in passport && validate(passport[field])
    )
);
console.log(valid);
