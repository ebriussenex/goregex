package fsm

import (
	"testing"
)

func TestConstructedFSM(t *testing.T) {
	// arrange
	initialState := State{}
	stateA := State{}
	stateB := State{}
	stateC := State{}

	initialState.transitions = append(initialState.transitions, Transition{
		to: &stateA,
		predicate: func(input rune) bool {
			return input == 'a'
		},
	})

	stateA.transitions = append(stateA.transitions, Transition{
		to: &stateB,
		predicate: func(input rune) bool {
			return input == 'b'
		},
	})

	stateB.transitions = append(stateB.transitions, Transition{
		to: &stateC,
		predicate: func(input rune) bool {
			return input == 'c'
		},
	})

	type testCase struct {
		name           string
		input          string
		expectedStatus Status
	}

	testCases := []testCase{
		{"empty string", "", Normal},
		{"non matching string", "x", Fail},
		{"matching string", "abc", Success},
		{"partial matching string", "ab", Normal},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			runner := NewRunner(&initialState)

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
