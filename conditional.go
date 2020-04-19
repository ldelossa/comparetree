package comparetree

import "log"

// And will AND the result of it's
// left and right children
type And struct {
	LNode  Node
	RNode  Node
	LValue bool
	RValue bool
}

func (n *And) LeftNode() Node {
	return n.LNode
}

func (n *And) RightNode() Node {
	return n.RNode
}

func (n *And) LeftValue() bool {
	return n.LValue
}

func (n *And) RightValue() bool {
	return n.RValue
}

// Operate will perform a POT
// of it's children then self.
func (n *And) Operate() bool {
	if n.LNode != nil {
		log.Printf("And: evaluating LNode")
		n.LValue = n.LNode.Operate()
		log.Printf("And: evaluated LNode and got %v", n.LValue)
	}
	if n.RNode != nil {
		log.Printf("And: evaluating RNode")
		n.RValue = n.RNode.Operate()
		log.Printf("And: evaluated RNode and got %v", n.RValue)
	}
	log.Printf("And: evaluating self %v && %v", n.LValue, n.RValue)
	return n.LValue && n.RValue
}

// Or will OR the result of it's
// left and right children
type Or struct {
	LNode  Node
	RNode  Node
	LValue bool
	RValue bool
}

func (n *Or) LeftNode() Node {
	return n.LNode
}

func (n *Or) RightNode() Node {
	return n.RNode
}

func (n *Or) LeftValue() bool {
	return n.LValue
}

func (n *Or) RightValue() bool {
	return n.RValue
}

func (n *Or) Operate() bool {
	if n.LNode != nil {
		log.Printf("Or: evaluating LNode")
		n.LValue = n.LNode.Operate()
		log.Printf("Or: evaluated LNode and got %v", n.LValue)
	}
	if n.RNode != nil {
		log.Printf("Or: evaluating RNode")
		n.RValue = n.RNode.Operate()
		log.Printf("Or: evaluated RNode and got %v", n.RValue)
	}
	log.Printf("And: evaluating self %v || %v", n.LValue && n.RValue)
	return n.LValue || n.RValue
}
