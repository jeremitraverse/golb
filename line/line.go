package line

import "github.com/jeremitraverse/golb/token"

type LineType string

type Line struct {
	Type   LineType
	Tokens []token.Token
}

const (
	FIRST_TITLE  = "#"
	SECOND_TITLE = "##"
	THIRD_TITLE  = "###"
	FOURTH_TITLE = "####"

	TEXT = "TEXT"

	IMAGE = "IMAGE"

	BREAK = "BREAK"

	EOF = "EOF"

	CODE = "CODE"
)
