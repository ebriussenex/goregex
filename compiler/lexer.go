package compiler

const (
	AnyChar symbol = iota
	Pipe
	LParen
	RParen
	Char
	ZeroOrMore
	ZeroOrOne
	OneOrMore
)

type (
	symbol int
	token  struct {
		symbol symbol
		letter rune
	}
)

func Lex(input string) []token {
	var tokens []token
	for _, v := range []rune(input) {
		tokens = append(tokens, lexRune(v))
	}
	return tokens
}

func lexRune(r rune) token {
	var token token
	switch r {
	case '(':
		token.symbol = LParen
	case ')':
		token.symbol = RParen
	case '|':
		token.symbol = Pipe
	case '*':
		token.symbol = ZeroOrMore
	case '+':
		token.symbol = OneOrMore
	case '?':
		token.symbol = ZeroOrOne
	case '.':
		token.symbol = AnyChar
	default:
		token.symbol = Char
		token.letter = r
	}
	return token
}
