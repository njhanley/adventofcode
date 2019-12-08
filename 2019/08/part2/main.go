package main

import (
	"io/ioutil"
	"os"
)

func combineLayers(layers [][]byte) []byte {
	image, layers := layers[len(layers)-1], layers[:len(layers)-1]
	for i := len(layers) - 1; i >= 0; i-- {
		for i, c := range layers[i] {
			if c != '2' {
				image[i] = c
			}
		}
	}
	return image
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	data = data[:len(data)-1]
	var layers [][]byte
	const width, height = 25, 6
	for i := 0; i < len(data); i += width * height {
		layers = append(layers, data[i:i+width*height])
	}
	image := combineLayers(layers)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if image[y*width+x] == '0' {
				os.Stdout.Write([]byte{' '})
			} else {
				os.Stdout.Write([]byte("█"))
			}
		}
		os.Stdout.Write([]byte{'\n'})
	}
}
