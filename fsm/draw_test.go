package fsm

import (
	"reflect"
	"testing"
)

func TestDrawFSM(t *testing.T) {
	type test struct {
		name, expected string
		fsmBuilder     func() *State
	}

	tests := []test{
		{
			name:       "simple example",
			fsmBuilder: abcBuilder,
			expected: `
			graph LR
			0((0)) --"a"--> 1((1))
			1((1)) --"b"--> 2((2))
			2((2)) --"c"--> 3((3))
			style 3 stroke:green,stroke-width:4px;`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			drawing, _ := tt.fsmBuilder().Draw()

			if drawing != tt.expected {
				t.Fatalf("Expected drawing to be \n\"%s\", got\n\"%s\"", tt.expected, &drawing)
			}
		})
	}
}

func aaaBuilder() *State {
	state1, state2, state3, state4 := &State{}, &State{}, &State{}, &State{}
	state1.AddTransition(state2, Predicate{AllowedChars: "a"}, "a")
	state2.AddTransition(state3, Predicate{AllowedChars: "a"}, "a")
	state3.AddTransition(state4, Predicate{AllowedChars: "a"}, "a")

	return state1
}

func abcBuilder() *State {
	state1, state2, state3, state4 := &State{}, &State{}, &State{}, &State{}
	state1.AddTransition(state2, Predicate{AllowedChars: "a"}, "a")
	state2.AddTransition(state3, Predicate{AllowedChars: "b"}, "b")
	state3.AddTransition(state4, Predicate{AllowedChars: "c"}, "c")
	return state1
}

func TestDrawSnapshot(t *testing.T) {
	type test struct {
		name, input, expected string
		fsmBuilder            func() *State
	}

	tests := []test{
		{
			name:  "initial snapshot",
			input: "",
			expected: `graph LR  
0((0)) --"a"--> 1((1))  
1((1)) --"b"--> 2((2))  
2((2)) --"c"--> 3((3))  
style 3 stroke:green,stroke-width:4px;
style 0 fill:#ff5555;`,
			fsmBuilder: abcBuilder,
		},
		{
			name:       "last state highlighted",
			fsmBuilder: aaaBuilder,
			input:      "aaa",
			expected: `graph LR  
0((0)) --"a"--> 1((1))  
1((1)) --"a"--> 2((2))  
2((2)) --"a"--> 3((3))  
style 3 stroke:green,stroke-width:4px;
style 3 fill:#00ab41;`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := NewRunner(tt.fsmBuilder())
			for _, char := range tt.input {
				runner.Next(char)
			}
			snapshot := runner.DrawSnapshot()

			if !reflect.DeepEqual(tt.expected, snapshot) {
				t.Fatalf("Expected drawing to be \n\"%v\"\ngot\n\"%v\"", tt.expected, snapshot)
			}
		})
	}
}
