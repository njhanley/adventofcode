package main

import (
	"fmt"
	"io/ioutil"
	"math/bits"
	"os"
)

func countBytes(b []byte, c byte) int {
	var n int
	for i := range b {
		if b[i] == c {
			n++
		}
	}
	return n
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	data = data[:len(data)-1]
	var layers [][]byte
	const stride = 25 * 6
	for i := 0; i < len(data); i += stride {
		layers = append(layers, data[i:i+stride])
	}
	fewestZeros := int(1<<(bits.UintSize-1) - 1)
	var fewestZerosLayer []byte
	for _, layer := range layers {
		zeros := countBytes(layer, '0')
		if zeros < fewestZeros {
			fewestZeros = zeros
			fewestZerosLayer = layer
		}
	}
	fmt.Println(countBytes(fewestZerosLayer, '1') * countBytes(fewestZerosLayer, '2'))
}
