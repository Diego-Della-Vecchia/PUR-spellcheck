package trees

import "spellcheck/lev"

func (t *tree) Lookup(query string, maxDistance int) []string {
	if t.Root == nil {
		return nil
	}

	var results []string
	var search func(n *Node)

	search = func(n *Node) {
		distance := lev.Iterative(n.Value, query)
		if distance <= maxDistance {
			results = append(results, n.Value)
		}

		// Search in range [distance - maxDistance, distance + maxDistance]
		for d, child := range n.Children {
			if d >= distance-maxDistance && d <= distance+maxDistance {
				search(child)
			}
		}
	}

	search(t.Root)
	return results
}
