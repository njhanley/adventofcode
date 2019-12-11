package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/cmplx"
	"os"
	"sort"
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

func findNeighborhood(asteroids map[complex128]bool, location complex128) map[float64][]complex128 {
	neighborhood := make(map[float64][]complex128)
	for asteroid := range asteroids {
		if asteroid == location {
			continue
		}
		direction := cmplx.Phase(asteroid - location)
		neighborhood[direction] = append(neighborhood[direction], asteroid)
	}
	for _, ray := range neighborhood {
		sort.Slice(ray, func(i, j int) bool { return cmplx.Abs(ray[i]-location) < cmplx.Abs(ray[j]-location) })
	}
	return neighborhood
}

func main() {
	asteroids, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	var highestDetection int
	var bestLocation complex128
	for location := range asteroids {
		neighborhood := findNeighborhood(asteroids, location)
		if n := len(neighborhood); n > highestDetection {
			highestDetection = n
			bestLocation = location
		}
	}
	neighborhood := findNeighborhood(asteroids, bestLocation)
	directions := make([]float64, 0, len(neighborhood))
	for direction := range neighborhood {
		directions = append(directions, direction)
	}
	sort.Float64s(directions)
	var a, b []float64
	for _, direction := range directions {
		if direction < -math.Pi/2 {
			a = append(a, direction)
		} else {
			b = append(b, direction)
		}
	}
	copy(directions[:len(b)], b)
	copy(directions[len(b):], a)
	var destroyed []complex128
	for len(neighborhood) > 0 {
		for _, direction := range directions {
			if len(neighborhood[direction]) == 0 {
				delete(neighborhood, direction)
				continue
			}
			destroyed = append(destroyed, neighborhood[direction][0])
			neighborhood[direction] = neighborhood[direction][1:]
		}
	}
	fmt.Println(int(real(destroyed[199]))*100 + int(imag(destroyed[199])))
}
