# Correção: 1,0
def expr():
    term()
    while (nextToken == ADD_OP | nextToken == SUB_OP):
        lex()
        term()
    return expr


def term():
    print("Enter <term>: ")
    factor()
    while (nextToken == MULT_OP | nextToken == DIV_OP):
        lex()
        factor()
    print("Exit <term> \n")
    return term


def factor():
    if (nextToken == ID_CODE | nextToken == INT_CODE):
        lex()
    elif (nextToken == LP_CODE):
        lex()
        expr()
        if (nextToken == RP_CODE):
            lex()
        else:
            error()
    return factor
