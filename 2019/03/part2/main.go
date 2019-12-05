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
	wires [2]int // wire length
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func parse(r io.Reader) (grid map[point]wires, err error) {
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
	grid = make(map[point]wires)
	for i := range tokens {
		var p point
		var wireLength int
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
					wireLength++
					w := grid[p]
					if w[i] == 0 {
						w[i] = wireLength
					} else {
						w[i] = min(w[i], wireLength)
					}
					grid[p] = w
					if p.y >= q.y {
						break
					}
				}
			case 'D':
				q := point{p.x, p.y - length}
				for {
					p.y--
					wireLength++
					w := grid[p]
					if w[i] == 0 {
						w[i] = wireLength
					} else {
						w[i] = min(w[i], wireLength)
					}
					grid[p] = w
					if p.y <= q.y {
						break
					}
				}
			case 'R':
				q := point{p.x + length, p.y}
				for {
					p.x++
					wireLength++
					w := grid[p]
					if w[i] == 0 {
						w[i] = wireLength
					} else {
						w[i] = min(w[i], wireLength)
					}
					grid[p] = w
					if p.x >= q.x {
						break
					}
				}
			case 'L':
				q := point{p.x - length, p.y}
				for {
					p.x--
					wireLength++
					w := grid[p]
					if w[i] == 0 {
						w[i] = wireLength
					} else {
						w[i] = min(w[i], wireLength)
					}
					grid[p] = w
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

func main() {
	grid, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	type cross struct {
		p point
		w wires
	}
	var intersections []cross
	for p, w := range grid {
		if w[0] != 0 && w[1] != 0 {
			intersections = append(intersections, cross{p, w})
		}
	}
	sort.Slice(intersections, func(i, j int) bool {
		wi, wj := intersections[i].w, intersections[j].w
		return wi[0]+wi[1] < wj[0]+wj[1]
	})
	minimal := intersections[0]
	fmt.Println(minimal, minimal.w[0]+minimal.w[1])
}
