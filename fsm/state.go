package fsm

const (
	Success Status = "success"
	Fail           = "fail"
	Normal         = "normal"
)

type Status string
type State struct {
	transitions []Transition
}

func (s *State) firstMatchingTransition(input rune) *State {
	for _, transition := range s.transitions {
		if transition.predicate(input) {
			return transition.to
		}
	}

	return nil
}

func (s *State) isSuccessState() bool {
	if len(s.transitions) == 0 {
		return true
	}

	return false
}
