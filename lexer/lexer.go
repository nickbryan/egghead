package lexer

import "github.com/nickbryan/egghead/token"

type Tokenizer struct {
	input        string
	position     int  // Points to the current character in the input.
	readPosition int  // Points to the current reading position in the input (after current character).
	ch           byte // Current character under examination.
}

func New(input string) *Tokenizer {
	tokenizer := &Tokenizer{input: input}
	tokenizer.readCharacter()

	return tokenizer
}

func (tokenizer *Tokenizer) NextToken() token.Definition {
	var tok token.Definition

	tokenizer.skipWhitespace()

	switch tokenizer.ch {
	case '=':
		tok = newToken(token.ASSIGN, tokenizer.ch)
	case '+':
		tok = newToken(token.PLUS, tokenizer.ch)
	case '(':
		tok = newToken(token.LPAREN, tokenizer.ch)
	case ')':
		tok = newToken(token.RPAREN, tokenizer.ch)
	case '{':
		tok = newToken(token.LBRACE, tokenizer.ch)
	case '}':
		tok = newToken(token.RBRACE, tokenizer.ch)
	case ';':
		tok = newToken(token.SEMICOLON, tokenizer.ch)
	case ',':
		tok = newToken(token.COMMA, tokenizer.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(tokenizer.ch) {
			tok.Literal = tokenizer.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		}

		if isDigit(tokenizer.ch) {
			tok.Type = token.INT
			tok.Literal = tokenizer.readNumber()
			return tok
		}

		tok = newToken(token.ILLEGAL, tokenizer.ch)
	}

	tokenizer.readCharacter()

	return tok
}

func (tokenizer *Tokenizer) readCharacter() {
	if tokenizer.readPosition >= len(tokenizer.input) {
		tokenizer.ch = 0 // ASCII code for NUL (EOF)
	} else {
		tokenizer.ch = tokenizer.input[tokenizer.readPosition]
	}

	tokenizer.position = tokenizer.readPosition
	tokenizer.readPosition++
}

func (tokenizer *Tokenizer) readIdentifier() string {
	position := tokenizer.position

	for isLetter(tokenizer.ch) {
		tokenizer.readCharacter()
	}

	return tokenizer.input[position:tokenizer.position]
}

func (tokenizer *Tokenizer) readNumber() string {
	position := tokenizer.position
	for isDigit(tokenizer.ch) {
		tokenizer.readCharacter()
	}

	return tokenizer.input[position:tokenizer.position]
}

func (tokenizer *Tokenizer) skipWhitespace() {
	for tokenizer.ch == ' ' || tokenizer.ch == '\t' || tokenizer.ch == '\n' || tokenizer.ch == '\r' {
		tokenizer.readCharacter()
	}
}

func newToken(tokenType token.Type, ch byte) token.Definition {
	return token.Definition{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
