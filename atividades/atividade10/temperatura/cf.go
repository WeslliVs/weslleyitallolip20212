package main

import (
	"fmt"
	"os"
	"strconv"
	"temperatura/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		c := tempconv.Celsius(t)
		f := tempconv.Fahrenheit(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("Celsius: %s = %s, %s = %s \n", c, tempconv.CToF(c), c, tempconv.CToK(c))
		fmt.Printf("Fahrenheit: %s = %s, %s = %s \n", f, tempconv.FToC(f), f, tempconv.FToK(f))
		fmt.Printf("Kelvin: %s = %s, %s = %s \n", k, tempconv.KToC(k), k, tempconv.KToF(k))
	}
}
