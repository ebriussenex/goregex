package fsm

type Runner struct {
	head    *State
	current *State
}

func NewRunner(initialState *State) *Runner {
	r := &Runner{
		head:    initialState,
		current: initialState,
	}

	return r
}

func (r *Runner) Next(input rune) {
	if r.current == nil {
		return
	}

	r.current = r.current.firstMatchingTransition(input)
}

func (r *Runner) GetStatus() Status {
	if r.current == nil {
		return Fail
	}

	if r.current.isSuccessState() {
		return Success
	}

	return Normal
}

func (r *Runner) Reset() {
	r.current = r.head
}
