package main

import (
	"fmt"
	"play/lib"
)

func dump(tokenReader *lib.SimpleTokenReader) {
	fmt.Println("text\ttype")
	for {
		token := tokenReader.Read()
		if token == nil {
			break
		}
		fmt.Printf("%s\t\t%s\n", token.Text, token.Type.String())
	}
}

func main() {
	lexer := &lib.SimpleLexer{}

	script := "int age = 45;"
	fmt.Println("parse :", script)
	tokenReader := &lib.SimpleTokenReader{Tokens: lexer.Tokenize(script)}
	dump(tokenReader)

	script = "inta age = 45;"
	fmt.Println("\nparse :", script)
	tokenReader = &lib.SimpleTokenReader{Tokens: lexer.Tokenize(script)}
	dump(tokenReader)

	script = "in age = 45;"
	fmt.Println("\nparse :", script)
	tokenReader = &lib.SimpleTokenReader{Tokens: lexer.Tokenize(script)}
	dump(tokenReader)

	script = "age >= 45;"
	fmt.Println("\nparse :", script)
	tokenReader = &lib.SimpleTokenReader{Tokens: lexer.Tokenize(script)}
	dump(tokenReader)

	script = "age > 45;"
	fmt.Println("\nparse :", script)
	tokenReader = &lib.SimpleTokenReader{Tokens: lexer.Tokenize(script)}
	dump(tokenReader)
}
