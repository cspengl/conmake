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
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/cspengl/conmake/api/v1"
	"github.com/cspengl/conmake/pkg/agent"

	ocispec "github.com/opencontainers/runtime-spec/specs-go"

	"github.com/tj/go-spin"
)

const (
	//CONMAKEREF defines the conmake station image reference
	CONMAKEREF = "conmake"

	//WORKDIR defines the conmake working directory in the station container
	WORKDIR = "/workspace/"
)

// Conmaker models the conmaker based on an agent to use, a conmakefile and
// the project path.
type Conmaker struct {
	agent       agent.Agent
	conmakefile *v1.Conmakefile
	projectpath string
}

// NewConmaker returns an instance of Conmaker based on existing agent and conmakefile
func NewConmaker(a agent.Agent, c *v1.Conmakefile, p string) *Conmaker {

	return &Conmaker{
		agent:       a,
		conmakefile: c,
		projectpath: p,
	}
}

// Public functions

// PerformStep performs a step specified in a Conmakefile
func (cm *Conmaker) PerformStep(step string) error {

	//Validate step
	valid := cm.validateStep(step)

	if !valid {
		return errors.New("Step not found")
	}

	//Retrieving step definition
	stepdef := cm.conmakefile.Steps[step]

	//Defining station image id
	stationImageID := constructStationImageID(cm.conmakefile.Project, stepdef.Workstation)

	//Check if station is initialized
	stationPresent, err := cm.agent.ImagePresent(stationImageID)

	if err != nil {
		return err
	}

	//If station not present init station
	if !stationPresent {
		err := cm.InitStation(stepdef.Workstation)

		if err != nil {
			return err
		}
	}

	//Construct containerID
	containerID := constructStationContainerID(
		cm.conmakefile.Project, stepdef.Workstation)

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
				UID: 1,
				GID: 1,
			},
			Args: cm.generateArgs(stepdef.Script),
		},
	}

	//Run station to perform step
	err = cm.runStation(config, false)

	if err != nil {
		return err
	}

	//Destroy station
	return cm.agent.DestroyStationContainer(containerID)
}

// InitStation initializes station and leaves a new image stored in the
// underlying image store
func (cm *Conmaker) InitStation(station string) error {

	//Validate station
	valid := cm.validateStation(station)

	if !valid {
		return errors.New("Station not found")
	}

	//Retrieve Station definition
	stationdef := cm.conmakefile.Workstations[station]
	containerID := constructStationContainerID(
		cm.conmakefile.Project, station)
	imageID := constructStationImageID(
		cm.conmakefile.Project, station)

	//Check if there is an existing station
	oldStationPresent, err := cm.agent.ImagePresent(imageID)

	if err != nil {
		return err
	}

	//Deleting old station
	if oldStationPresent {
		err = cm.DeleteStation(station)

		if err != nil {
			return err
		}
	}

	//Construct station config
	config := agent.StationConfig{
		ContainerID: containerID,
		ImageID:     stationdef.Base,
		Mounts:      []ocispec.Mount{},
		Process: ocispec.Process{
			Terminal: true,
			Cwd:      WORKDIR,
			User: ocispec.User{
				UID: 1000,
				GID: 1000,
			},
			Args: cm.generateArgs(stationdef.Script),
		},
	}

	//Run initialization script
	err = cm.runStation(config, true)

	if err != nil {
		return err
	}

	//Save container state to new image
	err = cm.agent.SaveStationContainer(containerID, imageID)

	if err != nil {
		return err
	}

	//Destroy station container
	return cm.agent.DestroyStationContainer(containerID)
}

// DeleteStation deletes a workstation of the Conmakefile associated to the Conmaker
// with the associated agent.
func (cm *Conmaker) DeleteStation(station string) error {

	valid := cm.validateStation(station)

	if !valid {
		return errors.New("Station Not Found")
	}

	return cm.agent.DeleteImage(constructStationImageID(
		cm.conmakefile.Project, station))
}

//StationList is currently not implemented and returns nothing
func (cm *Conmaker) StationList() error {
	return nil
}

// Private functions

func (cm *Conmaker) runStation(config agent.StationConfig, quiet bool) error {

	//Preparing station
	err := cm.prepareStation(config)

	if err != nil {
		return err
	}

	//Creating station container
	err = cm.agent.CreateStationContainer(config)

	if err != nil {
		return err
	}

	//Running the station container
	output, err := cm.agent.RunStationContainer(config.ContainerID, quiet)

	if err != nil {
		return err
	}

	if (!quiet) {
		//Printing output to console
		buf := new(bytes.Buffer)
		buf.ReadFrom(output)
		fmt.Print(buf.String())
	}

	return err
}

func (cm *Conmaker) prepareStation(config agent.StationConfig) error {

	imagePresent, err := cm.agent.ImagePresent(config.ImageID)

	if err != nil {
		return err
	}

	if !imagePresent {

		progress, err := cm.agent.DownloadImage(config.ImageID)

		if err != nil {
			return err
		}

		progressBuffer := make([]byte, 32*2048)
		downloadSpinner := spin.New()
		for {
			_, downloadErr := progress.Read(progressBuffer)
			if downloadErr != nil {
				if downloadErr == io.EOF {
					break
				}
			}
			fmt.Printf("\rDownloading image %s", downloadSpinner.Next())
		}

		//clearing console
		fmt.Println("")
		defer progress.Close()
	}

	return err
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

func (cm *Conmaker) generateArgs(script []string) []string {

	if len(script) != 0 {

		//Creating new shell
		args := []string{"sh", "-c"}

		//oneLineScript
		oneLineScript := ""

		//Appending commands
		for _, command := range script {
			oneLineScript = oneLineScript + command
			oneLineScript = oneLineScript + ";"
		}

		oneLineScript = oneLineScript[:len(oneLineScript)-1]

		args = append(args, oneLineScript)

		return args
	}

	return []string{"pwd"}
}

// Static functions

func constructStationContainerID(project, station string) string {
	return project + "-" + station
}

func constructStationImageID(project, station string) string {
	return project + "/" + station + ":" + CONMAKEREF
}
