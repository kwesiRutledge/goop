package solvers

import (
	gurobi "github.com/kwesiRutledge/gurobi.go/gurobi"
)

// Type Definition

type GurobiSolver struct {
	Env          *gurobi.Env
	CurrentModel *gurobi.Model
}

// Function

/*
NewGurobiSolver
Description:
	Create a new gurobi solver object.
*/
func NewGurobiSolver() *GurobiSolver {
	// Constants
	modelName := "goopModel"

	// Algorithm

	newGS := GurobiSolver{}
	newGS.CreateModel(modelName)

	return &newGS

}

/*
CreateModel
Description:

*/
func (gs *GurobiSolver) CreateModel(modelName string) {
	// Constants

	// Algorithm
	env, err := gurobi.NewEnv(modelName + ".log")
	if err != nil {
		panic(err.Error())
	}

	gs.Env = env

	// Create an empty model.
	model, err := gurobi.NewModel(modelName, env)
	if err != nil {
		panic(err.Error())
	}
	gs.CurrentModel = model

}

/*
FreeEnv
Description:
	Frees the Env() method. Useful after the problem is solved.
*/
func (gs *GurobiSolver) FreeEnv() {
	gs.Env.Free()
}

/*
FreeModel
Description
	Frees the Model member. Useful after the problem is solved.
*/
func (gs *GurobiSolver) FreeModel() {
	gs.CurrentModel.Free()
}

/*
Free
Description:
	Frees the Env and Model elements of the system.
*/
func (gs *GurobiSolver) Free() {
	gs.FreeModel()
	gs.FreeEnv()
}
