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

const (
	_ = iota
	north
	south
	west
	east
)

type point struct{ x, y int }

func main() {
	code, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}

	ic := newIntcode(code)
	go ic.run()

	accessible := make(map[point]bool)
	var oxygen point
	var direction int64 = north
	var pos point
	for {
		var p point
		for originalDirection := direction; ; {
			p = pos
			switch direction {
			case north:
				p.y++
			case south:
				p.y--
			case west:
				p.x--
			case east:
				p.x++
			}
			if _, visited := accessible[p]; !visited {
				break
			}
			if direction == originalDirection {
				break
			}
			direction = []int64{
				north: west,
				south: east,
				west:  south,
				east:  north,
			}[direction]
		}

		ic.input <- direction
		status := <-ic.output

		fmt.Println(p, status)
		if status == 0 {
			accessible[p] = false
			direction = []int64{
				north: west,
				south: east,
				west:  south,
				east:  north,
			}[direction]
		} else {
			pos = p
			accessible[p] = true
			if status == 2 {
				oxygen = p
				break
			}
		}
	}

	var min, max point
	for p := range accessible {
		if p.x < min.x {
			min.x = p.x
		} else if p.x > max.x {
			max.x = p.x
		}
		if p.y < min.y {
			min.y = p.y
		} else if p.y > max.y {
			max.y = p.y
		}
	}
	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			if p := (point{x, y}); p == oxygen {
				fmt.Print("O")
			} else if accessible[p] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
