package ast

import "github.com/jeremitraverse/golb/token"

type Node interface {
	getTokens() []token.Token
}

type Post struct {
	Nodes []Node
}

type TitleNode struct {
	Type   token.TokenType
	Tokens []token.Token
}

type ImageNode struct {
	Type   token.TokenType
	Tokens []token.Token
}

type ListNode struct {
	Type   token.TokenType
	Tokens []token.Token
}

func (tn *TitleNode) getTokens() []token.Token {
	return tn.Tokens
}

func (in *ImageNode) getTokens() []token.Token {
	return in.Tokens
}

func (ln *ListNode) getTokens() []token.Token {
	return ln.Tokens
}
