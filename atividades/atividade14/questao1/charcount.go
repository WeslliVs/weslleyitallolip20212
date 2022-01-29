package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	count_digit := 0
	count_letter := 0
	count_lower := 0
	count_punct := 0
	count_space := 0
	count_symbol := 0
	count_upper := 0

	// Contagem de tamanho das codificações UTF-8 (quantas runas de 1, 2, 3 ou 5 bytes etc)
	var utflen [utf8.UTFMax + 1]int

	// Contagem de caracteres inválidos em UTF-8
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		// Retorna runa, número de bytes que ela ocupa e possível erro
		r, n, err := in.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++

		if (unicode.IsDigit(r)){
			count_digit++
		}

		if (unicode.IsLetter(r)){
			count_letter++
		}

		if (unicode.IsLower(r)){
			count_lower++
		}

		if (unicode.IsPunct(r)){
			count_punct++
		}

		if (unicode.IsSpace(r)){
			count_space++
		}

		if (unicode.IsSymbol(r)){
			count_symbol++
		}

		if (unicode.IsUpper(r)){
			count_upper++
		}

	}

	fmt.Printf("Quantidade de Letras minúsculas: %d\n", count_lower)
	fmt.Printf("Quantidade de Letras maiúsculas: %d\n", count_upper)
	fmt.Printf("Quantidade de Letras: %d\n", count_letter)
	fmt.Printf("Quantidade de Dígitos: %d\n", count_digit)
	fmt.Printf("Quantidade de Pontos: %d\n", count_punct)
	fmt.Printf("Quantidade de Espaços: %d\n", count_space)
	fmt.Printf("Quantidade de Símbolos: %d\n", count_symbol)

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("len\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}