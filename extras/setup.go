/*
setup.go
Description:
	This script is designed to help setup the gurobi.go dependency which is needed to run goop.
*/

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Type Definitions
type SetupInfo struct {
	InitialDirectory      string
	GurobiGoDirectory     string
	LatestGurobiGoVersion GurobiGoVersionInfo
}

type GurobiGoVersionInfo struct {
	MajorVersionNumber    int
	MinorVersionNumber    int
	TertiaryVersionNumber int
	TimeString            int
	Hash                  string
}

// Functions

/*
GetDefaultSetupInfo
Description:
	Defines the default setup flags for the setup script.
*/
func GetDefaultSetupInfo() (SetupInfo, error) {
	// Create Default Struct
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return SetupInfo{}, fmt.Errorf("There was an error collecting the user's home directory: %v", err)
	}

	sf0 := SetupInfo{
		InitialDirectory:  "~/go/pkg/mod/github.com/kwesi\\!rutledge/",
		GurobiGoDirectory: fmt.Sprintf("%v/go/pkg/mod/github.com/kwesi!rutledge/", homeDir),
		LatestGurobiGoVersion: GurobiGoVersionInfo{
			MajorVersionNumber:    0,
			MinorVersionNumber:    0,
			TertiaryVersionNumber: 0,
			TimeString:            20220110234629,
			Hash:                  "2556814a1f69",
		},
	}

	// Get current working directory
	tempPwd, err := os.Getwd()
	if err != nil {
		return sf0, err
	}
	sf0.InitialDirectory = tempPwd
	log.Printf("The current directory is \"%v\"", tempPwd)

	// Search through the pkg Library for all instances of gurobi.go
	libraryContents, err := os.ReadDir(sf0.GurobiGoDirectory)
	if err != nil {
		return sf0, err
	}
	gurobiGoDirectories := []string{}
	for _, content := range libraryContents {
		if content.IsDir() && strings.Contains(content.Name(), "gurobi.go") {
			fmt.Println(content.Name())
			gurobiGoDirectories = append(gurobiGoDirectories, content.Name())
		}
	}

	return sf0, nil

	// Convert Directories into Gurobi Version Info
	// gurobiVersionList, err := StringsToGurobiVersionInfoList(gurobiDirectories)
	// if err != nil {
	// 	return mlf, err
	// }

	// highestVersion, err := FindHighestVersion(gurobiVersionList)
	// if err != nil {
	// 	return mlf, err
	// }

	// // Write the highest version's directory into the GurobiHome variable
	// mlf.GurobiHome = fmt.Sprintf("/Library/gurobi%v%v%v/mac64", highestVersion.MajorVersion, highestVersion.MinorVersion, highestVersion.TertiaryVersion)

	// return mlf, nil

}

/*
ToVersionName
Description:
	Converts the GurobiGoVersionInfo object into the single string which is used to contain the gurobi.go installation.
*/
func (infoIn GurobiGoVersionInfo) ToVersionName() string {
	return fmt.Sprintf(
		"v%v.%v.%v-%v-%v",
		infoIn.MajorVersionNumber,
		infoIn.MinorVersionNumber,
		infoIn.TertiaryVersionNumber,
		infoIn.TimeString,
		infoIn.Hash,
	)
}

/*
SetupLogs
Description:
	Sets up the logging file and the connection of log to the terminal.
*/

/*
SetUpLog
Description:
	Creates a log file in the new directory "logs"
	1. Checks to see if log file already exists. If it exists, then the function deletes the old log file.
	2. Creates a new log file.
	3. Sets the log module to point to new log file.
	4. Adds initial message to log file.
*/
func SetUpLog() error {
	// Constants
	logFileName := "extras/setup_log.txt"

	// Check to see if logFile exists
	_, err := os.Stat(logFileName)
	if os.IsNotExist(err) {
		//Do Nothing. The later lines will create the file.
	} else {
		//Delete the old file.
		err = os.Remove(logFileName)
		if err != nil {
			return err
		}
	}

	// Create Logging file
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		// log.Fatal(err)
		return err
	}
	log.SetOutput(io.MultiWriter(file, os.Stdout))

	// Create initial file
	log.Println("Log file created.")

	return nil

}

func main() {
	_, err := GetDefaultSetupInfo()
	if err != nil {
		log.Printf("There was an error collecting default setup info: %v", err)
	}

	// Next, parse the arguments to make_lib and assign values to the mlf appropriately
	//sf, err = ParseMakeLibArguments(sf)
}
