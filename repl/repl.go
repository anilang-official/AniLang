package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/anilang-official/AniLang/evaluator"
	"github.com/anilang-official/AniLang/lexer"
	"github.com/anilang-official/AniLang/object"
	"github.com/anilang-official/AniLang/parser"
)

const PROMPT = ">> "

// Start starts the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := strings.TrimSpace(scanner.Text())

		if line == ".exit" {
			fmt.Println("Bye 👋")
			return
		}

		if line[len(line)-1] == '{' {
			// multiline input
			for {
				fmt.Print("... ")
				scanned := scanner.Scan()
				if !scanned {
					return
				}
				line += "\n" + strings.TrimSpace(scanner.Text())
				if line[len(line)-1] == '}' {
					break
				}
			}
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func ReplFile(filename string, out io.Writer) {
	env := object.NewEnvironment()

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	l := lexer.New(string(fileContent))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(io.Writer(os.Stderr), p.Errors())
		return
	}

	evaluator.Eval(program, env)
	io.WriteString(out, "\n")
}
