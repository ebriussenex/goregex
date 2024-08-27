package compiler

import "github.com/ebriussenex/goregex/fsm"

type (
	Node interface {
		Compile() (head *fsm.State, tail *fsm.State)
	}

	CompositeNode interface {
		Node
		Append(node Node)
	}

	Group struct {
		ChildNodes []Node
	}

	CharacterLiteral struct {
		Character rune
	}

	WildcardLiteral struct{}
)

func (g *Group) Compile() (*fsm.State, *fsm.State) {
	initialState := fsm.State{}
	curTail := &initialState

	for _, expression := range g.ChildNodes {
		nextStateHead, nextStateTail := expression.Compile()
		curTail.Merge(nextStateHead)
		curTail = nextStateTail
	}

	return &initialState, curTail
}

func (l CharacterLiteral) Compile() (*fsm.State, *fsm.State) {
	initialState, endState := fsm.State{}, fsm.State{}

	initialState.AddTransition(&endState, fsm.Predicate{
		AllowedChars: string(l.Character),
	},
		string(l.Character),
	)

	return &initialState, &endState
}

func (w WildcardLiteral) Compile() (*fsm.State, *fsm.State) {
	initialState, endState := fsm.State{}, fsm.State{}

	initialState.AddTransition(&endState, fsm.Predicate{
		DisallowedChars: "\n",
	}, ".")
	return &initialState, &endState
}
func (g *Group) Append(node Node) {
	g.ChildNodes = append(g.ChildNodes, node)
}
