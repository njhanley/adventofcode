package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type (
	point struct{ x, y int }
	wire  byte
)

func parse(r io.Reader) (grid map[point]wire, err error) {
	reader := bufio.NewReader(r)
	var lines [2][]byte
	for i := range lines {
		lines[i], err = reader.ReadSlice('\n')
		if err != nil {
			return nil, err
		}
		lines[i] = bytes.TrimSuffix(lines[i], []byte{'\n'})
	}
	var tokens [2][][]byte
	for i := range tokens {
		tokens[i] = bytes.Split(lines[i], []byte{','})
	}
	grid = make(map[point]wire)
	for i := range tokens {
		var p point
		for j, segment := range tokens[i] {
			direction := segment[0]
			length, err := strconv.Atoi(string(segment[1:]))
			if err != nil {
				return nil, err
			}
			switch direction {
			case 'U':
				q := point{p.x, p.y + length}
				for {
					p.y++
					grid[p] |= 1 << i
					if p.y >= q.y {
						break
					}
				}
			case 'D':
				q := point{p.x, p.y - length}
				for {
					p.y--
					grid[p] |= 1 << i
					if p.y <= q.y {
						break
					}
				}
			case 'R':
				q := point{p.x + length, p.y}
				for {
					p.x++
					grid[p] |= 1 << i
					if p.x >= q.x {
						break
					}
				}
			case 'L':
				q := point{p.x - length, p.y}
				for {
					p.x--
					grid[p] |= 1 << i
					if p.x <= q.x {
						break
					}
				}
			default:
				return nil, fmt.Errorf("unknown direction %c in token %d", direction, j)
			}
		}
	}
	return grid, nil
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func distance(p, q point) int {
	return abs(p.x-q.x) + abs(p.y-q.y)
}

func main() {
	grid, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	var intersections []point
	for p, w := range grid {
		if w == 0x3 {
			intersections = append(intersections, p)
		}
	}
	sort.Slice(intersections, func(i, j int) bool {
		return distance(point{}, intersections[i]) < distance(point{}, intersections[j])
	})
	fmt.Println(intersections[0], distance(point{}, intersections[0]))
}
