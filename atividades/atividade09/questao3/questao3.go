// Correção: 1,0
package main

import (
	"flag"
	"fmt"
	"math"
)

var π float64 = 22.0 / 7.0
var raio_circ float64

//Flags
var r = flag.Float64("r", 1.0, "raio")
var c = flag.Bool("c", false, "calcular comprimento")
var a = flag.Bool("a", false, "calcular area")

func calcularArea() float64 {
	var area_circ = π * math.Pow(raio_circ, 2)
	return area_circ
}

func calcularComprimento() float64 {
	var comp_circ = 2 * π * raio_circ
	return comp_circ
}

func main() {
	flag.Parse()
	raio_circ = *r

	if *c == true {
		fmt.Printf("Comprimento da Circunferência: %f\n", calcularComprimento())
	}
	if *a == true {
		fmt.Printf("Área da Circunferência: %f\n", calcularArea())
	}
}
