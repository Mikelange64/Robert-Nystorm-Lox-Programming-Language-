package main

import (
	"fmt";
	"os";
	"bufio"
)

var hadError bool = false

func runFile(path string) error {
	bytes, err := os.ReadFile(path)
	// we read the file as bytes then explicitly convert it to strings later.
	// If we read it as strings, the system would assume the encoding for the file
	// Here we have more control, raw bytes first, then explicit conversion to text (UTF-8)
	if err != nil {
		return err
	}
	run(string(bytes))

	if hadError {
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
        run(line)
        hadError = false
    }
}

func run(source string) {
    scanner := CreateScanner(source)
    tokens := scanner.ScanTokens()

    for _, token := range tokens {
    	fmt.Println(token)
    }
}

func showError(line int, message string) {
	report(line, "", message)
}

func report(line int, where, message string) {
	// full-featured languages have multiple ways of displaying errors; stderr, IDE error window, etc.
	// Ideally we'd use an ErrorReporter stuct/interface of some sort that can be passed around
	// We're not doing that here though,
	fmt.Fprintf(os.Stderr, "[line %d] Error%v: %v", line, where, message)
	hadError = true
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
