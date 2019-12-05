package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

type intcode []int

func parse(r io.Reader) (code intcode, err error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, ','); i >= 0 {
			return i + 1, data[:i], nil
		}
		if atEOF {
			return len(data), bytes.TrimSuffix(data, []byte{'\n'}), nil
		}
		return 0, nil, nil
	})
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		code = append(code, n)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return code, err
}

const (
	opADD = 1
	opMUL = 2
	opHLT = 99
)

func (code intcode) run() error {
	for i := 0; i < len(code); i += 4 {
		switch op := code[i]; op {
		case opADD:
			code[code[i+3]] = code[code[i+1]] + code[code[i+2]]
		case opMUL:
			code[code[i+3]] = code[code[i+1]] * code[code[i+2]]
		case opHLT:
			return nil
		default:
			return fmt.Errorf("unknown opcode %d at %d", op, i)
		}
	}
	return nil
}

func main() {
	code, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}

	_code := make(intcode, len(code))
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			copy(_code, code)
			_code[1] = i
			_code[2] = j
			_code.run()
			if _code[0] == 19690720 {
				fmt.Println(i, j)
				return
			}
		}
	}
}
