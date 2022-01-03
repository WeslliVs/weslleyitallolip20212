package main

import "fmt"

func expr() {
	term()
	for nextToken == ADD_OP || nextToken == SUB_OP {
		lex()
		term()
	}
	return expr
}

func term() {
	fmt.Printf("\nEnter <term>: ")
	factor()

	for nextToken == MULT_OP || nextToken == DIV_OP {
		lex()
		factor()
	}
	fmt.Printf("\nExit <term>\n")
	return term()
}

func factor() {
	if nextToken == ID_CODE || nextToken == INT_CODE {
		lex()
	} else if nexToken == LP_CODE {
		lex()
		expr()
		if nextToken == RP_CODE {
			lex()
		} else {
			error()
		}
	}
	return factor
}
