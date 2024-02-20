package parser

import (
	"strings"

	"github.com/jeremitraverse/golb/line"
	"github.com/jeremitraverse/golb/token"
)

type Parser struct {
	currentLine line.Line
	currentToken token.Token
}

func New(li line.Line) Parser {
	return Parser{ currentLine: li }
}

func (p *Parser) ParseLine() string {
	lineType := p.currentLine.Type

	switch lineType {
		case line.FIRST_TITLE:
			lit := p.parseTextTokens()
			return createHtmlNode("h1", lit)
		case line.SECOND_TITLE:
			lit := p.parseTextTokens()
			return createHtmlNode("h2", lit)
		case line.THIRD_TITLE:
			lit := p.parseTextTokens()
			return createHtmlNode("h3", lit)
		case line.FOURTH_TITLE:
			lit := p.parseTextTokens()
			return createHtmlNode("h4", lit)
		case line.TEXT:
			lit := p.parseTextTokens()
			return createHtmlNode("div", lit)
		case line.IMAGE:
			return p.parseImageLine()	
		case line.BREAK:
			return "<br/>"
		case line.CODE:
			return p.parseCodeLine()
	}

	return ""
}

func (p *Parser) parseTextTokens() string {
	var sb strings.Builder
	
	for _, tok := range p.currentLine.Tokens {
		switch tok.Type {
			case token.TEXT:
				sb.WriteString(tok.Literal)
			case token.BOLD:
				node := createHtmlNode("b", tok.Literal) 
				sb.WriteString(node)
			case token.ITALIC:
				node := createHtmlNode("i", tok.Literal) 
				sb.WriteString(node)
		}
	}

	return sb.String()
}

func (p *Parser) parseImageLine() string {
	return "<img src=\"" + p.currentLine.Tokens[0].Literal + "\" />"
}

func (p *Parser) parseCodeLine() string {
	return getOpeningNode("code") + p.currentLine.Tokens[0].Literal + getClosingNode("code")
}

func getOpeningNode(nodeType string) string {
	return "<" + nodeType + ">"
}

func getClosingNode(nodeType string) string {
	return "</" + nodeType + ">"
}

func createHtmlNode(nodeType string, nodeLiteral string) string {
	var sb strings.Builder

	openingNode := getOpeningNode(nodeType)
	closingNode := getClosingNode(nodeType)
	node := openingNode + nodeLiteral + closingNode

	sb.WriteString(node)

	return sb.String()
}
