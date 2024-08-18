package fsm

type runner struct {
	head    *State
	current *State
}

func NewRunner(initialState *State) *runner {
	r := &runner{
		head:    initialState,
		current: initialState,
	}

	return r
}

func (r *runner) Next(input rune) {
	if r.current == nil {
		return
	}

	r.current = r.current.firstMatchingTransition(input)
}

func (r *runner) GetStatus() Status {
	if r.current == nil {
		return Fail
	}

	if r.current.isSuccessState() {
		return Success
	}

	return Normal
}
