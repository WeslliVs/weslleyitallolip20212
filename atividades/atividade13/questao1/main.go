package main

import "fmt"

var array = [5]int {1, 2, 3, 4, 5}
var reverso [5]int

func reverse(s []int) []int {
	var aux = [5]int {1, 2, 3, 4, 5}

	for i, j:= 0, 4; i<j; i, j = i + 1, j - 1 {
		s[i] = s[j]
		s[j] = aux[i]
	}
	return s
}

func main() {
	fmt.Println("Array Original: ", array)
	reverse(array[:])
	fmt.Println("Array Reverso: ", array)
}