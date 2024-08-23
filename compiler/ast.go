package compiler

import "github.com/ebriussenex/goregex/fsm"

type (
	Node interface {
		compile() (head *fsm.State, tail *fsm.State)
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
)

func (g *Group) compile() (*fsm.State, *fsm.State) {
	initialState := fsm.State{}
	curTail := &initialState

	for _, expression := range g.ChildNodes {
		nextStateHead, nextStateTail := expression.compile()
		curTail.Merge(nextStateHead)
		curTail = nextStateTail
	}

	return &initialState, curTail
}

func (l CharacterLiteral) compile() (*fsm.State, *fsm.State) {
	initialState := fsm.State{}
	endState := &fsm.State{}

	initialState.AddTransition(endState, func(input rune) bool {
		return input == l.Character
	})

	return &initialState, endState

}
func (g *Group) Append(node Node) {
	g.ChildNodes = append(g.ChildNodes, node)
}
