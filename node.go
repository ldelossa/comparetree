package comparetree

type Node interface {
	LeftNode() *Node
	RightNode() *Node
	LeftValue() bool
	RightValue() bool
	Operate() bool
}
