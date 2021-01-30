/*
Copyright 2019 cspengl

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package conmaker is the main package for conmake
package conmaker

import (
	"errors"
	"fmt"
	"io"
	"strings"

	v1 "github.com/cspengl/conmake/api/v1"
	"github.com/cspengl/conmake/pkg/agent"

	ocispec "github.com/opencontainers/runtime-spec/specs-go"
)

const (
	// CONMAKEREF defines the conmake station image reference
	CONMAKEREF = "conmake"

	// WORKDIR defines the conmake working directory in the station container
	WORKDIR = "/workspace/"

	// SCRIPT_BASE defines the base command for executing 'shell' script on
	// a station
	scriptBase = "/bin/sh -c"

	//SCRIPT_DEFAULT sets the default script if there is no script provided
	scriptDefault = "pwd"
)

// Conmaker models the conmaker based on an agent to use, a conmakefile and
// the project path.
type Conmaker struct {
	agent       agent.Agent
	conmakefile *v1.Conmakefile
	projectpath string
	output      io.WriteCloser
}

// NewConmaker returns an instance of Conmaker based on existing agent and conmakefile
func NewConmaker(a agent.Agent, c *v1.Conmakefile, p string, o io.WriteCloser) *Conmaker {

	return &Conmaker{
		agent:       a,
		conmakefile: c,
		projectpath: p,
		output:      o,
	}
}

// Public functions

// PerformStep performs a step specified in a Conmakefile
func (cm *Conmaker) PerformStep(step string, args ...string) error {

	cm.printf("Performing step [%s]\n", step)

	//Validate step
	valid := cm.validateStep(step)

	if !valid {
		return errors.New("Step not found")
	}

	//Retrieving step definition
	stepdef := cm.conmakefile.Steps[step]

	//Prepare station for step
	stationImageID, err := cm.prepareStation(stepdef)

	if err != nil {
		return errors.New("Failed to prepare station image")
	}

	//Construct containerID
	containerID := constructStationContainerID(
		cm.conmakefile.Project, step)

	//Create station config
	config := agent.StationConfig{
		ContainerID: containerID,
		ImageID:     stationImageID,
		Mounts: []ocispec.Mount{{
			Destination: WORKDIR,
			Type:        "bind",
			Source:      cm.projectpath,
			Options:     []string{"rw", "rbind"},
		}},
		Process: ocispec.Process{
			Terminal: true,
			Cwd:      WORKDIR,
			User: ocispec.User{
				UID: 1000,
				GID: 1000,
			},
			Args: cm.generateArgs(
				stepdef.Command,
				stepdef.Script,
				args...,
			),
		},
	}

	//Run station to perform step
	//Creating station container
	err = cm.agent.CreateStationContainer(config)

	if err != nil {
		return err
	}

	//Running the station container
	output, err := cm.agent.RunStationContainer(config.ContainerID, false)

	if err != nil {
		return err
	}

	//Piping output to client
	io.Copy(cm.output, output)

	//Destroy station container
	if err = cm.agent.DestroyStationContainer(containerID); err != nil {
		return err
	}

	//Closing the output
	return cm.output.Close()
}

// InitStation initializes station and leaves a new image stored in the
// underlying image store
func (cm *Conmaker) InitStation(station string) error {

	cm.printf("Initializing station [%s]...\n", station)

	//Validate station
	valid := cm.validateStation(station)

	if !valid {
		return fmt.Errorf("station [%s] not found", station)
	}

	//Retrieve Station definition
	stationdef := cm.conmakefile.Workstations[station]

	err := cm.buildStation(station, stationdef)

	if err != nil {
		return fmt.Errorf("failed to initialize station [%s]", station)
	}

	defer cm.output.Close()

	cm.printf("Successfully initialized station [%s]\n", station)

	return err

}

// DeleteStation deletes a workstation of the Conmakefile associated to the Conmaker
// with the associated agent.
func (cm *Conmaker) DeleteStation(station string) error {

	valid := cm.validateStation(station)

	if !valid {
		return errors.New("Station Not Found")
	}

	//Closing the output
	defer cm.output.Close()

	return cm.agent.DeleteImage(constructStationImageID(
		cm.conmakefile.Project, station))
}

//StationList is currently not implemented and returns nothing
func (cm *Conmaker) StationList() error {
	return nil
}

// Private functions

func (cm *Conmaker) prepareStation(step v1.Step) (string, error) {
	//Defining station image id
	stationImageID := constructStationImageID(cm.conmakefile.Project, step.Workstation)

	//Check if station is initialized
	stationPresent, err := cm.agent.ImagePresent(stationImageID)

	if err != nil {
		return "", err
	}

	//If station not present init station
	if !stationPresent {

		valid := cm.validateStation(step.Workstation)

		if valid {
			err := cm.buildStation(step.Workstation, cm.conmakefile.Workstations[step.Workstation])

			if err != nil {
				return "", err
			}
		} else {

			stationImageID = step.Workstation

			stationPresent, err = cm.agent.ImagePresent(step.Workstation)

			if err != nil {
				return "", err
			}

			if !stationPresent {
				cm.printf("Station: %s not found locally, trying to download...\n", step.Workstation)
				cm.printf("Downloading image...\n")
				err = cm.agent.DownloadImage(step.Workstation)

				if err != nil {
					return "", err
				}
				cm.printf("Download completed\n")
			}
		}
	}

	return stationImageID, nil
}

func (cm *Conmaker) buildStation(station string, stationdef v1.Workstation) error {

	imageID := constructStationImageID(
		cm.conmakefile.Project, station)

	containerID := constructStationContainerID(
		cm.conmakefile.Project, station)

	//Check if there is an existing station
	oldStationPresent, err := cm.agent.ImagePresent(imageID)

	if err != nil {
		return err
	}

	//Deleting old station
	if oldStationPresent {
		err = cm.agent.DeleteImage(imageID)

		if err != nil {
			return err
		}
	}

	//Construct station config
	buildConfig := agent.StationConfig{
		ContainerID: containerID,
		ImageID:     stationdef.Base,
		Mounts: []ocispec.Mount{{
			Destination: WORKDIR,
			Type:        "bind",
			Source:      cm.projectpath,
			Options:     []string{"rw", "rbind"},
		}},
		Process: ocispec.Process{
			Terminal: true,
			Cwd:      WORKDIR,
			User: ocispec.User{
				UID: 1000,
				GID: 1000,
			},
			Args: cm.generateArgs("", stationdef.Script),
		},
	}

	//Build station
	return cm.agent.BuildStation(imageID, buildConfig)
}

func (cm *Conmaker) validateStep(step string) bool {
	if _, ok := cm.conmakefile.Steps[step]; !ok {
		return false
	}
	return true
}

func (cm *Conmaker) validateStation(station string) bool {
	if _, ok := cm.conmakefile.Workstations[station]; !ok {
		return false
	}
	return true
}

func (cm *Conmaker) generateArgs(command string, script []string, additionalArgs ...string) []string {

	var args []string
	if command != "" {
		args = strings.Fields(command)
		args = append(args, additionalArgs...)
	} else if len(script) != 0 {

		//Creating new shell
		args = strings.Fields(scriptBase)

		//oneLineScript
		oneLineScript := ""

		//Appending commands
		for _, command := range script {
			oneLineScript = oneLineScript + command
			oneLineScript = oneLineScript + ";"
		}

		oneLineScript = oneLineScript[:len(oneLineScript)-1]

		args = append(args, oneLineScript)
	}

	return args
}

//Printing functions
func (cm *Conmaker) print(a ...interface{}) (int, error) {
	return fmt.Fprint(cm.output, a...)
}

func (cm *Conmaker) printf(format string, a ...interface{}) (int, error) {
	format = fmt.Sprintf("\033[36m%s>\033[0m %s", CONMAKEREF, format)
	return fmt.Fprintf(cm.output, format, a...)
}

// Static functions

func constructStationContainerID(project, station string) string {
	return fmt.Sprintf("%s-%s", project, station)
}

func constructStationImageID(project, station string) string {
	return fmt.Sprintf("%s/%s:%s", project, station, CONMAKEREF)
}
