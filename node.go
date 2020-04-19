package comparetree

// Node is an interface that abstracts
// a Condition and an Operation
type Node interface {
	LeftNode() Node
	RightNode() Node
	LeftValue() bool
	RightValue() bool
	Operate() bool
}
