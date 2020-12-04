const fs = require("fs");

const input = fs.readFileSync("input.txt", "utf-8");
const map = input.split("\n");

function check(x, y) {
    const row = map[y];
    return row[x % row.length];
}

const d = { x: 3, y: 1 };

let trees = 0;
for (let x = 0, y = 0; y < map.length; x += d.x, y += d.y)
    if (check(x, y) === "#") trees++;
console.log(trees);
