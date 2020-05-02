package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/benjamin-daniel/monkey/lexer"
	"github.com/benjamin-daniel/monkey/token"
)

// PROMPT is the presuffix for a new line
const PROMPT = ">> "

// Start basically starts the repl in the console
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
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
