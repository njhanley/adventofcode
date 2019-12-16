package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

type material struct {
	quantity int
	chemical string
}

type reaction struct {
	inputs []material
	output material
}

func parse(r io.Reader) (map[string]reaction, error) {
	recipes := make(map[string]reaction)
	scanner := bufio.NewScanner(r)
	re := regexp.MustCompile(`([0-9]+) ([A-Z]+)`)
	for scanner.Scan() {
		var a []material
		for _, match := range re.FindAllStringSubmatch(scanner.Text(), -1) {
			quantity, err := strconv.Atoi(match[1])
			if err != nil {
				return nil, err
			}
			a = append(a, material{quantity, match[2]})
		}
		r := reaction{a[:len(a)-1], a[len(a)-1]}
		recipes[r.output.chemical] = r
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return recipes, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countOre(recipes map[string]reaction, materials map[string]int, target material) (ore int) {
	if target.chemical == "ORE" {
		return target.quantity
	}
	if target.quantity == 0 {
		return 0
	}

	recipe := recipes[target.chemical]
	multiple := target.quantity / recipe.output.quantity
	if target.quantity%recipe.output.quantity != 0 {
		multiple++
	}

	for _, input := range recipe.inputs {
		input.quantity *= multiple
		take := min(input.quantity, materials[input.chemical])
		input.quantity -= take
		materials[input.chemical] -= take
		ore += countOre(recipes, materials, input)
	}

	materials[recipe.output.chemical] += recipe.output.quantity*multiple - target.quantity

	return ore
}

func main() {
	recipes, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	var fuel int
	for countOre(recipes, make(map[string]int), material{fuel + 1, "FUEL"}) <= 1000000000000 {
		fuel++
	}
	fmt.Println(fuel)
}
