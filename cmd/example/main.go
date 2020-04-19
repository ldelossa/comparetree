package main

import (
	"log"

	ct "github.com/ldelossa/comparetree"
)

func main() {
	// lets see if ( 1=1 AND 2=2 ) AND ( 1=1 AND 2=2 ) reports true
	leftLeftEq := ct.NewEqualInt(1, 1)
	leftRightEq := ct.NewEqualInt(2, 2)
	rightLeftEq := ct.NewEqualInt(1, 1)
	rightRightEq := ct.NewEqualInt(2, 2)

	rootAnd := &ct.And{
		LNode: &ct.And{
			LNode: leftLeftEq,
			RNode: leftRightEq,
		},
		RNode: &ct.And{
			LNode: rightLeftEq,
			RNode: rightRightEq,
		},
	}
	out := rootAnd.Operate()
	log.Printf("%v", out)

	log.Printf("============")

	// lets see if ( 1=1 AND 2=2 ) AND ( 1=1 AND 2=3 ) reports false
	leftLeftEq = ct.NewEqualInt(1, 1)
	leftRightEq = ct.NewEqualInt(2, 2)
	rightLeftEq = ct.NewEqualInt(1, 1)
	rightRightEq = ct.NewEqualInt(2, 3)

	rootAnd = &ct.And{
		LNode: &ct.And{
			LNode: leftLeftEq,
			RNode: leftRightEq,
		},
		RNode: &ct.And{
			LNode: rightLeftEq,
			RNode: rightRightEq,
		},
	}
	out = rootAnd.Operate()
	log.Printf("%v", out)

}
