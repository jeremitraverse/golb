package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	FIRST_TITLE		= "#"
	SECOND_TITLE	= "##"
	THIRD_TITLE		= "###"
	FOURTH_TITLE	= "####"

	CODE = "CODE"

	TEXT     = "TEXT"
	NUM_LIST = "NUM_LIST"

	EOF = "EOF"
	EOL = "EOL"

	BOLD   = "BOLD"
	ITALIC = "ITALIC"

	IMAGE = "IMAGE"
)
