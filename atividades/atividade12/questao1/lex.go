// Correção: 0,5. Não executa, analisa um arquivo diferente (word.txt?!?)
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

//Global Declarations
var charClass int
var lexeme string
var nextChar string
var lexLen int
var token int
var nextToken int

//Character Classes
var LETTER = 0
var DIGIT = 1
var UNKNOWN = 99

//Token Codes
var INT_LIT = 10
var IDENT = 11
var ASSIGN_OP = 20
var ADD_OP = 21
var SUB_OP = 22
var MULT_OP = 23
var DIV_OP = 24
var LEFT_PAREN = 25
var RIGHT_PAREN = 26

func main() {
	in_fp, err := ioutil.ReadFile("words.txt")

	if err != nil {
		log.Fatal(err)
	}

	getChar()

	for nextToken != EOF {
		lex()
	}
}

func lookup(string ch) {
	switch ch {
	case '(':
		addChar()
		nextToken = LEFT_PAREN
	case ')':
		addChar()
		nextToken = RIGHT_PAREN
	case '+':
		addChar()
		nextToken = ADD_OP
	case '-':
		addChar()
		nextToken = SUB_OP
	case '*':
		addChar()
		nextToken = MULT_OP
	case '/':
		addChar()
		nextToken = DIV_OP
	default:
		addChar()
		nextToken = EOF
	}
	return nextToken
}

func addChar(){
	if (lexLen <= 98){
		lexeme[lexLen++] = nextChar
		lexeme[lexLen] = 0
	} else {
		fmt.Printf ("Error - Lexeme is too long \n")
	}
}

func getChar(){
	nextChar = getc(in_fp)

	if (isalpha != EOF){
		if isalpha(nextChar){
			charClass = LETTER
		}else if (isdigit(nextChar)){
			charClass = DIGIT
		} else {
			charClass = UNKNOWN
		}
	} else {
		charClass = EOF
	}
}

func getNonBlank(){
	for isspace(nextChar){
		getChar()
	}
}

func lex(){
	lexLen = 0
	getNonBlank()
	
	switch charClass{
	case LETTER:
		addChar()
		getChar()
		for (charClass == LETTER || charClass == DIGIT){
			addChar()
			getChar()
		}
		nextToken = IDENT
	
	case DIGIT:
		addChar()
		getChar()

		for (charClass == DIGIT){
			addChar()
			getChar()
		}
		nextToken = INT_LIT

	case UNKNOWN:
		lookup(nextChar)
		getChar()
	
	case EOF:
		nextToken = EOF
		lexeme[0] = 'E'
		lexeme[1] = 'O'
		lexeme[2] = 'F'
		lexeme[3] = 0
	}
	fmt.Printf("Next Token is: %d, Next lexeme is %s \n", nextToken, lexeme)
	return nextToken
}