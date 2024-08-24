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

	return match(runner, []rune(input))
}

func match(runner *fsm.Runner, input []rune) bool {
	runner.Reset()
	for _, character := range input {
		runner.Next(character)
		status := runner.GetStatus()

		if status == fsm.Fail {
			return match(runner, input[1:])
		}

		if status == fsm.Success {
			return true
		}
	}
	return runner.GetStatus() == fsm.Success
}
