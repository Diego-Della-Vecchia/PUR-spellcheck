package trees

//This file defines the types used in the trees package.

type Node struct {
	Value    string
	Children map[int]*Node
}
type tree struct {
	Root *Node
}
