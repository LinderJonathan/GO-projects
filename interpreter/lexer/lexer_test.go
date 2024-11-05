package lexer

import (
	"fmt"
	"testing"

	"github.com/LinderJonathan/go-projects/interpreter/token"
)

func TestGetNextToken(t *testing.T) {
	// Original input test case
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

	// Extended test cases
	extraTests := []struct {
		input          string
		expectedTokens []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}
	}{
		{
			input: "let y = 5 10;",
			expectedTokens: []struct {
				expectedType    token.TokenType
				expectedLiteral string
			}{
				{token.LET, "let"},
				{token.IDENT, "y"},
				{token.ASS, "="},
				{token.INT, "5"},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			input: "func myFunction() { return; }",
			expectedTokens: []struct {
				expectedType    token.TokenType
				expectedLiteral string
			}{
				{token.FUNCTION, "func"},
				{token.IDENT, "myFunction"},
				{token.LPAR, "("},
				{token.RPAR, ")"},
				{token.LBR, "{"},
				{token.RETURN, "return"},
				{token.SEMICOLON, ";"},
				{token.RBR, "}"},
				{token.EOF, ""},
			},
		},
		{
			input: "for i = 0; i < 10; i {}",
			expectedTokens: []struct {
				expectedType    token.TokenType
				expectedLiteral string
			}{
				{token.FOR, "for"},
				{token.IDENT, "i"},
				{token.ASS, "="},
				{token.INT, "0"},
				{token.SEMICOLON, ";"},
				{token.IDENT, "i"},
				{token.LT, "<"},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.IDENT, "i"},
				{token.LBR, "{"},
				{token.RBR, "}"},
				{token.EOF, ""},
			},
		},
		{
			input: "let x = 10 <= 20; let y = x >= 5;",
			expectedTokens: []struct {
				expectedType    token.TokenType
				expectedLiteral string
			}{
				{token.LET, "let"},
				{token.IDENT, "x"},
				{token.ASS, "="},
				{token.INT, "10"},
				{token.LTEQ, "<="},
				{token.INT, "20"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "y"},
				{token.ASS, "="},
				{token.IDENT, "x"},
				{token.GTEQ, ">="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			input: "   \t \n let z = 1;   ",
			expectedTokens: []struct {
				expectedType    token.TokenType
				expectedLiteral string
			}{
				{token.LET, "let"},
				{token.IDENT, "z"},
				{token.ASS, "="},
				{token.INT, "1"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{

			input: "let x = 10 #;",
			expectedTokens: []struct {
				expectedType    token.TokenType
				expectedLiteral string
			}{
				{token.LET, "let"},     // Valid keyword
				{token.IDENT, "x"},     // Valid identifier
				{token.ASS, "="},       // Assignment operator
				{token.INT, "10"},      // Integer
				{token.ILLEGAL, "#"},   // Illegal character
				{token.SEMICOLON, ";"}, // Semicolon
				{token.EOF, ""},        // End of input
			},
		},
		{
			input: "",
			expectedTokens: []struct {
				expectedType    token.TokenType
				expectedLiteral string
			}{
				{token.EOF, ""},
			},
		},
	}

	// Run the original tests
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

	// Run additional tests
	for _, testCase := range extraTests {
		l := NewLexer(testCase.input)

		for _, tt := range testCase.expectedTokens {
			tok := l.get_next_token()
			fmt.Println(tok.Lit)
			if tok.Type != tt.expectedType {
				t.Fatalf("extra tests - token type wrong. expected=%q, got=%q", tt.expectedType, tok.Type)
			}
			if tok.Lit != tt.expectedLiteral {
				t.Fatalf("extra tests - literal wrong. expected=%q, got=%q", tt.expectedLiteral, tok.Lit)
			}
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
