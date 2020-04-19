package comparetree

import (
	"log"
	"regexp"
)

// RegexMatch is an operator
type RegexMatch struct {
	LNode  Node
	RNode  Node
	LValue string
	RValue string
	target string
	reg    regexp.Regexp
}

// NewRegex returns a RegexMatch Node ready for have it's
// Operate method called
func NewRegex(target string, reg regexp.Regexp) *RegexMatch {
	return &RegexMatch{
		target: target,
		reg:    reg,
	}
}

func (n *RegexMatch) Operate() bool {
	log.Printf("RegexMatch: evaluating %v ~ %v", n.LValue, n.RValue)
	return n.reg.Match([]byte(n.target))
}

func (n *RegexMatch) LeftNode() Node {
	return nil
}

func (n *RegexMatch) RightNode() Node {
	return nil
}

func (n *RegexMatch) LeftValue() bool {
	return false
}

func (n *RegexMatch) RightValue() bool {
	return false
}

type EqualInt struct {
	LNode  Node
	RNode  Node
	LValue int
	RValue int
}

func NewEqualInt(a int, b int) *EqualInt {
	return &EqualInt{
		LValue: a,
		RValue: b,
	}
}

func (n *EqualInt) Operate() bool {
	log.Printf("EqualInt: now evaluating %v == %v", n.LValue, n.RValue)
	return n.LValue == n.RValue
}

func (n *EqualInt) LeftNode() Node {
	return nil
}

func (n *EqualInt) RightNode() Node {
	return nil
}

func (n *EqualInt) LeftValue() bool {
	return false
}

func (n *EqualInt) RightValue() bool {
	return false
}
