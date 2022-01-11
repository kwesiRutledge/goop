package solvers

import (
	"fmt"

	"github.com/kwesiRutledge/goop"
	"github.com/kwesiRutledge/gurobi.go/gurobi"
)

func VarTypeToGRBVType(goopTypeIn goop.VarType) (rune, error) {
	// Double check

	switch goopTypeIn {
	case goop.Continuous:
		return gurobi.CONTINUOUS, nil
	case goop.Binary:
		return gurobi.BINARY, nil
	case goop.Integer:
		return gurobi.INTEGER, nil
	default:
		return -1, fmt.Errorf("The goop variable type \"%v\" is not currently supported by VarTypeToGRBVType.")

	}
}
