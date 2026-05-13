package main

import (
	"bufio"
	"fmt"
	"os"
)

var interpreter Interpreter

func runFile(path string) error {
	bytes, err := os.ReadFile(path)
	// we read the file as bytes then explicitly convert it to strings later.
	// If we read it as strings, the system would assume the encoding for the file
	// Here we have more control, raw bytes first, then explicit conversion to text (UTF-8)
	if err != nil {
		return err
	}
	if err := run(string(bytes)); err != nil {
		os.Exit(65)
	}
	return nil
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		// Scan() waits for the next line and returns false on EOF/Error
		if !scanner.Scan() {
			break
		}

		line := scanner.Text() // This is already clean of \n or \r\n
		if err := run(line); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func run(source string) error {
	scanner := CreateScanner(source)
	tokens, errs := scanner.ScanTokens()
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Fprintln(os.Stderr, err)
		}
		return errs[0]
	}
	parser := CreateParser(tokens)
	expression, err := parser.Parse()
	if err != nil {
		return err
	}

	if err := interpreter.Interpret(expression); err != nil {
		return err
	}
	return nil
}

func main() {
	// expr := Binary{
	// 	Left : Unary {
	// 		Operator : Token{ Type : MINUS, Lexeme : "-", Literal : nil, Line : 1},
	// 		Right : Literal{Value : 123},
	// 	},
	// 	Operator : Token{ Type : STAR, Lexeme : "*", Literal : nil, Line : 1},
	// 	Right : Grouping{ Expression : Literal{Value : 45.67} },
	// }

	// fmt.Println(printExpr(expr))

	if len(os.Args) > 2 {
		fmt.Println("Usage: jlox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}
