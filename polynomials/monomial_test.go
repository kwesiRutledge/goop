package polynomials

import (
	"testing"

	"github.com/mit-drl/goop"
)

/*
TestMonomial_NumVars1
Description:
	Tests whether or not a monomial of two variables can properly detect the two variables in
	a simple monomial.
*/
func TestMonomial_NumVars1(t *testing.T) {
	// Constants
	v := goop.Var{
		ID:    1,
		Lower: -10,
		Upper: 10,
		Vtype: goop.Continuous}
	w := goop.Var{2, -10, 10, goop.Continuous}

	// Algorithm

	m1 := Monomial{2.5, []int{1, 2}, []*goop.Var{&v, &w}}

	if m1.NumVars() != 2 {
		t.Errorf("There are only two variables in the monomial but the code detects %v.", m1.NumVars())
	}

}

/*
TestMonomial_NumVars2
Description:
	Tests whether or not a monomial of two variables can properly detect the four monomials
	in the simple monomial.
*/
func TestMonomial_NumVars2(t *testing.T) {
	// Constants
	v := goop.Var{
		ID:    1,
		Lower: -10,
		Upper: 10,
		Vtype: goop.Continuous}
	w := goop.Var{2, -10, 10, goop.Continuous}
	xSlice := goop.NewModel().AddVarVector(3, -10, 10, goop.Continuous)

	// Algorithm

	m2 := Monomial{2.5, []int{1, 2, 3, 4, 5}, append([]*goop.Var{&v, &w}, xSlice...)}

	if m2.NumVars() != 5 {
		t.Errorf("There are only two variables in the monomial but the code detects %v.", m2.NumVars())
	}

}

/*
TestMonomial_Vars1
Description:
	Tests whether or not we can retrieve the variables that we should be able to do.
*/
func TestMonomial_Vars1(t *testing.T) {
	// Constants
	model1 := goop.NewModel()
	v := model1.AddVar(-10, 10, goop.Continuous)
	w := model1.AddVar(-10, 10, goop.Continuous)

	xSlice := model1.AddVarVector(3, -10, 10, goop.Continuous)

	// Algorithm

	m1 := Monomial{2.5, []int{1, 2, 3, 4, 5}, append([]*goop.Var{v, w}, xSlice...)}
	varIDSlice1 := m1.Vars()

	for _, tempID := range varIDSlice1 {
		if (tempID < 0) || (tempID > 4) {
			t.Errorf("Unexpected ID number when there are 5 IDs: %v", tempID)
		}
	}

}

/*
TestMonomial_Coeffs1
Description:
	Tests whether or not we can retrieve the variables that we should be able to do.
*/
func TestMonomial_Coeffs1(t *testing.T) {
	// Constants
	model1 := goop.NewModel()
	v := model1.AddVar(-10, 10, goop.Continuous)
	w := model1.AddVar(-10, 10, goop.Continuous)

	xSlice := model1.AddVarVector(3, -10, 10, goop.Continuous)

	// Algorithm

	m1 := Monomial{2.5, []int{1, 2, 3, 4, 5}, append([]*goop.Var{v, w}, xSlice...)}

	if m1.Coeffs()[0] != 2.5 {
		t.Errorf("The coefficient should be 2.5, but received %v", m1.Coeffs()[0])
	}

}
