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

	initialState.Transitions = append(initialState.Transitions, Transition{
		to: &stateA,
		predicate: Predicate{
			AllowedChars: "a",
		},
	})

	stateA.Transitions = append(stateA.Transitions, Transition{
		to: &stateB,
		predicate: Predicate{
			AllowedChars: "b",
		},
	})

	stateB.Transitions = append(stateB.Transitions, Transition{
		to: &stateC,
		predicate: Predicate{
			AllowedChars: "c",
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
