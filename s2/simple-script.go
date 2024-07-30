package main

import (
	"bufio"
	"fmt"
	"os"
	"play/lib"
	"strings"
)

type SimpleScript struct {
	variables map[string]int
	verbose   bool
}

func NewSimpleScript() *SimpleScript {
	return &SimpleScript{variables: make(map[string]int), verbose: true}
}

func main() {
	parser := &lib.SimpleParser{}
	script := NewSimpleScript()
	reader := bufio.NewReader(os.Stdin)

	scriptText := ""
	fmt.Print("\n>")

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		line = strings.TrimSpace(line)
		if line == "exit();" {
			fmt.Println("good bye!")
			break
		}
		scriptText += line + "\n"
		if strings.HasSuffix(line, ";") {
			tree := parser.Parse(scriptText)
			if script.verbose {
				parser.DumpAST(tree, "")
			}
			parser.EvaluateNode(tree, "")
			fmt.Print("\n>")
			scriptText = ""
		}
	}
}
