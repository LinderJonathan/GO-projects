package lexer

import (
	"testing"

	"github.com/LinderJonathan/go-projects/interpreter/token"
)

func TestGetNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = func(x, y) {
		x + y;
	};
	let result = add(five, ten);
	if (5 < 10) {
		return true;
	} else {
		return false;
	}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASS, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASS, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASS, "="},
		{token.FUNCTION, "func"},
		{token.LPAR, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAR, ")"},
		{token.LBR, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBR, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASS, "="},
		{token.IDENT, "add"},
		{token.LPAR, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAR, ")"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAR, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAR, ")"},
		{token.LBR, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBR, "}"},
		{token.ELSE, "else"},
		{token.LBR, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBR, "}"},
		{token.EOF, ""},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.get_next_token()
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
		l := NewLexer(tt.input)
		sequence := l.read_sequence()
		if sequence != tt.expected {
			t.Fatalf("test[%d] - read_sequence() wrong. Expected %q, got %q", i, tt.expected, sequence)
		}
	}

	t.Log("read_sequence passed all tests")
}
