package repl

import (
	"bufio" //provides buffered I/O utilities
	"fmt"   //provides formatted I/O, like `Printf`, `Fprintf`, etc.
	"io"    //defines interfaces for I/O (Reader, Writer, etc.)
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)

		// It returns true if there *is* a line, false on EOF or error.
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		// scanner.Text() returns the text of the line that was just scanned.
		// This is the raw Monkey source code the user typed.
		line := scanner.Text()

		// Create a new lexer for this line of input.
		// The lexer will turn this line into a stream of tokens.
		l := lexer.New(line)

		// This inner `for` loop pulls tokens out of the lexer until it hits EOF.
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
