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
#### Fourth Title
`
	
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
		{
			line.EOF, []token.Token{},
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

func TestModifiers(t *testing.T) {
	input := `# First Title Text **Bold** **Italic* *Italic* _Italic_ __Bold__
__bold__`
	
	testCases := []struct {
		expectedLineType line.LineType
		expectedTokens []token.Token
	}{
		{
			line.FIRST_TITLE, []token.Token {
				{ Type: token.TEXT, Literal: "First Title Text " },
				{ Type: token.BOLD, Literal: "Bold" },
				{ Type: token.TEXT, Literal: " " },
				{ Type: token.ITALIC, Literal: "*Italic" },
				{ Type: token.TEXT, Literal: " " },
				{ Type: token.ITALIC, Literal: "Italic" },
				{ Type: token.TEXT, Literal: " " },
				{ Type: token.ITALIC, Literal: "Italic" },
				{ Type: token.TEXT, Literal: " " },
				{ Type: token.BOLD, Literal: "Bold" },
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
	input := `<image url>
< Text`

	testCases := []struct {
		expectedLineType line.LineType
		expectedTokens []token.Token
	}{
		{
			line.IMAGE, []token.Token {
				{ Type: token.IMAGE, Literal: "image url" },
	 		},
		},
		{
			line.TEXT, []token.Token {
				{ Type: token.TEXT, Literal: "< Text" },
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

func TestCodeBlock(t *testing.T) {
	input := "```\ni := 123\n```"

	testCases := []struct {
		expectedLineType line.LineType
		expectedTokens []token.Token
	}{
		{ line.CODE, []token.Token { { Type: token.CODE, Literal: "\ni := 123\n" } }},
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
