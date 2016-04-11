package ast

import "github.com/rogeralsing/GoMath/engine"

type Node interface {
	Eval(context *engine.Context) float64
	String() string
	Mutate() Node
}

type LiteralNode struct {
	Value float64
}

type BinaryOp int

const (
    AddOp BinaryOp = iota
    SubOp
    MulOp
    DivOp
)

type BinaryNode struct {
	Left  Node
	Right Node
    Operator BinaryOp
}


type AddNode struct {
	Left  Node
	Right Node
}

type DivNode struct {
	Left  Node
	Right Node
}

type MulNode struct {
	Left  Node
	Right Node
}

type SubNode struct {
	Left  Node
	Right Node
}

//VariableNode represents a variable
type VariableNode struct {
	Variable string
}
