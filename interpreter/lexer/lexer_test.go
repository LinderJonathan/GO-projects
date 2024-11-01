package lexer

import (
	"testing"

	"github.com/LinderJonathan/go-projects/interpreter/token"
)

func TestGetNextToken(t *testing.T) {
	input := `=+(){},; func let 12345 identifier`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// Simple single-character tokens
		{token.ASS, "="},
		{token.PLUS, "+"},
		{token.LPAR, "("},
		{token.RPAR, ")"},
		{token.LBR, "{"},
		{token.RBR, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},

		// Multi-character operators (assuming you've implemented these)
		//{token.LTEQ, "<="},
		//{token.GTEQ, ">="},

		// Keywords
		{token.FUNCTION, "func"},
		{token.LET, "let"},

		// Integer literal
		{token.INT, "12345"},

		// Identifier
		{token.IDENT, "identifier"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.get_next_token()
		t.Logf("Token %d: %+v", i, tok)
		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - token type wrong. Expected =%q, got %q", i, tt.expectedType, tok.Type)
		}
		if tok.Lit != tt.expectedLiteral {
			t.Fatalf("test[%d] - token literal wrong. Expected =%q, got %q", i, tt.expectedLiteral, tok.Lit)
		}
	}
	t.Log("get_next_token passed all tests")
}

func TestReadSequence(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"let x = 5", "let"},              // keyword
		{"fn myFunction() {", "fn"},       // keyword (function)
		{"myVariable + 10", "myVariable"}, // identifier with letters
		{"variable123", "variable123"},    // identifier with letters and numbers
		{"foo_bar = 3", "foo_bar"},        // identifier with underscore
		{"if (condition) {}", "if"},       // keyword
		{"return value", "return"},        // keyword
		{"true false", "true"},            // boolean keyword
		{"while_loop ", "while_loop"},     // identifier with underscore
	}

	for i, tt := range tests {
		l := New(tt.input)
		sequence := l.read_sequence()
		if sequence != tt.expected {
			t.Fatalf("test[%d] - read_sequence() wrong. Expected %q, got %q", i, tt.expected, sequence)
		}
	}

	t.Log("read_sequence passed all tests")
}
