# Lox — Learning Project

I decided to learn how to write programming languages so that I can write my own someday — a pretty ambitious project.

This is an implementation of the Lox programming language, a toy language from Robert Nystrom's book [_Crafting Interpreters_](https://craftinginterpreters.com). I am not writing the full interpreter, as that is not required for what I want to do. The focus here is the **lexer** and **parser**.

The original book implements Lox in Java. This implementation is written in **Go**.

Since Lox is a dynamically typed language, there is no type checking in this implementation. That is something I will learn separately.

---

## What is a Lexer and Parser?

A **lexer** (also called a scanner) reads raw source code as text and breaks it into tokens — the smallest meaningful units of a language, like keywords, identifiers, numbers, and symbols.

A **parser** takes those tokens and builds a tree structure that represents the grammatical meaning of the program. That tree is what a compiler or interpreter actually works with.

---

## Project Structure

```
lox/
├── go.mod
├── main.go
└── internal/
    ├── scanner/
    │   └── scanner.go
    ├── parser/
    │   └── parser.go
    └── token/
        └── token.go
```

---

## Running the Project

**Interactive mode** — type Lox code directly in the terminal:

```
go run main.go
```

**File mode** — run a `.lox` source file:

```
go run main.go script.lox
```

---

## Resources

- [_Crafting Interpreters_ by Robert Nystrom](https://craftinginterpreters.com) — available for purchase
- [The Go Programming Language](https://go.dev)
