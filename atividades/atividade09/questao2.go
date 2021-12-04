package main

import (
	"fmt"
	"math"
)

var π float64 = 22.0 / 7.0
var raio_circ float64 = 1.984

func calcularArea() float64 {
	var area_circ = π * math.Pow(raio_circ, 2)
	return area_circ
}

func calcularComprimento() float64 {
	var comp_circ = 2 * π * raio_circ
	return comp_circ
}

func main() {
	fmt.Printf("Área da Circunferência: %f\n", calcularArea())
	fmt.Printf("Comprimento da Circunferência: %f\n", calcularComprimento())
}
