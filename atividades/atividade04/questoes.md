## Atividade 04
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

<br/>

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

<br/>

__4. Escreva uma gramática de atributos considerando a seguinte gramática:__
```
<assign> → <var> = <expr> 
<expr> → <var> + <var>
       | <var>
<var> → A|B|C
```
__Resposta:__
```

```
