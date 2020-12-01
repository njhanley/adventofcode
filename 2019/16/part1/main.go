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
	signal, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	pattern := []int{0, 1, 0, -1}
	_signal := make([]int, len(signal))
	for i := 0; i < 100; i++ {
		for j := range signal {
			for k := range signal {
				_signal[j] += signal[k] * pattern[(k+1)/(j+1)%len(pattern)]
			}
			_signal[j] = abs(_signal[j]) % 10
		}
		for j := range signal {
			signal[j] = _signal[j]
			_signal[j] = 0
		}
	}
	for i := range signal[:8] {
		fmt.Print(signal[i])
	}
	fmt.Println()
}
