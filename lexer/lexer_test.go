package lexer

import (
	"testing"

	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	} {
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		// assert the type
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong expected : %s got : %s", 
						i, tt.expectedType, tok.Type)	
		}

		// assert the literal
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong expected : %s got : %s", 
						i, tt.expectedLiteral, tok.Literal)
		}
	}
}