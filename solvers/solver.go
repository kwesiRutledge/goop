package solvers

import "github.com/kwesiRutledge/goop"

/*
solver.go
Description:
	Defines the new interface Solver which should define
*/

type Solver interface {
	ShowLog(tf bool)
	SetTimeLimit(timeLimit float64)
	AddVar(varIn goop.Var)
	AddVars(varSlice []goop.Var)
	AddConstr(constrIn goop.Constr)
	SetObjective(expressionIn goop.Expr)
	Optimize()
}
