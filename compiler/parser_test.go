package compiler

import (
	"reflect"
	"testing"

	"github.com/ebriussenex/goregex/fsm"
)

func TestParser(t *testing.T) {
	type test struct {
		name, input    string
		expectedResult Node
	}

	tests := []test{
		{
			name:  "simple string",
			input: "acV",
			expectedResult: &Group{
				ChildNodes: []Node{
					CharacterLiteral{Character: 'a'},
					CharacterLiteral{Character: 'c'},
					CharacterLiteral{Character: 'V'},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := Lex(tt.input)
			p := NewParser(tokens)
			result := p.Parse()

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Fatalf("Expected:\n%+v\nGot:\n%+v\n", tt.expectedResult, result)
			}
		})
	}
}

func TestCompiledFSM(t *testing.T) {
	tokens := Lex("abc")
	parser := NewParser(tokens)
	ast := parser.Parse()

	initialState, _ := ast.Compile()

	type testCase struct {
		name           string
		input          string
		expectedStatus fsm.Status
	}

	testCases := []testCase{
		{"empty string", "", fsm.Normal},
		{"non matching string", "x", fsm.Fail},
		{"matching string", "abc", fsm.Success},
		{"partial matching string", "ab", fsm.Normal},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			runner := fsm.NewRunner(initialState)

			for _, character := range tc.input {
				runner.Next(character)
			}

			actualStatus := runner.GetStatus()
			if tc.expectedStatus != actualStatus {
				t.Fatalf("expected FSM to have final state of '%v', got '%v'", tc.expectedStatus, actualStatus)
			}
		})
	}
}
