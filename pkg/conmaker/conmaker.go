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

//Package conmaker is the core package of conmake implementing the functions
//called by the commands.
package conmaker

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/cspengl/conmake/pkg/agent"
	//"github.com/cspengl/conmake/pkg/agent/docker"
	"github.com/cspengl/conmake/pkg/agent/containerd"
	"github.com/cspengl/conmake/pkg/models"
)

//Conmaker models the conmaker based on an agent to use, a conmakefile and
//the project path.
type Conmaker struct {
	agent       agent.Agent
	conmakefile *models.Conmakefile
	projectpath string
}

//NewConmaker returns an instance of Conmaker based on existing agent and conmakefile
func NewConmaker(a agent.Agent, c *models.Conmakefile, p string) *Conmaker {
	return &Conmaker{
		agent:       a,
		conmakefile: c,
		projectpath: p,
	}
}

//InitConmaker initializes a Conmaker based on a projectpath and a path to a Conmakefile
//Currently using the local docker daemon as agent per default
//since its the only supported
func InitConmaker(projectpath, conmakefile string) (*Conmaker, error) {

	//Read file
	f, err := ioutil.ReadFile(conmakefile)

	if err != nil {
		log.Fatal("Conmakefile not found")
		return nil, err
	}

	//Parse file and construct models
	c, err := models.NewConmakefile(f)

	if err != nil {
		log.Fatal("Could not parse Conmakefile")
		return nil, err
	}

	//Construct agent
	//a, err := docker.NewDockerAgentFromEnv()
	a, err := containerd.NewContainerdAgent()

	if projectpath == "./" {
		projectpath, err = os.Getwd()
	}

	return &Conmaker{
		agent:       a,
		conmakefile: c,
		projectpath: projectpath,
	}, err
}

//Perform performs a step of the Conmakefile associated to the Conmaker with the
//associated agent.
func (c *Conmaker) Perform(step string) error {

	//Check if used step is in conmakefile
	if _, ok := c.conmakefile.Steps[step]; !ok {
		log.Fatalf("Step %s not found in conmakefile", step)
		return errors.New("Step not found!")
	}

	imageName, err := c.InitStation(c.conmakefile.Steps[step].Workstation)

	if err != nil {
		return err
	}

	config := &agent.StationConfig{
		ProjectName: c.conmakefile.Project,
		StationName: step,
		Image:       imageName,
		Script:      c.conmakefile.Steps[step].Script,
		Workspace:   c.projectpath,
	}

	return c.agent.PerformStep(config)
}

//InitStation initializes a workstation of the Conmakefile associated
//to the Conmaker with the associated agent.
func (c *Conmaker) InitStation(station string) (string, error) {

	err := c.validateStation(station)

	if err != nil {
		return "", err
	}

	config := &agent.StationConfig{
		ProjectName: c.conmakefile.Project,
		StationName: station,
		Image: agent.ConstructStationImageNameFromRaw(
			c.conmakefile.Project,
			station,
		),
		Script:    c.conmakefile.Workstations[station].Script,
		Workspace: c.projectpath,
	}

	stationExists, err := c.agent.StationExists(config)

	if err != nil {
		return "", err
	}

	if !stationExists {
		log.Println("Station not found, initializing from base...")
		config.Image = c.conmakefile.Workstations[station].Base
	}

	return c.agent.InitStation(config, stationExists)
}

//DeleteStation deletes a workstation of the Conmakefile associated to the Conmaker
//with the associated agent.
func (c *Conmaker) DeleteStation(station string) error {

	err := c.validateStation(station)

	if err != nil {
		return err
	}

	config := &agent.StationConfig{
		ProjectName: c.conmakefile.Project,
		StationName: station,
		Image: agent.ConstructStationImageNameFromRaw(
			c.conmakefile.Project,
			station,
		),
		Script:    c.conmakefile.Workstations[station].Script,
		Workspace: c.projectpath,
	}

	return c.agent.DeleteStation(config)
}

//CleanStation cleans a workstation of the Conmakefile associated to the Conmaker
//with the associated agent. This basically deletes the existing one and
//initializes a new one from the given base image.
func (c *Conmaker) CleanStation(station string) error {

	err := c.validateStation(station)

	if err != nil {
		return err
	}

	config := &agent.StationConfig{
		ProjectName: c.conmakefile.Project,
		StationName: station,
		Image: agent.ConstructStationImageNameFromRaw(
			c.conmakefile.Project,
			station,
		),
		Script:    c.conmakefile.Workstations[station].Script,
		Workspace: c.projectpath,
	}

	err = c.agent.DeleteStation(config)

	if err != nil {
		return err
	}

	config.Image = c.conmakefile.Workstations[station].Base

	_, err = c.InitStation(station)

	return err
}

func (c *Conmaker) validateStation(station string) error {
	if _, ok := c.conmakefile.Workstations[station]; !ok {
		log.Fatalf("Station %s not found in conmakefile", station)
		return errors.New("Station not found!")
	}
	return nil
}
