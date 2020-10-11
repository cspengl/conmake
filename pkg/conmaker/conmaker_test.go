/*
Copyright 2020 cspengl

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

package conmaker_test

import (
	"testing"

	"bytes"
	"io"
	"io/ioutil"
	"os"

	"github.com/cspengl/conmake/pkg/agent"
	"github.com/cspengl/conmake/pkg/conmaker"
	"github.com/cspengl/conmake/pkg/utils"
)

//Implementing a fake agent for testing the conmaker

type imageStub struct{}

type agentStub struct {
	images     map[string]imageStub
	containers map[string]agent.StationConfig
}

func (a *agentStub) ImagePresent(imageID string) (bool, error) {
	//Check if there is an empty struct at imageID
	_, ok := a.images[imageID]
	return ok, nil
}

func (a *agentStub) DownloadImage(imageID string) (error) {
	//Add empty struct at imageID
	a.images[imageID] = imageStub{}

	return nil
}

func (a *agentStub) DeleteImage(imageID string) error {
	//Delete empty struct at imageID
	delete(a.images, imageID)

	return nil
}

func (a *agentStub) CreateStationContainer(config agent.StationConfig) error {
	//Add station config to "containers"
	a.containers[config.ContainerID] = config

	return nil
}

func (a *agentStub) RunStationContainer(containerID string, quiet bool) (io.ReadCloser, error) {
	output := ioutil.NopCloser(bytes.NewReader([]byte("Simulated container output")))

	return output, nil
}

func (a *agentStub) DestroyStationContainer(containerID string) error {
	//Delete container with containerID
	delete(a.containers, containerID)

	return nil
}

func (a *agentStub) BuildStation(imageID string, config agent.StationConfig) error {

	if _, ok := a.images[config.ImageID]; !ok {
		a.images[config.ImageID] = imageStub{}
	}

	//Add a new image with specified id
	a.images[imageID] = imageStub{}

	return nil
}

type discardCloser struct {
	io.Writer
}

func (discardCloser) Close() error {return nil}

// Testing the conmaker

const filePath = "/../../testdata/Conmakefile.yaml"

var aStub = &agentStub{
	images:     make(map[string]imageStub),
	containers: make(map[string]agent.StationConfig),
}

func getConmaker() *conmaker.Conmaker {

	//Reading the conmakefile
	// - getting path
	pwd, _ := os.Getwd()
	// - reading the file
	conmakefile, _ := utils.ConmakefileFromFile(pwd + filePath)

	return conmaker.NewConmaker(
		aStub,
		conmakefile,
		"",
		discardCloser{ioutil.Discard})
}

func TestPerformStep(t *testing.T) {

	cm := getConmaker()

	err := cm.PerformStep("build")
	if err != nil {
		t.Fail()
	}
}

func TestPerformInvalidStep(t *testing.T) {
	cm := getConmaker()

	err := cm.PerformStep("doesnotexist")
	if err == nil {
		t.Fail()
	}
}

func TestInitStation(t *testing.T) {
	cm := getConmaker()

	err := cm.InitStation("building")

	//Check that there is no error
	if err != nil {
		t.Fatal("Error init")
	}

	//Check that there is a prepared station image
	if _, ok := aStub.images["testapp/building:conmake"]; !ok {
		t.Fatal("Image does not exist")
	}

	//Check that the base image of the station is present
	if _, ok := aStub.images["gcc:latest"]; !ok {
		t.Fatal("Base image not present")
	}
}

func TestInitStationWithBasePresent(t *testing.T) {
	cm := getConmaker()

	//Downloading base image
	aStub.DownloadImage("gcc:latest")

	//Init station
	err := cm.InitStation("building")

	//Check that there is no error
	if err != nil {
		t.Fail()
	}

	//Check that there is a prepared station image
	if _, ok := aStub.images["testapp/building:conmake"]; !ok {
		t.Fail()
	}
}

func TestInitInvalidStation(t *testing.T) {
	cm := getConmaker()

	err := cm.InitStation("doesnotexist")

	if err == nil {
		t.Fail()
	}
}

func TestDeleteStation(t *testing.T) {
	cm := getConmaker()

	err := cm.DeleteStation("building")

	//Check that there is no error
	if err != nil {
		t.Fail()
	}

	//Check that there is no station image
	if _, ok := aStub.images["testapp/building:conmake"]; ok {
		t.Fail()
	}
}

func TestDeleteInvalidStation(t *testing.T) {
	cm := getConmaker()

	err := cm.DeleteStation("doesnotexist")

	//Check that there is an error
	if err == nil {
		t.Fail()
	}
}
