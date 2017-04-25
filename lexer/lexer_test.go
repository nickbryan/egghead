package lexer

import (
	"testing"

	"github.com/nickbryan/egghead/token"
	"github.com/stretchr/testify/assert"
)

var input string = `
let five = 5;
let ten = 10;

let add = func(x, y) {
    x + y;
};

let result = add(five, ten);
`

var tests = []struct {
	expectedType    token.Type
	expectedLiteral string
}{
	{token.LET, "let"},
	{token.IDENT, "five"},
	{token.ASSIGN, "="},
	{token.INT, "5"},
	{token.SEMICOLON, ";"},
	{token.LET, "let"},
	{token.IDENT, "ten"},
	{token.ASSIGN, "="},
	{token.INT, "10"},
	{token.SEMICOLON, ";"},
	{token.LET, "let"},
	{token.IDENT, "add"},
	{token.ASSIGN, "="},
	{token.FUNCTION, "func"},
	{token.LPAREN, "("},
	{token.IDENT, "x"},
	{token.COMMA, ","},
	{token.IDENT, "y"},
	{token.RPAREN, ")"},
	{token.LBRACE, "{"},
	{token.IDENT, "x"},
	{token.PLUS, "+"},
	{token.IDENT, "y"},
	{token.SEMICOLON, ";"},
	{token.RBRACE, "}"},
	{token.SEMICOLON, ";"},
	{token.LET, "let"},
	{token.IDENT, "result"},
	{token.ASSIGN, "="},
	{token.IDENT, "add"},
	{token.LPAREN, "("},
	{token.IDENT, "five"},
	{token.COMMA, ","},
	{token.IDENT, "ten"},
	{token.RPAREN, ")"},
	{token.SEMICOLON, ";"},
	{token.EOF, ""},
}

func TestNew(t *testing.T) {
	assert.IsType(t, new(Tokenizer), New(""))
}

func TestAnalyser_NextToken(t *testing.T) {
	lex := New(input)

	for _, test := range tests {
		tok := lex.NextToken()

		assert.Equal(t, test.expectedType, tok.Type, "Wrong token type")
		assert.Equal(t, test.expectedLiteral, tok.Literal, "Wrong token literal")
	}
}
