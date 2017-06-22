package syntax

import "code.justin.tv/edge/visage/__vendor/github.com/pkg/errors"

type Node interface {
	Eval() (Node, error)
}

/*
Unary operators:
a++: ++ a
!true: ! true

Binary operators
2 + 2: + 2 2
true and false: and true false

Ternary operators
true ? 1 : 2: ?: true 1 2
*/
type BinaryOpNode struct {
	Operator string
	Left     Node
	Right    Node
}

func (n BinaryOpNode) Eval() (Node, error) {
	leftNode, err := n.Left.Eval()
	if err != nil {
		return nil, err
	}
	rightNode, err := n.Right.Eval()
	if err != nil {
		return nil, err
	}
	switch n.Operator {
	case "+":
		return NumberNode{Number: leftNode.(NumberNode).Number + rightNode.(NumberNode).Number}, nil
	case "-":
		return NumberNode{Number: leftNode.(NumberNode).Number - rightNode.(NumberNode).Number}, nil
	default:
		return nil, errors.New("unrecognized operator")
	}
}

type NumberNode struct {
	Number int
}

func (n NumberNode) Eval() (Node, error) {
	return n, nil
}
