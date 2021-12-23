// Correção: 1,0
package main

import (
	"fmt"
	"medidas/medconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		m := medconv.Metros(v)
		p := medconv.Pés(v)
		q := medconv.Quilos(v)
		l := medconv.Libras(v)

		fmt.Printf("Metros: %s = %s \n", m, medconv.MToP(m))
		fmt.Printf("Pés: %s = %s \n", p, medconv.PToM(p))
		fmt.Printf("Quilos: %s = %s \n", q, medconv.QToL(q))
		fmt.Printf("Libras: %s = %s \n", l, medconv.LToQ(l))
	}
}
