package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func parse(r io.Reader) ([]int, error) {
	reader := bufio.NewReader(r)
	var a []int
	for {
		c, err := reader.ReadByte()
		if err == io.EOF || err == nil && c == '\n' {
			break
		}
		if err != nil {
			return nil, err
		}
		a = append(a, int(c-'0'))
	}
	return a, nil
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	input, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	signal := input
	offset := 0
	//signal := make([]int, len(input)*10000)
	//for i := range signal {
	//	signal[i] = input[i%len(input)]
	//}
	//var offset int
	//for i := range signal[:7] {
	//	offset = offset*10 + signal[i]
	//}
	_signal := make([]int, len(signal))
	for i := 0; i < 100; i++ {
		//for j := len(signal) - 1; j >= offset; j-- {
		for j := range signal {
			var negate bool
			for k := j; k < len(signal); k += 2 * (j + 1) {
				fmt.Println(i, j, k)
				for l := 0; l <= j && k+l < len(signal); l++ {
					if negate {
						_signal[j] -= signal[k+l]
					} else {
						_signal[j] += signal[k+l]
					}
				}
				negate = !negate
			}
			_signal[j] = abs(_signal[j]) % 10
		}
		for j := range signal {
			signal[j] = _signal[j]
			_signal[j] = 0
		}
	}
	for _, n := range signal[offset : offset+8] {
		fmt.Print(n)
	}
	fmt.Println()
}
