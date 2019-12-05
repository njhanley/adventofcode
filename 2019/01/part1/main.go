package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fuelRequired(mass int) (fuel int) {
	return mass/3 - 2
}

func main() {
	var fuel int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		fuel += fuelRequired(mass)
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Println(fuel)
}
