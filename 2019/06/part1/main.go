package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func parse(r io.Reader) (tree map[string][]string, err error) {
	tree = make(map[string][]string)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		token := scanner.Text()
		parent, child := token[:3], token[4:]
		tree[parent] = append(tree[parent], child)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return tree, nil
}

func countOrbits(tree map[string][]string, node string, depth int) int {
	orbits := depth
	for _, child := range tree[node] {
		orbits += countOrbits(tree, child, depth+1)
	}
	return orbits
}

func main() {
	tree, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(countOrbits(tree, "COM", 0))
}
