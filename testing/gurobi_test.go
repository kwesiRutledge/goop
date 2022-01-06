// +build !travis

package goop_test

import (
	"testing"

	"github.com/kwesiRutledge/goop/solvers"
)

func TestGurobi(t *testing.T) {
	t.Run("SimpleMIP", func(t *testing.T) {
		solveSimpleMIPModel(t, solvers.NewGurobiSolver())
	})

	t.Run("SumRowsCols", func(t *testing.T) {
		solveSumRowsColsModel(t, solvers.NewGurobiSolver())
	})
}
