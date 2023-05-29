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

var PROMPT = ">> "

// Start starts the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	var BRACECOUNTER int = 0
	var line string = ""

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		if BRACECOUNTER > 0 {
			line = line + "\n" + strings.TrimSpace(scanner.Text())
		} else {
			line = strings.TrimSpace(scanner.Text())
		}

		if line == ".exit" && BRACECOUNTER == 0 {
			fmt.Println("Bye 👋")
			return
		}

		if line[len(line)-1] == '{' {
			BRACECOUNTER++
			PROMPT = "... "
		}

		if line[len(line)-1] == '}' {
			BRACECOUNTER--
			if BRACECOUNTER == 0 {
				PROMPT = ">> "
			}
		}

		if BRACECOUNTER > 0 {
			continue
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
		printParserErrors(out, p.Errors())
		return
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		if evaluated.Type() == object.ERROR_OBJ {
			io.WriteString(out, evaluated.Inspect())
		}
	}
	io.WriteString(out, "\n")
}
