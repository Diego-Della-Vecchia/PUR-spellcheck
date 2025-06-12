package trees

//This file defines the types used in the trees package.

type node struct {
	Value    string
	Children map[int]*node
}
type tree struct {
	Root *node
}
