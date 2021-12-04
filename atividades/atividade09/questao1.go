package main

import (
	"fmt"
	"math"
)

var π float64 = 22.0 / 7.0
var raio_circ float64 = 1.984

func main() {
	var area_circ = π * math.Pow(raio_circ, 2)
	var comp_circ = 2 * π * raio_circ

	fmt.Printf("Área da Circunferência: %f\n", area_circ)
	fmt.Printf("Comprimento da Circunferência: %f\n", comp_circ)
}
