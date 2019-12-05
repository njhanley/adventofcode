package main

import (
	"fmt"
	"os"
	"strconv"
)

func repeated(s string) int {
	var i int
	for c := s[0]; i < len(s) && s[i] == c; i++ {
	}
	return i
}

func doubled(s string) bool {
	for len(s) > 0 {
		n := repeated(s)
		if n == 2 {
			return true
		}
		s = s[n:]
	}
	return false
}

func nondecreasing(s string) bool {
	c := s[0]
	for i := 1; i < 6; i++ {
		if s[i] < c {
			return false
		}
		c = s[i]
	}
	return true
}

func valid(s string) bool {
	return len(s) == 6 && doubled(s) && nondecreasing(s)
}

func main() {
	var input [2]int
	for i := range input {
		var err error
		input[i], err = strconv.Atoi(os.Args[i+1])
		if err != nil {
			panic(err)
		}
	}
	var n int
	for password := input[0]; password <= input[1]; password++ {
		if valid(strconv.Itoa(password)) {
			n++
		}
	}
	fmt.Println(n)
}
