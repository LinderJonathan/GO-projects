package lexer

import (
	"testing"

	"github.com/LinderJonathan/go-projects/interpreter/token"
)

func Testget_next_token(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASS, "="},
		{token.PLUS, "+"},
	}
}
