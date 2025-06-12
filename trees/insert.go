package trees

import (
	"spellcheck/lev"
)

func (tree *tree) Insert(node *Node) {
	if tree.Root == nil {
		tree.Root = node
		return
	}

	current := tree.Root

	for {
		distance := lev.Iterative(current.Value, node.Value)
		if distance == 0 {
			// Node already exists, do not insert
			return
		} else if current.Children[distance] != nil {
			//child node at this length already exists, continue recursively inserting into it
			current = current.Children[distance]
			continue
		} else {
			// Child node does not exist, create it
			current.Children[distance] = node
			return
		}
	}

}
