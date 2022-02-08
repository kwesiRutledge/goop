package solvers

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kwesiRutledge/goop"
	gurobi "github.com/kwesiRutledge/gurobi.go/gurobi"
)

// Type Definition

type GurobiSolver struct {
	Env                    *gurobi.Env
	CurrentModel           *gurobi.Model
	ModelName              string
	GoopIDToGurobiIndexMap map[uint64]int32 // Maps each Goop ID (uint64) to the idx value used for each Gurobi variable.
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
ShowLog
Description:
	Decides whether or not to print logs to the terminal?
*/
func (gs *GurobiSolver) ShowLog(tf bool) {
	// Constants
	logFileName := gs.ModelName + ".txt"

	// Check to see if logFile exists
	_, err := os.Stat(logFileName)
	if os.IsNotExist(err) {
		//Do Nothing. The later lines will create the file.
	} else {
		//Delete the old file.
		err = os.Remove(logFileName)
		if err != nil {
			panic(err.Error())
		}
	}

	// Create Logging file
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		// log.Fatal(err)
		panic(err.Error())
	}

	// Attach logger to terminal only if tf is true
	if tf {
		log.SetOutput(io.MultiWriter(file, os.Stdout))
	} else {
		log.SetOutput(file)
	}

	// Create initial file
	log.Println("Log file created.")

}

/*
SetTimeLimit
Description:
	Sets the time limit of the current model in gurobi solver gs.
Input:
	limitInS = Value of time limit in seconds (float)
*/
func (gs *GurobiSolver) SetTimeLimit(limitInS float64) error {

	err := gs.Env.SetDBLParam("TimeLimit", limitInS)
	if err != nil {
		return fmt.Errorf("There was an issue using SetDBLParam(): %v", err)
	}

	// If there was no error, return nil
	return nil
}

/*
GetTimeLimit
Description:
	Gets the time limit of the current model in gurobi solver gs.
Input:
	None
Output
	limitInS = Value of time limit in seconds (float)
*/
func (gs *GurobiSolver) GetTimeLimit() (float64, error) {

	limitOut, err := gs.Env.GetDBLParam("TimeLimit")
	if err != nil {
		return -1, fmt.Errorf("There was an error getting the double param TimeLimit: %v", err)
	}

	// If all things succeeded, return good data.
	return limitOut, err
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

	// Create an empty map
	gs.GoopIDToGurobiIndexMap = make(map[uint64]int32)

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

/*
AddVar
Description:
	Adds a variable to the Gurobi Model.
*/
func (gs *GurobiSolver) AddVar(varIn *goop.Var) error {
	// Constants

	// Convert Variable Type
	vType, err := VarTypeToGRBVType(varIn.Vtype)
	if err != nil {
		return fmt.Errorf("There was an error defining gurobi type: %v", err)
	}

	// Add Variable to Current Model
	_, err = gs.CurrentModel.AddVar(int8(vType), 0.0, varIn.Lower, varIn.Upper, fmt.Sprintf("x%v", varIn.ID), []*gurobi.Constr{}, []float64{})

	// Update Map from GoopID to Gurobi Idx
	gs.GoopIDToGurobiIndexMap[varIn.ID] = len(gs.CurrentModel.Variables)

	return err
}

/*
AddVars
Description:
	Adds a set of variables to the Gurobi Model.
*/
func (gs *GurobiSolver) AddVars(varSliceIn []*goop.Var) error {
	// Constants

	// Iterate through ALL variable address in varSliceIn
	for _, varPointer := range varSliceIn {
		err := AddVar(varPointer)
		if err != nil {
			// Terminate early.
			return fmt.Errorf("Error in AddVars(): %v", err)
		}
	}

	// If we successfully made it through all Var objects, then return no errors.
	return nil
}

/*
AddConstraint
Description:
	Adds a single constraint to the gurobi model object inside of the current GurobiSolver object.
*/
func (gs *GurobiSolver) AddConstr(constrIn *goop.Constr) error {
	// Constants

	// Identify the variables in the left hand side of this constraint
	var tempVarSlice []*gurobi.Var
	for _, tempGoopID := range constrIn.lhs.Vars() {
		tempGurobiIdx := gs.GoopIDToGurobiIndexMap[tempGoopID]

		// Locate the gurobi variable in the current model that has matching ID
		for _, tempGurobiVar := range gs.CurrentModel.Variables {
			if tempGurobiIdx == tempGurobiVar.Index {
				tempVarSlice = append(tempVarSlice, &tempGurobiVar)
			}
		}
	}

	// Call Gurobi library's AddConstr() function
	gs.CurrentModel.AddConstr(
		tempVarSlice,
		constrIn.lhs.Coeffs(),
		constrIn.sense,
		constrIn.rhs.Constant(),
		fmt.Sprintf("goop Constraint #%v", len(gs.CurrentModel.Constraints)),
	)
}
