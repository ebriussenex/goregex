package compiler

type Parser struct {
	tokens []token
}

func NewParser(tokens []token) *Parser {
	return &Parser{tokens}
}

func (p *Parser) Parse() Node {
	root := Group{}

	for _, t := range p.tokens {
		switch t.symbol {
		case Char:
			root.Append(CharacterLiteral{Character: t.letter})
		case AnyChar:
			root.Append(WildcardLiteral{})
		}
	}

	return &root
}
