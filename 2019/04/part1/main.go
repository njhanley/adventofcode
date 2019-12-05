package main

import (
	"fmt"
	"os"
	"strconv"
)

func valid(password int) bool {
	s := strconv.Itoa(password)
	if len(s) != 6 {
		return false
	}
	repeated := false
	c := s[0]
	for i := 0; i < 5; i++ {
		if s[i] == s[i+1] {
			repeated = true
		}
		if s[i+1] < c {
			return false
		}
		c = s[i+1]
	}
	if !repeated {
		return false
	}
	return true
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
		if valid(password) {
			n++
		}
	}
	fmt.Println(n)
}
