package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

type intcode struct {
	mem    []int
	ip     int
	input  chan int
	output chan int
}

func parse(r io.Reader) (code *intcode, err error) {
	var memory []int
	reader := bufio.NewReader(r)
	for eof := false; !eof; {
		token, err := reader.ReadBytes(',')
		if err != nil {
			if err == io.EOF {
				eof = true
			} else {
				return nil, err
			}
		}
		token = bytes.TrimRight(token, ",\n")
		value, err := strconv.Atoi(string(token))
		if err != nil {
			return nil, err
		}
		memory = append(memory, value)
	}
	return &intcode{
		mem:    memory,
		input:  make(chan int),
		output: make(chan int),
	}, nil
}

func (code *intcode) op() int {
	return code.mem[code.ip] % 100
}

func (code *intcode) param(n int) int {
	pos := [...]int{0, 100, 1000, 10000}
	mode := code.mem[code.ip] / pos[n] % 10
	if mode == 0 {
		return code.mem[code.mem[code.ip+n]]
	}
	return code.mem[code.ip+n]
}

func (code *intcode) paramref(n int) *int {
	return &code.mem[code.mem[code.ip+n]]
}

func (code *intcode) run() {
	for {
		switch code.op() {
		case 1:
			*code.paramref(3) = code.param(1) + code.param(2)
			code.ip += 4
		case 2:
			*code.paramref(3) = code.param(1) * code.param(2)
			code.ip += 4
		case 3:
			*code.paramref(1) = <-code.input
			code.ip += 2
		case 4:
			code.output <- code.param(1)
			code.ip += 2
		case 99:
			close(code.input)
			close(code.output)
			return
		}
	}
}

func main() {
	code, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	go code.run()
	code.input <- 1
	for n := range code.output {
		fmt.Println(n)
	}
}
