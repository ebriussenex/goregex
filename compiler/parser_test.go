package compiler

import (
	"reflect"
	"testing"
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
			tokens := lex(tt.input)
			p := NewParser(tokens)
			result := p.Parse()

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Fatalf("Expected:\n%+v\nGot:\n%+v\n", tt.expectedResult, result)
			}
		})
	}
}

