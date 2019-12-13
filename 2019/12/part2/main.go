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

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	moons, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	_pairs := pairs(len(moons))
	// hardcoded for 4 bodies
	cycles := [4]int{1, 0, 0, 0}
	offsets := [4]int{0, 0, 0, 0}
	for i := range (vec3{}) {
		occurred := make(map[[4][2]int]int)
		step := offsets[i]
		for {
			var state [4][2]int
			for j := range moons {
				state[j][0] = moons[j].pos[i]
				state[j][1] = moons[j].vel[i]
			}
			if prev, exists := occurred[state]; exists {
				cycles[i+1] = step - offsets[i]
				offsets[i+1] = prev
				break
			}
			occurred[state] = step

			for _, pair := range _pairs {
				if moons[pair[0]].pos[i] < moons[pair[1]].pos[i] {
					moons[pair[0]].vel[i]++
					moons[pair[1]].vel[i]--
				} else if moons[pair[0]].pos[i] > moons[pair[1]].pos[i] {
					moons[pair[0]].vel[i]--
					moons[pair[1]].vel[i]++
				}
			}
			for j := range moons {
				moons[j].pos[i] += moons[j].vel[i]
			}
			step++
		}
	}
	fmt.Println(lcm(cycles[1], lcm(cycles[2], cycles[3])))
}
