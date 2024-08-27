package fsm

import (
	"fmt"
	"strings"

	"github.com/ebriussenex/goregex/orderedset"
)

func (s *State) Draw() (string, orderedset.OrderedSet[*State]) {
	transitionSet := orderedset.OrderedSet[Transition]{}
	nodeSet := orderedset.OrderedSet[*State]{}

	visitNodes(s, &transitionSet, &nodeSet)

	output := []string{
		"graph LR",
	}

	for _, transition := range transitionSet.List() {
		fromID := nodeSet.GetIndex(transition.from)
		toID := nodeSet.GetIndex(transition.to)
		output = append(
			output,
			fmt.Sprintf(
				"%d((%d)) --\"%s\"--> %d((%d))",
				fromID, fromID, transition.debugSymbol, toID, toID,
			),
		)
	}

	for _, state := range nodeSet.List() {
		if state.isSuccessState() {
			output = append(output, fmt.Sprintf("style %d stroke:green,stroke-width:4px", nodeSet.GetIndex(state)))
		}
	}

	return strings.Join(output, "\n"), nodeSet
}

func (r Runner) DrawSnapshot() string {
	graph, nodeSet := r.head.Draw()
	switch r.GetStatus() {

	case Normal:
		graph += fmt.Sprintf("\nstyle %d fill:#ff5555;", nodeSet.GetIndex(r.current))
	case Success:
		graph += fmt.Sprintf("\nstyle %d fill:#ff5555;", nodeSet.GetIndex(r.current))
	}

	return graph
}

func visitNodes(
	node *State,
	transitions *orderedset.OrderedSet[Transition],
	visited *orderedset.OrderedSet[*State],
) {
	if visited.Has(node) {
		return
	}

	for _, transition := range node.Transitions {
		transitions.Add(transition)
	}

	visited.Add(node)

	for _, transition := range node.Transitions {
		destinationNode := transition.to
		visitNodes(destinationNode, transitions, visited)
	}
}
