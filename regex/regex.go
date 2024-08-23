package regex

import (
	"github.com/ebriussenex/goregex/compiler"
	"github.com/ebriussenex/goregex/fsm"
)

type regex struct {
	fsm *fsm.State
}

func NewRegex(re string) *regex {
	tokens := compiler.Lex(re)
	parser := compiler.NewParser(tokens)
	ast := parser.Parse()
	state, _ := ast.Compile()

	return &regex{state}
}

func (r *regex) MatchString(input string) bool {
	runner := fsm.NewRunner(r.fsm)

	for _, character := range input {
		runner.Next(character)
	}

	return runner.GetStatus() == fsm.Success
}
