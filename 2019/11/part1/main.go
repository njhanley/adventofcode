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
	io     chan int64
}

func newIntcode(code []int64) *intcode {
	return &intcode{
		mem: append([]int64(nil), code...),
		io:  make(chan int64),
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
			*ic.paramref(1) = <-ic.io
			ic.ip += 2
		case 4:
			ic.io <- ic.param(1)
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
			close(ic.io)
			ic.ip = -1
			return
		}
	}
}

func (ic *intcode) running() bool {
	return ic.ip >= 0
}

const (
	up byte = iota
	right
	down
	left
)

type point struct{ x, y int }

type robot struct {
	ic        *intcode
	position  point
	direction byte
}

func newRobot(code []int64) *robot {
	r := &robot{
		ic: newIntcode(code),
	}
	go r.ic.run()
	return r
}

func (r *robot) paint(image map[point]int64) {
	for r.ic.running() {
		r.ic.io <- image[r.position]
		image[r.position] = <-r.ic.io
		if <-r.ic.io == 0 {
			r.direction = (r.direction - 1) % 4
		} else {
			r.direction = (r.direction + 1) % 4
		}
		switch r.direction {
		case up:
			r.position.y++
		case down:
			r.position.y--
		case right:
			r.position.x++
		case left:
			r.position.x--
		}
	}
}

func main() {
	code, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	image := make(map[point]int64)
	newRobot(code).paint(image)
	fmt.Println(len(image))
}
