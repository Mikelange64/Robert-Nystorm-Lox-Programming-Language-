# Context-Free_Grammar

The scanner we built was for parsing raw text and breaking it down into valid lexemes and tokens

For the next part we're gonna be building the languages **context-free grammar**, which uses the tokens are from the scanner and composes them into an 
arbitrary numbers of expressions. Formal grammars usually have two aspects to them :

- The alphabet : The smallest, valid atomic pieces that it works on
- The Strings  : A set of (usually finite) combination of letters in the alphabet

Comparison table
| **Terminology**  | **Lexical grammar** | **Syntactic grammar** |
|------------------|---------------------|-----------------------|
| The "alphabet"   | Characters          | Tokens                | 
| The "strings"    | Lexeme or token     | Expression            |
| Implemented by   | Scanner             | Parser                |

Since we can't list out all possible strings the way we did with lexemese we define rule, formally called **productions** that tell us what is and is 
not allowed to happens. Each rule contains a **head** (its name) and a **body** (the description). A **terminal** are the letters from the alphabet 
(here the tokens, literal values) and a **nonterminal** is reference to another rule.
 
### Example
```
breakfast  -> protein "with" breakfast "on the side"
breakfast  -> bread

protein    -> crispiness "crispy" "bacon"
protein    -> "sausage"
protein    -> cooked "eggs"

crispiness -> "really"
crispiness -> "really" crispiness

cooked     -> "scrambled"
cooked     -> "poached"
cooked     -> "fried"

bread      -> "toast"
bread      -> "biscuits"
bread      -> "English muffin"
```

We can optimize by writing the head once only, and separating each option with a `|` e.g. ` breakfast  -> protein "with" breakfast "on the side" | bread ` 

We can get rid of **production** which only contains **terminals** and write each terminal in a set of parentheses separated by the same pipe operator
`|` e.g. the entire *cooked* description becomes ` protein -> ("scrambled" | "poached" | "fried") "eggs" ` and for repeating values like the *crispiness* 
description, we can use a `+` postfix to represent it ` "really"+ `. Then a `?` means the preceeding **production** happens at most once. Our revised version
looks like this

```
breakfast -> protein ("with" breakfast "on the side")? | bread

protein   ->    "really"+ "cripsy" "bacon" 
              | "sausage" 
              | ("scrambled" | "poached" | "fried") "eggs"
              
bread     -> "toast" | "biscuits" | "English muffin"
```

The Grammar for Lox so far will be
```
expression -> literal | unary | binary | operator | grouping ;
literal    -> NUMBER | STRING | "true" | "false" | "nil" ;
grouping   -> "(" expression ")" ;
unary      -> ( "-" | "!") expression ;
binary     -> expression operator expression ;
operator   -> "+" | "-" | "/" | "*" | "==" | "!=" | "<" | ">" | "<=" | >=
```


