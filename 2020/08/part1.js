"use strict";

const fs = require("fs");

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
    .chain(code => {
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
                while (true) {
                    const [instruction] = code.splice(this.pc, 1, null);
                    if (!instruction) break;
                    this.execute(instruction);
                }
            },
        };
        device.run(code);
        return device.a;
    })
    .log();
