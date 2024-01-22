package lexer

import (
	"testing"

	"github.com/jeremitraverse/golb/line"
	"github.com/jeremitraverse/golb/token"
)

func TestTitles(t *testing.T) {
	input := `# First Title
## Second Title
### Third Title
#### Fourth Title`
	
	testCases := []struct {
		expectedLineType line.LineType
		expectedTokens []token.Token
	}{
		{
			line.FIRST_TITLE, []token.Token {
				{ Type: token.TEXT, Literal: "First Title" },
	 		},
		},
		{
			line.SECOND_TITLE, []token.Token {
				{ Type: token.TEXT, Literal: "Second Title" },
	 		},
		},
		{
			line.THIRD_TITLE, []token.Token {
				{ Type: token.TEXT, Literal: "Third Title" },
	 		},
		},
		{
			line.FOURTH_TITLE, []token.Token {
				{ Type: token.TEXT, Literal: "Fourth Title" },
	 		},
		},
	}

	l := New(input)

	for i, tc := range testCases {
		line := l.GetLine()

		if line.Type != tc.expectedLineType {
			t.Fatalf("Test Case #%d - line type is wrong. expected %q, got %q",
				i+1, tc.expectedLineType, line.Type)
		}

		if len(line.Tokens) != len(tc.expectedTokens) {
			t.Fatalf("Test Case #%d - Tokens array are not the same size. expected %d, got %d",
				i+1, len(tc.expectedTokens), len(line.Tokens))
		}

		for j, tok := range line.Tokens {
			expectedToken := tc.expectedTokens[j] 

			if tok.Type != expectedToken.Type {
				t.Fatalf("Error at token #%d - Token type doesn't match. expected %q, got %q",
					j+1, tok.Type, expectedToken.Type)
			}

			if tok.Literal != expectedToken.Literal {
				t.Fatalf("Error at token #%d - Literal doesn't match. expected %q, got %q",
					j+1, tok.Literal, expectedToken.Literal)
			}
		}
	}
}

func TestTestModifier(t *testing.T) {
	input := `# First Title **bold** **Italic* *Italic*`
	
	testCases := []struct {
		expectedLineType line.LineType
		expectedTokens []token.Token
	}{
		{
			line.FIRST_TITLE, []token.Token {
				{ Type: token.TEXT, Literal: "First Title " },
				{ Type: token.BOLD, Literal: "bold" },
				{ Type: token.TEXT, Literal: " " },
				{ Type: token.ITALIC, Literal: "*Italic" },
				{ Type: token.TEXT, Literal: " " },
				{ Type: token.ITALIC, Literal: "Italic" },
	 		},

		},
	}

	l := New(input)

	for i, tc := range testCases {
		line := l.GetLine()

		if line.Type != tc.expectedLineType {
			t.Fatalf("Test Case #%d - line type is wrong. expected %q, got %q",
				i, tc.expectedLineType, line.Type)
		}

		for j, tok := range line.Tokens {
			expectedToken := tc.expectedTokens[j] 

			if tok.Literal != expectedToken.Literal {
				t.Fatalf("Error at token #%d - Literal doesn't match. expected %q, got %q",
					j+1, expectedToken.Literal, tok.Literal)
			}

			if tok.Type != expectedToken.Type {
				t.Fatalf("Error at token #%d - Token type doesn't match. expected %q, got %q",
					j+1, expectedToken.Type, tok.Type)
			}

		}
	}
}

func TestImage(t *testing.T) {
	input := `<image url>`

	testCases := []struct {
		expectedLineType line.LineType
		expectedTokens []token.Token
	}{
		{
			line.IMAGE, []token.Token {
				{ Type: token.IMAGE, Literal: "image url" },
	 		},
		},
	}
}

/*
func TestTitles(t *testing.T) {
	input := `# First Title
## Second Title
### Third Title
#### Fourth Title`

	testCases := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.FIRST_TITLE, "#"},
		{token.TEXT, "First Title"},
		{token.EOL, ""},
		{token.SECOND_TITLE, "##"},
		{token.TEXT, "Second Title"},
		{token.EOL, ""},
		{token.THIRD_TITLE, "###"},
		{token.TEXT, "Third Title"},
		{token.EOL, ""},
		{token.FOURTH_TITLE, "####"},
		{token.TEXT, "Fourth Title"},
		{token.EOF, ""},
	}

	lexer := New(input)
	for i, tc := range testCases {
		token := lexer.GetToken()
		if token.Literal != tc.expectedLiteral {
			t.Fatalf("Test Case #%d - literal wrong. expected %q, got %q",
				i, tc.expectedLiteral, token.Literal)
		}
		if token.Type != tc.expectedType {
			t.Fatalf("Test Case #%d - token type wrong. expected %q, got %q",
				i, tc.expectedType, token.Type)
		}
	}
}

func TestTextModifiers(t *testing.T) {
	input := `**Bold**
*Italic*
__Bold__
_Italic_
**Bold** And Normal Text
*Italic* And Normal Text`

	testCases := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.BOLD, "Bold"},
		{token.EOL, ""},
		{token.ITALIC, "Italic"},
		{token.EOL, ""},
		{token.BOLD, "Bold"},
		{token.EOL, ""},
		{token.ITALIC, "Italic"},
		{token.EOL, ""},
		{token.BOLD, "Bold"},
		{token.TEXT, " And Normal Text"},
		{token.EOL, ""},
		{token.ITALIC, "Italic"},
		{token.TEXT, " And Normal Text"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, tc := range testCases {
		tok := lexer.GetToken()
		if tok.Literal != tc.expectedLiteral {
			t.Fatalf("Test Case #%d - literal wrong. expected %q, got %q",
				i+1, tc.expectedLiteral, tok.Literal)
		}

		if tok.Type != tc.expectedType {
			t.Fatalf("Test Case #%d - token type wrong. expected %q, got %q",
				i+1, tc.expectedType, tok.Type)
		}
	}
}

func TestImages(t *testing.T) {
	input := `<name of image>
<simple text`

	testCases := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IMAGE, "name of image"},
		{token.EOL, ""},
		{token.TEXT, "<simple text"},
		{token.EOF, ""},
	}

	lexer := New(input)
	for i, tc := range testCases {
		tok := lexer.GetToken()
		if tok.Literal != tc.expectedLiteral {
			t.Fatalf("Test Case #%d - literal wrong. expected %q, got %q",
				i+1, tc.expectedLiteral, tok.Literal)
		}

		if tok.Type != tc.expectedType {
			t.Fatalf("Test Case #%d - token type wrong. expected %q, got %q",
				i+1, tc.expectedType, tok.Type)
		}
	}
}*/
