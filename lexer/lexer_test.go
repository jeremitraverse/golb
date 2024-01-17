package lexer

import (
	"testing"

	"github.com/jeremitraverse/golb/token"
)

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
}
