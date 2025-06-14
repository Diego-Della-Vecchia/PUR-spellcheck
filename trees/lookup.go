package trees

import "spellcheck/lev"

func (t *tree) Lookup(query string, maxDistance int) []string {
	if t.Root == nil {
		return nil
	}

	var results []string
	var search func(n *Node)

	search = func(n *Node) {
		distance, _ := lev.Iterative(n.Value, query)
		if distance <= maxDistance {
			results = append(results, n.Value)
		}

		for d, child := range n.Children {
			if d >= distance-maxDistance && d <= distance+maxDistance {
				search(child)
			}
		}
	}

	search(t.Root)
	return results
}
