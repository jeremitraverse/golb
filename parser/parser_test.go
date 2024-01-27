package parser

import (
	"testing"

	"github.com/jeremitraverse/golb/lexer"
)

func TestTitleParse(t* testing.T) {
	input := `# _First Title_
## **Second Title**
### __Third Title__
#### *Fourth Title*
# _First Title_ Mix
`

	expectedStrings := []string { 
		"<h1><i>First Title</i></h1>",
		"<h2><b>Second Title</b></h2>",
		"<h3><b>Third Title</b></h3>",
		"<h4><i>Fourth Title</i></h4>",
		"<h1><i>First Title</i> Mix</h1>",
	}

	lex := lexer.New(input)	
	for i, expectedNode := range expectedStrings {
		li := lex.GetLine()	
		p := New(li)
		
		parsedLine := p.ParseLine()
		if parsedLine != expectedNode {
			t.Fatalf("Error at node #%d, expected %s got %s", i, expectedNode, parsedLine)
		}
	}	
}

func TestTextParsing(t* testing.T) {
	input := `_Italic Text_
*Italic Text* Regular Text
**Bold Text** *Regular Text
__Bold Text__ Regular Text*`

	expectedStrings := []string { 
		"<div><i>Italic Text</i></div>",
		"<div><i>Italic Text</i> Regular Text</div>",
		"<div><b>Bold Text</b> *Regular Text</div>",
		"<div><b>Bold Text</b> Regular Text*</div>",
	}

	lex := lexer.New(input)	
	for i, expectedNode := range expectedStrings {
		li := lex.GetLine()	
		p := New(li)
			
		parsedLine := p.ParseLine()

		if parsedLine != expectedNode {
			t.Fatalf("Error at node #%d, expected %s got %s", i, expectedNode, parsedLine)
		}
	}	
}

func TestImageParsing(t *testing.T) {
	input := `<photo.png>
< Text`
	expectedStrings := []string { 
		"<img src=\"photo.png\" />",
		"<div>< Text</div>",
	}

	lex := lexer.New(input)	
	for i, expectedNode := range expectedStrings {
		li := lex.GetLine()	
		p := New(li)
			
		parsedLine := p.ParseLine()

		if parsedLine != expectedNode {
			t.Fatalf("Error at node #%d, expected %s got %s", i, expectedNode, parsedLine)
		}
	}	

}
