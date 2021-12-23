// Correção: 2,0
package main

import (
	"flag"
	"fmt"
	"math"
)

//Flags
var a = flag.Float64("a", 0.0, "Limite Inferior")
var b = flag.Float64("b", 100.0, "Limite Superior")
var n = flag.Int("n", 1000, "Número de Intervalos")

func main() {
	flag.Parse()

	var h float64
	var sum float64 = 0.0
	var num float64 = float64(*n)

	fmt.Printf("\nEste programa calcula o integral no intervalo [a,b]\n")
	fmt.Printf("Limite Inferior [a]: %g\n", *a)
	fmt.Printf("Limite Superior [b]: %g\n", *b)
	fmt.Printf("Número de Intervalos: %g\n", num)

	*b = (5.0**a)/2.0 - *b
	h = (*b - *a) / num

	for i := 0; i < *n; i++ {
		sum = sum + (f(*a)+f((*a+h)))*h/2
		*a = *a + h
	}

	fmt.Printf("Valor da Integral é: %g\n", sum)
}

func f(x float64) float64 {
	var sin float64 = math.Sin(x)
	return math.Sqrt(1.0 + (sin * sin))
}
