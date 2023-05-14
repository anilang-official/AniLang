package main

import (
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/anilang-official/AniLang/repl"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1][len(os.Args[1])-4:] == ".ani" {
			repl.ReplFile(os.Args[1], os.Stdout)
		} else {
			io.WriteString(os.Stdout, "Usage: anilang [filename.ani]\n")
		}
	} else {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Hello %s! This is the AniLang programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")

		repl.Start(os.Stdin, os.Stdout)
	}
}
