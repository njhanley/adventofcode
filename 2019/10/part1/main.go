package main

import (
	"bufio"
	"fmt"
	"io"
	"math/cmplx"
	"os"
)

func parse(r io.Reader) (asteroids map[complex128]bool, err error) {
	asteroids = make(map[complex128]bool)
	scanner := bufio.NewScanner(r)
	for y := 0; scanner.Scan(); y++ {
		for x, c := range scanner.Bytes() {
			if c == '#' {
				asteroids[complex(float64(x), float64(y))] = true
			}
		}
	}
	if scanner.Err() != nil {
		return nil, err
	}
	return asteroids, nil
}

func countAsteroids(asteroids map[complex128]bool, location complex128) int {
	rays := make(map[float64]bool)
	for asteroid := range asteroids {
		rays[cmplx.Phase(asteroid-location)] = true
	}
	return len(rays)
}

func main() {
	asteroids, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	var highestDetection int
	for location := range asteroids {
		if n := countAsteroids(asteroids, location); n > highestDetection {
			highestDetection = n
		}
	}
	fmt.Println(highestDetection)
}
