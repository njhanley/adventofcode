"use strict";

const fs = require("fs");

Array.prototype.filterMap = function (callback) {
    return this.reduce((newArray, currentValue, index, array) => {
        const value = callback(currentValue, index, array);
        if (value !== undefined) newArray.push(value);
        return newArray;
    }, []);
};

Object.prototype.chain = function (callback) {
    return callback(this);
};

Object.prototype.log = function () {
    console.log(this, ...arguments);
};

function parse(str, regexp, callback) {
    return Array.from(str.matchAll(regexp), callback);
}

fs.readFileSync("input.txt", "utf-8")
    .chain(input =>
        parse(input, /(\w+) ([+-]\d+)/g, match => [match[1], +match[2]])
    )
    .chain(corruptCode => {
        const device = {
            execute([op, arg]) {
                switch (op) {
                    case "acc":
                        this.a += arg;
                        break;
                    case "jmp":
                        this.pc += arg - 1;
                        break;
                    case "nop":
                        break;
                }
                this.pc++;
            },
            run(code) {
                this.a = this.pc = 0;
                while (this.pc !== code.length) {
                    const [instruction] = code.splice(this.pc, 1, null);
                    if (!instruction) return false;
                    this.execute(instruction);
                }
                return true;
            },
        };
        return corruptCode
            .filterMap(([op, arg], index) => {
                const swap = { jmp: "nop", nop: "jmp" };
                return swap[op]?.chain(op => [index, 1, [op, arg]]);
            })
            .chain(patches => {
                patches.find(patch => {
                    const code = [...corruptCode];
                    code.splice(...patch);
                    return device.run(code);
                });
                return device.a;
            });
    })
    .log();
