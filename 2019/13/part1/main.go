package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func parse(r io.Reader) (code []int64, err error) {
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
		value, err := strconv.ParseInt(string(token), 10, 64)
		if err != nil {
			return nil, err
		}
		code = append(code, value)
	}
	return code, nil
}

type intcode struct {
	mem    []int64
	ip, rb int64
	input  chan int64
	output chan int64
}

func newIntcode(code []int64) *intcode {
	return &intcode{
		mem:    append([]int64(nil), code...),
		input:  make(chan int64),
		output: make(chan int64),
	}
}

func (ic *intcode) grow(n int64) {
	mem := make([]int64, int64(len(ic.mem))+n) // len(mem) >= 2*len(ic.mem)
	copy(mem, ic.mem)
	ic.mem = mem
}

func (ic *intcode) get(i int64) int64 {
	if i >= int64(len(ic.mem)) {
		ic.grow(i)
	}
	return ic.mem[i]
}

func (ic *intcode) getref(i int64) *int64 {
	if i >= int64(len(ic.mem)) {
		ic.grow(i)
	}
	return &ic.mem[i]
}

func (ic *intcode) op() int64 {
	return ic.get(ic.ip) % 100
}

func (ic *intcode) mode(n int64) int64 {
	pos := [...]int64{0, 100, 1000, 10000}
	return ic.get(ic.ip) / pos[n] % 10
}

func (ic *intcode) param(n int64) int64 {
	switch ic.mode(n) {
	case 0:
		return ic.get(ic.get(ic.ip + n))
	case 1:
		return ic.get(ic.ip + n)
	case 2:
		return ic.get(ic.rb + ic.get(ic.ip+n))
	default:
		panic("invalid mode")
	}
}

func (ic *intcode) paramref(n int64) *int64 {
	switch ic.mode(n) {
	case 0:
		return ic.getref(ic.get(ic.ip + n))
	case 2:
		return ic.getref(ic.rb + ic.get(ic.ip+n))
	default:
		panic("invalid mode")
	}
}

func (ic *intcode) run() {
	for {
		switch ic.op() {
		case 1:
			*ic.paramref(3) = ic.param(1) + ic.param(2)
			ic.ip += 4
		case 2:
			*ic.paramref(3) = ic.param(1) * ic.param(2)
			ic.ip += 4
		case 3:
			*ic.paramref(1) = <-ic.input
			ic.ip += 2
		case 4:
			ic.output <- ic.param(1)
			ic.ip += 2
		case 5:
			if ic.param(1) != 0 {
				ic.ip = ic.param(2)
			} else {
				ic.ip += 3
			}
		case 6:
			if ic.param(1) == 0 {
				ic.ip = ic.param(2)
			} else {
				ic.ip += 3
			}
		case 7:
			if ic.param(1) < ic.param(2) {
				*ic.paramref(3) = 1
			} else {
				*ic.paramref(3) = 0
			}
			ic.ip += 4
		case 8:
			if ic.param(1) == ic.param(2) {
				*ic.paramref(3) = 1
			} else {
				*ic.paramref(3) = 0
			}
			ic.ip += 4
		case 9:
			ic.rb += ic.param(1)
			ic.ip += 2
		case 99:
			close(ic.output)
			return
		}
	}
}

type point struct{ x, y int64 }

func main() {
	code, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	ic := newIntcode(code)
	go ic.run()
	screen := make(map[point]int64)
	for {
		x, ok := <-ic.output
		if !ok {
			break
		}
		y := <-ic.output
		id := <-ic.output
		screen[point{x, y}] = id
	}
	var blocks int
	for _, id := range screen {
		if id == 2 {
			blocks++
		}
	}
	fmt.Println(blocks)
}
