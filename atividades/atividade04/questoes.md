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
__Resposta:__
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
![Árvore de Análise Sintática - Primeira Derivação](?raw=true)
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
