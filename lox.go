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
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil { // if there are no more lines
			break
		}
		run(line)
		hadError = false
	}
}

func run(source string) {
    // Scanner and tokens don't exist yet — coming in Chapter 4
    fmt.Println(source)
}

func showError(line int, message string) {
	report(line, "", message)
}

func report(line int, where, message string) {
	fmt.Fprintf("[line %d] Error%v: %v", line, where, message)
	hadError = true
}

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: jlox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}