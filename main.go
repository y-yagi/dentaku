package main

import (
	"fmt"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/y-yagi/goext/strext"
)

func main() {
	tmpfile, err := ioutil.TempFile("", "dentaku")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	out := os.Stdout
	l, err := readline.NewEx(&readline.Config{
		Prompt:      "> ",
		Stdout:      out,
		HistoryFile: tmpfile.Name(),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		expr, err := l.Readline()
		if err != nil {
			break
		}

		if strext.IsBlank(expr) {
			continue
		}

		expr = strings.Replace(expr, ",", "", -1)

		fset := token.NewFileSet()
		tv, err := types.Eval(fset, nil, token.NoPos, expr)
		if err != nil {
			fmt.Fprintf(out, "%v\n", err)
		} else {
			fmt.Fprintf(out, "%v\n", tv.Value)
		}
	}
}
