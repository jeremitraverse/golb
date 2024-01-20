package lexer

import (
	"errors"
	"fmt"

	"github.com/jeremitraverse/golb/line"
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

func (l *Lexer) GetLine() line.Line {
	var li line.Line
	switch l.currentChar {
	// When the first character of a line is a '#' it can either be a title line
	// or a text line
	case '#':
		li, err := l.getTitleLineType()
		if err != nil {
			fmt.Println(err)
			li.Type = line.TEXT
			li.Tokens = l.getLineTokens()
		}

		li.Tokens = l.getLineTokens()
		return li
	}

	return li
}

func (l *Lexer) getTitleLineType() (line.Line, error) {
	initialPos := l.currentPos
	var li line.Line

	for l.currentChar == '#' {
		l.readChar()
	}

	titleLit := l.input[initialPos:l.currentPos]

	if l.currentChar != ' ' || len(titleLit) > 4 {
		l.reset(initialPos)
		return li, errors.New("Line isn't a title")
	}
	
	switch titleLit {
		case "#":
			li.Type = line.FIRST_TITLE
		case "##":
			li.Type = line.SECOND_TITLE
		case "###":
			li.Type = line.THIRD_TITLE
		case "####":
			li.Type = line.FOURTH_TITLE
	}

	// consuming the ' ' after the title's '#'
	l.readChar()
	return li, nil
}

func (l *Lexer) getLineTokens() []token.Token {
	var tokens []token.Token

	currentToken := l.GetToken()
	for currentToken.Type != token.EOF {
		tokens = append(tokens, currentToken)
		currentToken = l.GetToken()
	}

	return tokens
}

/*
func (l *Lexer) GetToken() token.Token {
	var tok token.Token
	switch l.currentChar {
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
*/

func (l *Lexer) GetToken() token.Token {
	var tok token.Token

	switch l.currentChar {
		case 10:
			tok.Type = token.EOL
			tok.Literal = ""
			return tok
		case 13:
			tok.Type = token.EOL
			tok.Literal = ""
			return tok
	}
		
	for l.currentChar != 10 || l.currentChar != 13 {
		switch l.currentChar {
			case '*':
				t, err := l.isModifiedText()
		}
	}
}

func (l *Lexer) isModifiedText() (token.Token, error) {
	var tok token.token

	return tok, nil
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


// **BOLD** { Type: BOLD, Value: "BOLD" }
// *ITALIC* { Type: ITALIC, Value: "ITALIC" }
// **ITALIC* { Type: ITALIC, Value: "ITALIC" }
func (l *Lexer) handleTextModifier(mod byte) token.Token {
	var tok token.Token
	
	textStartingPos := l.currentPos
	
	if l.peekChar() == mod {
		tok.Type = token.BOLD
	} else {
		tok.Type = token.ITALIC
	}
		
	for !isReturnLine(l.currentChar) && !l.isEndOfFile() {
		if l.currentChar 
	}

  /*
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
	*/
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

func (l *Lexer) reset(position int) {
	l.currentChar = l.input[position]
	l.currentPos = position
	l.nextPos = position + 1
}
