## Atividade 03
#### Aluno: Weslley Itallo Vieira da Silva
#### Matrícula: 471699

<br/>

__1. Usando a gramática do exemplo 3.4 no livro texto, mostre uma árvore de análise sintática e uma derivação mais à esquerda para cada uma das seguintes sentenças:__
 1. A = A + (B + C)
 2. A = B * (C * (A + B))
  
  
Resposta: A = A + (B + C)
```
<assign> -> <id> = <expr>
<assign> -> A = <expr>
<assign> -> A = <expr> + <term>
<assign> -> A = <term> + <term>
<assign> -> A = <factor> + <term>
<assign> -> A = <id> + <term>
<assign> -> A = A + <term>
<assign> -> A = A + <factor>
<assign> -> A = A + (<expr>)
<assign> -> A = A + (<expr> + <term>)
<assign> -> A = A + (<term> + <term>)
<assign> -> A = A + (B + <term>)
<assign> -> A = A + (B + C)
```
![Árvore de Análise Sintática: A = A + (B + C)](https://github.com/WeslliVs/weslleyitallolip20212/blob/main/atividades/atividade03/Att03%20-%20Diag1.png?raw=true)  

A = B * (C * (A + B))
```
<assign> -> <id> = <expr>
<assign> -> A = <term>
<assign> -> A = <term> * <factor>
<assign> -> A = <factor> * <factor>
<assign> -> A = <id> * <factor>
<assign> -> A = B * <factor>
<assign> -> A = B * (<expr>)
<assign> -> A = B * (<term>)
<assign> -> A = B * (<term> * <factor>)
<assign> -> A = B * (<factor> * <factor>)
<assign> -> A = B * (<id> * <factor>)
<assign> -> A = B * (C * <factor>)
<assign> -> A = B * (C * (<expr>)
<assign> -> A = B * (C * (<expr> + <term>))
<assign> -> A = B * (C * (<term> + <term>))
<assign> -> A = B * (C * (<factor> + <term>))
<assign> -> A = B * (C * (<id> + <term>))
<assign> -> A = B * (C * (A + <term>))
<assign> -> A = B * (C * (A + <factor>))
<assign> -> A = B * (C * (A + <id>))
<assign> -> A = B * (C * (A + B))
```
![Árvore de Análise Sintática: A = B * (C * (A + B))](https://github.com/WeslliVs/weslleyitallolip20212/blob/main/atividades/atividade03/Att03%20-%20Diag2.png?raw=true)  

<br/>

__2. Escreva uma derivação mais à direita para a cadeia (a 23 (m x y)) e construa uma árvore de análise sintática.__  
Resposta: 
```
<lexp> -> <lista>
<lexp> -> (<lexp-seq>)
<lexp> -> (<lexp-seq> <lexp>)
<lexp> -> (<lexp-seq> <lista>)
<lexp> -> (<lexp-seq> (<lexp-seq>))
<lexp> -> (<lexp-seq> (<lexp-seq> <lexp>))
<lexp> -> (<lexp-seq> (<lexp-seq> <átomo>))
<lexp> -> (<lexp-seq> (<lexp-seq> <identificador>))
<lexp> -> (<lexp-seq> (<lexp-seq> y))
<lexp> -> (<lexp-seq> (<lexp-seq> <lexp> y))
<lexp> -> (<lexp-seq> (<lexp-seq> <átomo> y))
<lexp> -> (<lexp-seq> (<lexp-seq> <identificador> y))
<lexp> -> (<lexp-seq> (<lexp-seq> x y))
<lexp> -> (<lexp-seq> (<lexp> x y))
<lexp> -> (<lexp-seq> (<átomo> x y))
<lexp> -> (<lexp-seq> (<identificador> x y))
<lexp> -> (<lexp-seq> (m x y))
<lexp> -> (<lexp-seq> <lexp> (m x y))
<lexp> -> (<lexp-seq> <átomo> (m x y))
<lexp> -> (<lexp-seq> <número> (m x y))
<lexp> -> (<lexp-seq> 23 (m x y))
<lexp> -> (<lexp> 23 (m x y))
<lexp> -> (<átomo> 23 (m x y))
<lexp> -> (<identificador> 23 (m x y))
<lexp> -> (a 23 (m x y))
```
