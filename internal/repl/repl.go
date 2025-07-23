package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Michaelpalacce/go-interperter.git/internal/lexer"
	"github.com/Michaelpalacce/go-interperter.git/internal/token"
)

// PROMPT is what's shown at the start of each new REP loop.
const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
