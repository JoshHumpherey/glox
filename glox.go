package main

import (
	"bufio"
	"fmt"
	"glox/scanner"
	"os"
)

const errorExit = 65

var hadError = false

func main() {
	args := os.Args[1:]
	switch len(args) {
	case 0:
		runPrompt()
	case 1:
		runFile(args[1])
	default:
		panic(fmt.Errorf("usage: glox [script]"))
	}
}

func runFile(filePath string) {
	data, _ := os.ReadFile(filePath)
	run(string(data))
	if hadError {
		os.Exit(errorExit)
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		if line == "" {
			return
		}
		run(line)
		hadError = false
	}
}

func run(source string) {
	s := scanner.Scanner{Source: source}
	t := s.ScanTokens()

	for token := range t {
		fmt.Println(token)
	}
}

func loxError(l int, msg string) {
	report(l, "", msg)
}

func report(l int, where, msg string) {
	fmt.Printf("[line %d] Error%s: %s", l, where, msg)
	hadError = true
}
