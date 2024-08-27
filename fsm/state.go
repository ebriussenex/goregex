package fsm

const (
	Success Status = "success"
	Fail           = "fail"
	Normal         = "normal"
)

type Status string
type State struct {
	Transitions []Transition
}

func (s *State) firstMatchingTransition(input rune) *State {
	for _, transition := range s.Transitions {
		if transition.predicate.mustCheck(input) {
			return transition.to
		}
	}

	return nil
}

func (s *State) isSuccessState() bool {
	if len(s.Transitions) == 0 {
		return true
	}

	return false
}

func (s *State) AddTransition(dest *State, predicate Predicate, debugSymbol string) {
	t := Transition{
		debugSymbol: debugSymbol,
		to:          dest, from: s, predicate: predicate,
	}

	s.Transitions = append(s.Transitions, t)
}

func (s *State) clear() {
	s.Transitions = nil
}

func (s *State) Merge(destState *State) {
	// a -> b +  c -> d = a -> b -> d
	// remove leafs from c to b
	for _, t := range destState.Transitions {
		s.AddTransition(t.to, t.predicate, t.debugSymbol)
	}

	destState.clear()
}
