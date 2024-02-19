package lexer

import (
	"errors"

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

	if l.currentPos == len(l.input) {
		return line.Line{ Type: line.EOF }
	}

	switch l.currentChar {
	case '#':
		li, err := l.getTitleLine()

		if err != nil {
			return l.getTextLine()
		}

		li.Tokens = l.getLineTokens()
		return li
	case '<':
		li, err := l.getImageLine()

		if err != nil {
			return l.getTextLine()
		}

		return li
	case '`':
		li, err := l.getCodeLine()

		if err != nil {
			return l.getTextLine()	
		}

		return li
	case 10:
		l.readChar()
		return line.Line{ Type: line.BREAK }
	case 13:
		l.readChar()
		return line.Line{ Type: line.BREAK }
	default:
		return l.getTextLine()
	}
}

func (l *Lexer) getTextLine() line.Line {
	var li line.Line

	li.Type = line.TEXT
	li.Tokens = l.getLineTokens()

	return li
}

func (l *Lexer) getTitleLine() (line.Line, error) {
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

func (l *Lexer) getImageLine() (line.Line, error) {
	var li line.Line	
	li.Type = line.IMAGE

	token, err := l.getImageToken()

	if err != nil {
		return li, err	
	}
	
	li.Tokens = append(li.Tokens, token)	

	return li, nil
}

func (l *Lexer) getLineTokens() []token.Token {
	var tokens []token.Token

	currentToken := l.getToken()

	for currentToken.Type != token.EOF && currentToken.Type != token.EOL {
		tokens = append(tokens, currentToken)
		currentToken = l.getToken()
	}
	
	return tokens
}

func (l *Lexer) getToken() token.Token {
	var tok token.Token
	switch l.currentChar {
		case 0:
			tok.Type = token.EOF
			tok.Literal = ""
			return tok
		case 10:
			tok.Type = token.EOL
			tok.Literal = ""
			l.readChar() // Reading char to skip to next line
			return tok
		case 13:
			tok.Type = token.EOL
			tok.Literal = ""
			l.readChar() // Reading char to skip to next line
			return tok
		case '*':
			return l.getModifiedTextToken('*')
		case '_':
			return l.getModifiedTextToken('_')
		default:
			return l.getTextToken()	
	}
}

func (l *Lexer) getCodeLine() (line.Line, error) {
	var li line.Line
	initialPost := l.currentPos

	for i := 0; i == 2; i++ {
		if i < 2 && l.peekChar() == '`' {
			l.readChar()	
		} else if i == 2 && l.peekChar() == '\n' {
			l.readChar()	
		} else {
			l.reset(initialPost)
			return li, errors.New("Not a code block")
		}
	}
	
	li.Type = line.CODE

	return li, nil
}

func (l *Lexer) isEndOfFile() bool {
	return l.currentChar == 0
}

func (l *Lexer) isBold(mod byte) bool {
	return l.peekChar() == mod && l.currentChar == mod
}

func (l *Lexer) isContentDelimiter() bool {
	return isReturnLine(l.currentChar) || l.isEndOfFile()
}

func (l *Lexer) getModifiedTextToken(mod byte) token.Token {
	var tok token.Token

	if l.peekChar() == mod {
		tok.Type = token.BOLD
		l.readChar()
		l.readChar()
	} else {
		tok.Type = token.ITALIC
		l.readChar()
	}

	textStartingPos := l.currentPos

	for !l.isContentDelimiter() {
		if l.currentChar == mod {
			if l.peekChar() == mod  && tok.Type == token.BOLD {
				tok.Literal = l.input[textStartingPos: l.currentPos]
				// Consuming the two mod char
				l.readChar()
				l.readChar()
				return tok
			}

			if l.peekChar() != mod && tok.Type == token.BOLD {
				tok.Type = token.ITALIC
				tok.Literal = l.input[textStartingPos-1: l.currentPos]
				// Consuming the mod char
				l.readChar()
				return tok
			}

			tok.Literal = l.input[textStartingPos: l.currentPos]
			// Consuming the mod char
			l.readChar()
			return tok
		}

		l.readChar()
	}

	tok.Type = token.TEXT
	tok.Literal = l.input[textStartingPos-1: l.currentPos]

	return tok
}

func (l *Lexer) getTextToken() token.Token {
	initialPos := l.currentPos

	for isCharText(l.currentChar) {
		l.readChar()
	}

	return token.Token{ Type: token.TEXT, Literal: l.input[initialPos:l.currentPos] }
}

func (l *Lexer) getImageToken() (token.Token, error) {
	var tok token.Token	
	initPos := l.currentPos
	
	for !l.isContentDelimiter() {
		if l.currentChar == '>' {
			tok.Type = token.IMAGE
			tok.Literal = l.input[initPos + 1: l.currentPos]
			l.readChar() // consuming '>' char 
			l.readChar() // consuming '\n' char 
			return tok, nil
		}

		l.readChar()
	}

	l.reset(initPos)

	return tok, errors.New("Image line not properly ended.") 
}

func (l *Lexer) reset(position int) {
	l.currentChar = l.input[position]
	l.currentPos = position
	l.nextPos = position + 1
}

func isCharText(char byte) bool {
	return char != '*' && char != '_' && !isReturnLine(char) && char != 0
}

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

// Not an extension of lexer since we want to be able to check the peek char
func isReturnLine(b byte) bool {
	return b == 10 || b == 13
}
