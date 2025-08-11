package main

func main() {
}

type Exp interface {
	MatchExp(ExpBoolVars, ExpIntVars) ExpKind
}

type ety interface {
	bool | int
}

type exp[t ety] interface {
	Exp
}

// Продолжения для булевых выражений
type ExpBoolVars interface {
	Lit(bool)
	Or(exp[bool], exp[bool])
	Lt(exp[int], exp[int])
	If(exp[bool], exp[bool], exp[bool])
}

// Продолжения для int выражений
type ExpIntVars interface {
	Lit(int)
	Add(exp[int], exp[int])
	If(exp[bool], exp[int], exp[int])
}

type ExpKind interface {
	MatchKind(KindVars)
}

type KindVars interface {
	Bool(ExpBoolVars)
	Int(ExpIntVars)
}

type boolVal struct{ ev ExpBoolVars }
type intVal struct{ ev ExpIntVars }

func (b boolVal) MatchKind(v KindVars) { v.Bool(b.ev) }
func (n intVal) MatchKind(v KindVars)  { v.Int(n.ev) }
