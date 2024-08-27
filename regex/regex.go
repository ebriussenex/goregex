package regex

import (
	"github.com/ebriussenex/goregex/compiler"
	"github.com/ebriussenex/goregex/fsm"
)

type (
	regex struct {
		fsm *fsm.State
	}

	DebugStep struct {
		RunnerDrawing         string
		CurrentCharacterIndex int
	}
)

func NewRegex(re string) *regex {
	tokens := compiler.Lex(re)
	parser := compiler.NewParser(tokens)
	ast := parser.Parse()
	state, _ := ast.Compile()

	return &regex{state}
}

func (r *regex) MatchString(input string) bool {
	runner := fsm.NewRunner(r.fsm)

	return match(runner, []rune(input), nil, 0)
}

func match(runner *fsm.Runner, input []rune, debugChan chan<- DebugStep, offset int) bool {
	runner.Reset()

	if debugChan != nil {
		debugChan <- DebugStep{
			RunnerDrawing:         runner.DrawSnapshot(),
			CurrentCharacterIndex: offset,
		}
	}

	for i, character := range input {
		runner.Next(character)
		if debugChan != nil {
			debugChan <- DebugStep{
				RunnerDrawing:         runner.DrawSnapshot(),
				CurrentCharacterIndex: offset + i + 1,
			}
		}
		status := runner.GetStatus()

		if status == fsm.Fail {
			return match(runner, input[1:], debugChan, offset+1)
		}

		if status == fsm.Success {
			return true
		}
	}
	return runner.GetStatus() == fsm.Success
}

func (r *regex) DebugFSM() string {
	graph, _ := r.fsm.Draw()
	return graph
}

func (r *regex) DebugMatch(input string) []DebugStep {
	runner := fsm.NewRunner(r.fsm)
	debugStepChan := make(chan DebugStep)
	go func() {
		match(runner, []rune(input), debugStepChan, 0)
		defer close(debugStepChan)
	}()

	var debugSteps []DebugStep
	for step := range debugStepChan {
		debugSteps = append(debugSteps, step)
	}

	return debugSteps
}
