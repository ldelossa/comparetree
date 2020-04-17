package comparetree

// And will AND the result of it's
// left and right children
type And struct {
	lNode  *Node
	rNode  *Node
	lValue bool
	rValue bool
}

func (n *And) LeftNode() *Node {
	return n.lNode
}

func (n *And) RightNode() *Node {
	return n.rNode
}

func (n *And) LeftValue() bool {
	return n.lValue
}

func (n *And) RightValue() bool {
	return n.rValue
}

func (n *And) Operate() bool {
	return n.lValue && n.rValue
}

// Or will OR the result of it's
// left and right children
type Or struct {
	lNode  *Node
	rNode  *Node
	lValue bool
	rValue bool
}

func (n *Or) LeftNode() *Node {
	return n.lNode
}

func (n *Or) RightNode() *Node {
	return n.rNode
}

func (n *Or) LeftValue() bool {
	return n.lValue
}

func (n *Or) RightValue() bool {
	return n.rValue
}

func (n *Or) Operate() bool {
	return n.lValue || n.rValue
}
