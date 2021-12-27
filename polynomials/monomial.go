package polynomials

import "github.com/mit-drl/goop"

// Type Definitions

// Monomial is the product
// Example:
//	m.Coeff * power( m.Variables[0] , m.Degrees[0]) * power( m.Variables[1] , m.Degrees[1] ) * ...
type Monomial struct {
	Coeff     float64
	Degrees   []int // Degree of each variable in the monomial
	Variables []*goop.Var
}

// Functions

/*
NumVars
Description:
	Determines the number of variables in the monomial.
*/
func (m *Monomial) NumVars() int {
	return len(m.Variables)
}

/*
Vars
Description:
	Determines the ids of variables in the monomial.
*/
func (m *Monomial) Vars() []uint64 {
	idList := []uint64{}
	for _, tempVar := range m.Variables {
		idList = append(idList, tempVar.Vars()...)
	}
	return idList
}

/*
Coeffs
Description:
	returns a slice of the coefficients in the expression
*/
func (m *Monomial) Coeffs() []float64 {
	return []float64{m.Coeff}
}
