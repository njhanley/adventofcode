const fs = require("fs");

const input = fs
    .readFileSync("input.txt", "utf-8")
    .split("\n")
    .map(line => parseInt(line, 10));

for (let i = 0; i < input.length; i++)
    for (let j = i + 1; j < input.length; j++)
        for (let k = j + 1; k < input.length; k++)
            if (input[i] + input[j] + input[k] == 2020)
                console.log(input[i] * input[j] * input[k]);
