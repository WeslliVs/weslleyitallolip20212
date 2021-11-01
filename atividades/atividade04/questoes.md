## Atividade 04

> Correção: 1,1

#### Aluno: Weslley Itallo Vieira da Silva
#### Matrícula: 471699

<br/>

__1. Prove que a seguinte gramática é ambígua:__
```
<S> → <A>
<A> → <A> + <A> | <id>
<id> → a | b | c
```
__Resposta:__ Para a sentença: a + b + c, temos pelo menos duas árvores com derivação mais à esquerda.
```
<S> → <A>
<S> → <A> + <A>
<S> → <A> + <A> + <A>
<S> → <id> + <A> + <A>
<S> → a + <A> + <A>
<S> → a + <id> + <A>
<S> → a + b + <A>
<S> → a + b + <id>
<S> → a + b + c
```
![Árvore de Análise Sintática - Primeira Derivação](https://github.com/WeslliVs/weslleyitallolip20212/blob/main/atividades/atividade04/Att04-Diag1.png?raw=true)
```
<S> → <A>
<S> → <A> + <A>
<S> → <id> + <A>
<S> → a + <A>
<S> → a + <A> + <A>
<S> → a + <id> + <A>
<S> → a + b + <A>
<S> → a + b + <id>
<S> → a + b + c
```
![Árvore de Análise Sintática - Segunda Derivação](https://github.com/WeslliVs/weslleyitallolip20212/blob/main/atividades/atividade04/Att04-Diag2.png?raw=true)

<br/>
<br/>

> Correção: 0,5
>
> Tudo OK!!!

__2. Transforme a seguinte gramática em EBNF:__
```
<program> → begin <stmt_list> end 
<stmt_list> → <stmt>
            | <stmt> ; <stmt_list> 
<stmt> → <var> = <expression>
<var> → A | B | C 
<expression> → <var> + <var>
             | <var> – <var>
             | <var>
```
__Resposta:__
```
<program> → begin <stmt_list> end 
<stmt_list> → <stmt> {<stmt> ; <stmt_list>}
<stmt> → <var> = <expression>
<var> → A | B | C
<expression> → <var> {(+|-) <var>}
```

> Correção: 0,3
>
> ```<stmt> {<stmt> ; <stmt_list>}```
>
> era para ser
>
> ```<stmt> {;<stmt>}```

__3. Transforme a seguinte gramática em EBNF:__
```
<assign> → <id> = <expr> 
<id> → A|B|C
<expr> → <expr> + <expr>
       | <expr> * <expr>
       | (<expr>)
       | <id>
```
__Resposta:__
```
<assign> → <id> = <expr> 
<id> → A|B|C
<expr> → <expr> {(+|*) <expr>}
       | (<expr>)
       | <id>
```

> Correção: 0,3
>
> ```<expr> {(+|*) <expr>}```
>
> era para ser:
>
> ```<expr> [(+|*) <expr>]```
>
> pois a segunda parte precisa aparecer pelo menos uma vez.
            
__4. Escreva uma gramática de atributos considerando a seguinte gramática:__
```
<assign> → <var> = <expr> 
<expr> → <var> + <var>
       | <var>
<var> → A|B|C
```
__Resposta:__
```
1. Regra Sintática: <assign> -> <var> = <expr>
   Regra Semântica: <expr>.actual_type <- <var>.actual_type
2. Regra Sintática: <expr> → <var>[2] + <var>[3]
   Regra Semântica: <expr>.actual_type <- 
                          if (<var>[2].actual_type = int) and (<var>[3].actual_type = int)
                             then int
                          else real
                          end if
   Predicado: <var>[2].actual_type == <var>[3].actual_type
3. Regra Sintática: <expr> → <var>
   Regra Semântica: <expr>.actual_type <- <var>.actual_type
4. Regra Sintática: <var> → A|B|C
   Regra Semântica: <var>.actual_type <- look-up(<var>.string)
```
            

> Correção: 0,0
