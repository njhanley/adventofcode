package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func parse(r io.Reader) (memory []int, err error) {
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
	return memory, nil
}

type intcode struct {
	mem    []int
	ip     int
	input  chan int
	output chan int
}

func newIntcode(memory []int) *intcode {
	return &intcode{
		mem:    append([]int(nil), memory...),
		input:  make(chan int),
		output: make(chan int),
	}
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
		case 5:
			if code.param(1) != 0 {
				code.ip = code.param(2)
			} else {
				code.ip += 3
			}
		case 6:
			if code.param(1) == 0 {
				code.ip = code.param(2)
			} else {
				code.ip += 3
			}
		case 7:
			if code.param(1) < code.param(2) {
				*code.paramref(3) = 1
			} else {
				*code.paramref(3) = 0
			}
			code.ip += 4
		case 8:
			if code.param(1) == code.param(2) {
				*code.paramref(3) = 1
			} else {
				*code.paramref(3) = 0
			}
			code.ip += 4
		case 99:
			close(code.input)
			close(code.output)
			code.ip = -1
			return
		}
	}
}

func (code *intcode) running() bool {
	return code.ip >= 0
}

func permutations(a []int) [][]int {
	var p [][]int
	var fn func(int, []int)
	fn = func(n int, a []int) {
		if n == 1 {
			p = append(p, append([]int(nil), a...))
			return
		}
		for i := 0; i < n-1; i++ {
			fn(n-1, a)
			if n%2 == 0 {
				a[i], a[n-1] = a[n-1], a[i]
			} else {
				a[0], a[n-1] = a[n-1], a[0]
			}
		}
		fn(n-1, a)
	}
	fn(len(a), a)
	return p
}

func main() {
	code, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	var maxSignal int
	for _, settings := range permutations([]int{5, 6, 7, 8, 9}) {
		amps := make([]*intcode, len(settings))
		for i := range amps {
			amps[i] = newIntcode(code)
			go amps[i].run()
			amps[i].input <- settings[i]
		}
		var signal int
		for amps[len(settings)-1].running() {
			for _, amp := range amps {
				amp.input <- signal
				signal = <-amp.output
			}
		}
		if signal > maxSignal {
			maxSignal = signal
		}
	}
	fmt.Println(maxSignal)
}
