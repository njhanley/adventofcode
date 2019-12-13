package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type vec3 [3]int

type moon struct {
	pos vec3
	vel vec3
}

func parse(r io.Reader) (moons []moon, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var x, y, z int
		_, err := fmt.Sscanf(scanner.Text(), "<x=%d, y=%d, z=%d>", &x, &y, &z)
		if err != nil {
			return nil, err
		}
		moons = append(moons, moon{pos: vec3{x, y, z}})
	}
	if scanner.Err() != nil {
		return nil, err
	}
	return moons, nil
}

func pairs(n int) [][2]int {
	if n == 0 {
		return nil
	}
	n--
	a := make([][2]int, n)
	for i := range a {
		a[i] = [2]int{n, i}
	}
	return append(a, pairs(n)...)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	moons, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	_pairs := pairs(len(moons))
	for i := 0; i < 1000; i++ {
		for _, pair := range _pairs {
			for j := range (vec3{}) {
				if moons[pair[0]].pos[j] < moons[pair[1]].pos[j] {
					moons[pair[0]].vel[j]++
					moons[pair[1]].vel[j]--
				} else if moons[pair[0]].pos[j] > moons[pair[1]].pos[j] {
					moons[pair[0]].vel[j]--
					moons[pair[1]].vel[j]++
				}
			}
		}
		for j := range moons {
			for k, v := range moons[j].vel {
				moons[j].pos[k] += v
			}
		}
	}
	var energy int
	for _, m := range moons {
		var potential, kinetic int
		for _, p := range m.pos {
			potential += abs(p)
		}
		for _, v := range m.vel {
			kinetic += abs(v)
		}
		energy += potential * kinetic
	}
	fmt.Println(energy)
}
