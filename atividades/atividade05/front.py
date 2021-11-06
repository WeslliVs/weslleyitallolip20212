from types import NoneType
charClass = 0
lexeme = [100]
nextChar = ''
lexLen = 0
token = 0
nextToken = 0
arq = open(
    'C:/Users/wesll/Documents/Development/Python/LIP/Atividade 05/front.in')
leitura = arq.read(100)
null = NoneType
# Classes de Caracteres
LETTER = 0
DIGIT = 1
UNKNOWN = 99
# Códigos Tokens
INT_LIT = 10
IDENT = 11
ASSIGN_OP = 20
ADD_OP = 21
SUB_OP = 22
MULT_OP = 23
DIV_OP = 24
LEFT_PAREN = 25
RIGHT_PAREN = 26
EOF = 190
#------------------------------------------------------------------------------#
# função main - abrir o arquivo e processar seu conteúdo


def main():
    if arq.read(1) == null:
        print("ERROR - cannot open front.in \n")
    else:
        getChar()
        while nextToken != EOF:
            print("main -> lex")
            lex()
        print("main <- lex")
    return main
#------------------------------------------------------------------------------#
# função lookup - verifica os operadores e parênteses e retorna o token


def lookup(ch):
    if ch == '(':
        addChar()
        nextToken = LEFT_PAREN
    elif ch == ')':
        addChar()
        nextToken = RIGHT_PAREN
    elif ch == '+':
        addChar()
        nextToken = ADD_OP
    elif ch == '-':
        addChar()
        nextToken = SUB_OP
    elif ch == '*':
        addChar()
        nextToken = MULT_OP
    elif ch == '/':
        addChar()
        nextToken = DIV_OP
    else:
        addChar()
        nextToken = EOF
    return lookup
#------------------------------------------------------------------------------#
# função addChar - adiciona o próximo char ao lexema


def addChar():
    if lexLen <= 98:
        lexeme[lexLen] = nextChar
        lexeme[lexLen] = 0
    else:
        print("Error - lexeme is too long \n")
    return addChar
#------------------------------------------------------------------------------#
# função getChar - pega o próximo char da entrada e determina sua classe


def getChar():
    if (nextChar == arq.read(1)) != '':
        if nextChar.isalpha:
            charClass = LETTER
        elif nextChar.isdigit:
            charClass = DIGIT
        else:
            charClass = UNKNOWN
    else:
        charClass = EOF
    return getChar
# -----------------------------------------------------------------------------#
# função getNonBlank - chama getChar enquanto não retornar um espaço em branco


def getNonBlank():
    while (nextChar == nextChar.isspace):
        getChar()
        nextChar+1
    return getNonBlank
#------------------------------------------------------------------------------#
# função lex - analisador léxico para expressões aritméticas


def lex():
    lexLen = 0
    getNonBlank()
    if charClass == LETTER:
        addChar()
        getChar()
        while(charClass == LETTER or charClass == DIGIT):
            addChar()
            getChar()
        nextToken = IDENT
    elif charClass == DIGIT:
        addChar()
        getChar()
        while(charClass == DIGIT):
            addChar()
            getChar()
        nextToken = INT_LIT
    elif charClass == UNKNOWN:
        lookup(nextChar)
        getChar()
    elif charClass == EOF:
        nextToken = EOF
        lexeme = 'EOF0'
    print("Next token is: %u, Next lexeme is: %s" % (nextToken, lexeme))
    return nextToken and lex


# ------------------------------------------------------------------------------
main()
