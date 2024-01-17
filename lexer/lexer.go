package lexer

import (
	"errors"

	"github.com/jeremitraverse/golb/token"
)

type Lexer struct {
	input       string
	currentPos  int
	nextPos     int
	currentChar byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.nextPos = 1
	l.currentChar = l.input[l.currentPos]
	return l
}

/*
func (l *Lexer) GetLine() {
	var li line.Line
	switch l.currentChar {
	case '#':
		newToken, err := l.getTitleToken()
		if err != nil {
			panic(err)
		}
	}

}
*/

func (l *Lexer) GetToken() token.Token {
	var tok token.Token
	switch l.currentChar {
	case '#':
		newToken, err := l.getTitleToken()
		if err != nil {
			panic(err)
		}

		return newToken
	case '*':
		return l.handleTextModifier('*')
	case '_':
		return l.handleTextModifier('_')
	case '<':
		return l.handleImage()
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	case 10:
		tok.Type = token.EOL
		tok.Literal = ""
	case 13:
		tok.Type = token.EOL
		tok.Literal = ""
	default:
		if isText(l.currentChar) {
			tok.Type = token.TEXT
			tok.Literal = l.readText()
		}
	}

	l.readChar()
	return tok
}

// Advancing both char pointers
func (l *Lexer) readChar() {
	if l.nextPos >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.nextPos]
	}

	l.currentPos = l.nextPos
	l.nextPos += 1
}

func (l *Lexer) peekChar() byte {
	if l.nextPos >= len(l.input) {
		return 0
	}

	return l.input[l.nextPos]
}

func (l *Lexer) getTitleToken() (token.Token, error) {
	var tok token.Token
	var initialPos = l.currentPos

	for l.currentChar == '#' {
		l.readChar()
	}

	title := l.input[initialPos:l.currentPos]
	tok.Literal = title

	switch title {
	case "#":
		tok.Type = token.FIRST_TITLE
	case "##":
		tok.Type = token.SECOND_TITLE
	case "###":
		tok.Type = token.THIRD_TITLE
	case "####":
		tok.Type = token.FOURTH_TITLE
	default:
		return tok, errors.New("Title type not found")
	}

	// Consume the space between the title type and the title text
	l.readChar()

	return tok, nil
}

func isText(b byte) bool {
	return b != '_' && b != '*' && b != 0
}

// Not an extension of lexer since we want to be able to check the peek char
func isReturnLine(b byte) bool {
	return b == 10 || b == 13
}

func (l *Lexer) isEndOfFile() bool {
	return l.currentChar == 0
}

func (l *Lexer) isBold(mod byte) bool {
	return l.peekChar() == mod
}

func (l *Lexer) readText() string {
	initialPosition := l.currentPos
	for isText(l.currentChar) && !isReturnLine(l.peekChar()) {
		l.readChar()
	}

	endPos := l.currentPos

	if isReturnLine(l.peekChar()) {
		endPos++
	}

	return l.input[initialPosition:endPos]
}

func (l *Lexer) isContentDelimiter() bool {
	return isReturnLine(l.currentChar) || l.isEndOfFile()
}

func (l *Lexer) handleTextModifier(mod byte) token.Token {
	var tok token.Token
	textStartingPos := l.currentPos

	l.readChar()

	if l.currentChar == mod {
		// There's 2 char to skip
		textStartingPos += 2
		tok.Type = token.BOLD
		l.readChar()
	} else {
		// There's 1 char to skip
		textStartingPos += 1
		tok.Type = token.ITALIC
	}

	for !isReturnLine(l.currentChar) && !l.isEndOfFile() {
		if l.currentChar == mod {
			if l.isBold(mod) {
				tok.Literal = l.input[textStartingPos:l.currentPos]
				// Consume the two next char
				l.readChar()
				l.readChar()
				return tok
			}

			tok.Literal = l.input[textStartingPos:l.currentPos]
			// Consume the next char
			l.readChar()
			return tok
		}

		l.readChar()
	}

	return tok
}

func (l *Lexer) handleImage() token.Token {
	initialPos := l.currentPos
	var tok token.Token

	for !l.isContentDelimiter() {
		if l.currentChar == '>' {
			tok.Type = token.IMAGE
			tok.Literal = l.input[initialPos+1 : l.currentPos]

			// Consume ']'
			l.readChar()
			return tok
		}

		l.readChar()
	}

	tok.Type = token.TEXT
	tok.Literal = l.input[initialPos:l.currentPos]

	return tok
}
