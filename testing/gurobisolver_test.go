package goop_test

import (
	"testing"

	"github.com/kwesiRutledge/goop/solvers"
)

/*
TestGurobiSolver_CreateModel1
Description:
	Tests to see if CreateModel() actually works.
*/
func TestGurobiSolver_CreateModel1(t *testing.T) {
	// Constants
	gs1 := solvers.GurobiSolver{}
	modelName1 := "Anniversary"

	// Algorithm
	gs1.CreateModel(modelName1)
	if gs1.Model == nil {
		t.Errorf("The model was not successfully created!")
	}

	defer gs1.Free()
}
