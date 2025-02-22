package main

import (
	"bufio"
	"fmt"
	"glox/lox_error"
	"glox/scanner"
	"os"
)

const errorExit = 65

func main() {
	args := os.Args[1:]
	errorHandler := lox_error.ErrorHandler{}
	switch len(args) {
	case 0:
		runPrompt(errorHandler)
	case 1:
		runFile(args[1], errorHandler)
	default:
		panic(fmt.Errorf("usage: glox [script]"))
	}
}

func runFile(filePath string, errorHandler lox_error.ErrorHandler) {
	data, _ := os.ReadFile(filePath)
	run(string(data))
	if errorHandler.HadError {
		os.Exit(errorExit)
	}
}

func runPrompt(errorHandler lox_error.ErrorHandler) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		if line == "" {
			return
		}
		run(line)
		errorHandler.HadError = false
	}
}

func run(source string) {
	s := scanner.Scanner{Source: source}
	t := s.ScanTokens()

	for token := range t {
		fmt.Println(token)
	}
}
