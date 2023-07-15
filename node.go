package main

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Node struct {
	ID     uuid.UUID
	childs []*Node
}

type State struct {
	root     *Node
	allNodes []*Node
}

func (s *State) empty() {
	newRoot := Node{
		ID:     uuid.New(),
		childs: []*Node{},
	}

	s.root = &newRoot
	s.allNodes = append([]*Node{}, &newRoot)
}

func (s *State) indexOf(parentId uuid.UUID) int {
	return slices.IndexFunc(s.allNodes, func(node *Node) bool {
		return node.ID == parentId
	})
}

func (n *Node) findParentNodeOfNode(wantedNodeId uuid.UUID) *Node {
	for _, childNode := range n.childs {
		if childNode.ID == wantedNodeId {
			return n
		}
		foundNode := childNode.findParentNodeOfNode(wantedNodeId)
		if foundNode != nil {
			return foundNode
		}
	}
	return nil
}

func (s *State) findParent(wantedNodeId uuid.UUID) *Node {
	return s.root.findParentNodeOfNode(wantedNodeId)
}

func (s *State) add(parentId uuid.UUID) int {
	parentIndex := s.indexOf(parentId)
	if parentIndex == -1 {
		return -1
	}

	var newNode = Node{
		ID:     uuid.New(),
		childs: []*Node{},
	}

	s.allNodes = append(s.allNodes, &newNode)
	s.allNodes[parentIndex].childs = append(s.allNodes[parentIndex].childs, &newNode)
	return 0
}

func (s *State) delete(nodeId uuid.UUID) int {
	parent := s.findParent(nodeId)
	if parent == nil {
		return -1
	}

	for i, child := range parent.childs {
		if child.ID == nodeId {
			parent.childs = append(parent.childs[:i], parent.childs[i+1:]...)
			break
		}
	}

	for i, child := range s.allNodes {
		if child.ID == nodeId {
			s.allNodes = append(s.allNodes[:i], s.allNodes[i+1:]...)
			break
		}
	}
	return 0
}

func formatDisplay(index int, nodeId uuid.UUID, deepness int) string {
	marginLeft := deepness * 15
	return fmt.Sprintf(`
		<li style="margin-left: %dpx">
    	<button hx-post="/add/%s" hx-target="#list" hx-swap="innerHtml">
    		+
    	</button>
    	<button hx-post="/delete/%s" hx-target="#list" hx-swap="innerHtml">
    		-
    	</button>
    	%d
		</li>`, marginLeft, nodeId.String(), nodeId.String(), index+1)
}

func (n *Node) display(nodeIndex int, deepness int) string {
	display := formatDisplay(nodeIndex, n.ID, deepness)

	for index, node := range n.childs {
		display += node.display(index, deepness+1)
	}
	return display
}

func (s *State) display() string {
	return s.root.display(0, 0)
}
