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

func (s *State) AddTransition(dest *State, predicate Predicate) {
	t := Transition{
		to: dest, from: s, predicate: predicate,
	}

	s.transitions = append(s.transitions, t)
}

func (s *State) clear() {
	s.transitions = nil
}

func (s *State) Merge(destState *State) {
	// a -> b +  c -> d = a -> b -> d
	// remove leafs from c to b
	for _, t := range destState.transitions {
		s.AddTransition(t.to, t.predicate)
	}

	destState.clear()
}
