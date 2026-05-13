## Precedence rules for the Lox programming language

Our parser will be build on the idea of precedence and recursions. For an expression to evaluated correctly, we need to define rules
around the order in which operations are evaluated. That's something we already to in every day math. Like the old `PEMDAS` rule.
This tells us what do evalaluate first rather than just doing a left to right evaluation: `2 * (7 + 5) - 2^3/4` has to follow the
precedence order. `(7 + 5)` first, then `2^3` giving us `2 * 12 - 8/4`. From then on we have multiplication and division on the same
level so we do `2 * 12` and `8 / 4` and we have `24 - 2` which is `22`

Our parser follows this order of precedence

```
expression  ->  equality ;
equality    ->  comparison ( ( "!=" | "==" ) comparison )* ;
comparison  ->  term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
term        ->  factor ( ( "-" | "+" ) unary )* ;
factor      ->  unary ( ( "/" | "*" ) unary )* ;
unary       ->  ( "!" | "-" ) unary | primary ;
primary     ->  NUMBER | STRING | "true" | "false" | "nil"
                | "(" expression ")" ;
```

The trick for our parser is that each "level" is represented by a function. And each token starts being evaluated at the lowest precendence
level. Each function calls the level right above it so that higher level are evaluated first. `expression()` calls `equality()`, which calls
`comparison()`, which calls `term()` etc. until it reaches the top `primary()` This returns an AST already implemented with all the correct
evaluation order
