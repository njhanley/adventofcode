const fs = require("fs");

const input = fs.readFileSync("input.txt", "utf-8");
const map = input.split("\n");

function check(x, y) {
    const row = map[y];
    return row[x % row.length];
}

const slopes = [
    { x: 1, y: 1 },
    { x: 3, y: 1 },
    { x: 5, y: 1 },
    { x: 7, y: 1 },
    { x: 1, y: 2 },
];

let product = 1;
for (const d of slopes) {
    let trees = 0;
    for (let x = 0, y = 0; y < map.length; x += d.x, y += d.y)
        if (check(x, y) === "#") trees++;
    product *= trees;
}
console.log(product);
