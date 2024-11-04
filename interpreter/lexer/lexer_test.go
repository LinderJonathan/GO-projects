package lexer

import (
	"testing"

	"github.com/LinderJonathan/go-projects/interpreter/token"
)

func TestGetNextToken(t *testing.T) {
	input := `let x = 10 <= 20 >= 15;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "x"},
		{token.ASS, "="},
		{token.INT, "10"},
		{token.LTEQ, "<="},
		{token.INT, "20"},
		{token.GTEQ, ">="},
		{token.INT, "15"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.GetNextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Lit != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Lit)
		}
	}
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
