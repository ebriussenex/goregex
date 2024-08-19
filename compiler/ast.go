package compiler

import "github.com/ebriussenex/goregex/fsm"

type (
	Node interface {
		compile() (head *fsm.State, tail *fsm.State)
	}

	Group struct {
		ChildNodes []Node
	}

	CharacterLiteral struct {
		Character rune
	}
)

func (g *Group) compile() (*fsm.State, *fsm.State) {
	panic("unimplemented")
}

func (l CharacterLiteral) compile() (*fsm.State, *fsm.State) {
	panic("unimplemented")
}
func (g *Group) Append(node Node) {
	g.ChildNodes = append(g.ChildNodes, node)
}
