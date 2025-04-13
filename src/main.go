package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Use: ./flowleaf <path-file>")
		return
	}

	filePath := os.Args[1]

	sourceCode, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error al abrir el archivo: %s\n", err)
		return
	}

	tokens := Tokenizer(string(sourceCode))

	PrintTokens(tokens)

	htmlCode, err := ParseTokens(tokens)
	if err != nil {
		fmt.Printf("Error for compiling: %s\n", err)
		return
	}

	fmt.Println("HTML generated:")
	fmt.Println(htmlCode)

	StartServer(filePath)
}
