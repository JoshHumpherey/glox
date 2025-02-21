package main

import (
	"bufio"
	"fmt"
	"glox/scanner"
	"os"
)

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

func runFile(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	run(string(data))
	return nil
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
	}
}

func run(source string) {
	s := scanner.Scanner{Source: source}
	t := s.ScanTokens()

	for token := range t {
		fmt.Println(token)
	}
}
