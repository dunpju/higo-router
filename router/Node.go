package router

type Node struct {
	isEnd    bool
	Children map[string]*Node
}

func NewNode() *Node {
	return &Node{Children: make(map[string]*Node)}
}
