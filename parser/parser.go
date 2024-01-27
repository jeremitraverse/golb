package parser

import (
	"fmt"

	"github.com/jeremitraverse/golb/line"
	"github.com/jeremitraverse/golb/token"
)

type Parser struct {
	currentLine line.Line
	currentToken token.Token
}


func (p *Parser) ParseLine(li line.Line) {
	for tok := range li.Tokens {
		fmt.Println(tok)
	}
}
