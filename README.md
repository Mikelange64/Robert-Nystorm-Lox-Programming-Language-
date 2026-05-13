# Lox — Learning Project

I decided to learn how to write programming languages so that I can write my own someday, a pretty ambitious project.

This is an implementation of the Lox programming language, a toy language from Robert Nystrom's book [_Crafting Interpreters_](https://craftinginterpreters.com). The original book implements Lox in Java, but I don't know (or like) Java, so this implementation is written in **Go**.

---

## Current State

This project currently implements the first three stages of a language pipeline:

**Scanner** — reads raw source code and breaks it into tokens. Handles single and multi-character operators, string and number literals, keywords, identifiers, comments (line and block), and whitespace. Reports all errors found before stopping rather than crashing on the first one.

**Parser** — takes the token list and builds an Abstract Syntax Tree (AST) that represents the grammatical structure of the program. Implements a recursive descent parser with correct operator precedence and associativity. Recovers from errors and attempts to keep parsing rather than stopping at the first mistake.

**Interpreter** — walks the AST and evaluates expressions. Handles arithmetic, string concatenation, comparison, equality, unary operations, and grouping. Reports runtime errors with line information.

---

## What is a Scanner, Parser, and Interpreter?

A **scanner** (also called a lexer) reads raw source code as text and breaks it into tokens, the smallest meaningful units of a language, like keywords, identifiers, numbers, and symbols.

A **parser** takes those tokens and builds a tree structure that represents the grammatical meaning of the program. Operator precedence is baked into the structure of the tree, multiplication nodes sit deeper than addition nodes, so they evaluate first.

An **interpreter** walks that tree and computes the result. It recurses down to the leaves (literal values) and works back up, computing each operation along the way.

---

## Project Structure

```
lox/
├── lox_test_files/     -- sample .lox files for testing
├── notes/              -- personal notes on concepts covered
├── ast_printer.go      -- prints the AST as a string for debugging
├── expr.go             -- AST node types (Binary, Unary, Literal, Grouping)
├── interpreter.go      -- tree-walking interpreter
├── lox.go              -- entry point, REPL, and file runner
├── parser.go           -- recursive descent parser
├── runtime_error.go    -- runtime error type
├── scanner.go          -- lexer
├── token.go            -- Token struct
├── token_types.go      -- TokenType enum and Stringer
├── go.mod
└── README.md
```

---

## Running the Project

**Interactive mode** — type Lox code directly in the terminal:

```
go run .
```

**File mode** — run a `.lox` source file:

```
go run . script.lox
```

---

## Resources

- [_Crafting Interpreters_ by Robert Nystrom](https://craftinginterpreters.com) : available for purchase
- [The Go Programming Language](https://go.dev)
