package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type node struct {
	parent   string
	children []string
}

func parse(r io.Reader) (map[string]*node, error) {
	tree := make(map[string]*node)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		token := scanner.Text()
		parent, child := token[:3], token[4:]
		if _node, exists := tree[child]; exists {
			_node.parent = parent
		} else {
			tree[child] = &node{parent: parent}
		}
		if _node, exists := tree[parent]; exists {
			_node.children = append(_node.children, child)
		} else {
			tree[parent] = &node{children: []string{child}}
		}
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return tree, nil
}

func findPath(tree map[string]*node, from, to string, path []string) []string {
	path = append(path, from)
	if from == to {
		return path
	}
	node := tree[from]
	neighbors := append([]string(nil), node.children...)
	if node.parent != "" {
		neighbors = append(neighbors, node.parent)
	}
	for _, neighbor := range neighbors {
		if n := len(path); n >= 2 && neighbor == path[n-2] {
			continue
		}
		path := findPath(tree, neighbor, to, append([]string(nil), path...))
		if path != nil {
			return path
		}
	}
	return nil
}

func main() {
	tree, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(findPath(tree, tree["YOU"].parent, tree["SAN"].parent, nil)) - 1)
}
