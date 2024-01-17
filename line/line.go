package line

import "go/token"

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

	TEXT = ""

	IMAGE = "IMAGE"
)
