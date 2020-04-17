package comparetree

import (
	"regexp"
)

// RegexMatch will perfrom the regex
type RegexMatch struct {
	target string
	reg    regexp.Regexp
}

func NewRegex(target string, reg regexp.Regexp) *RegexMatch {
	return &RegexMatch{
		target: target,
		reg:    reg,
	}
}

func (r *RegexMatch) Operate() bool {
	return r.reg.Match([]byte(r.target))
}

func (r *RegexMatch) LeftNode() *Node {
	return n.lNode
}

func (r *RegexMatch) RightNode() *Node {
	return n.rNode
}

func (r *RegexMatch) LeftValue() bool {
	return n.lValue
}

func (r *RegexMatch) RightValue() bool {
	return n.rValue
}
