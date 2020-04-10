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

package agent_test

import (
	"testing"

	"github.com/cspengl/conmake/pkg/agent"
)

const (
	testProjectName = "testproject"
	testStationName = "teststation"
	expectedStationImageName = testProjectName +"/" + testStationName + ":conmake"
	expectedStationContainerName = testProjectName + "-" + testStationName
)

var testStationConfig = &agent.StationConfig{
	ProjectName: testProjectName,
	StationName: testStationName,
}

func TestConstructStationImageName(t *testing.T) {
	got := agent.ConstructStationImageName(testStationConfig)
	if got != expectedStationImageName {
		t.Fail()
	}
}

func TestConstructStationImageNameFromRaw(t *testing.T) {
	got := agent.ConstructStationImageNameFromRaw(testStationConfig.ProjectName, testStationConfig.StationName)
	if got != expectedStationImageName {
		t.Fail()
	}
}

func TestConstructStationContainerName(t *testing.T) {
	got := agent.ConstructStationContainerName(testStationConfig)
	if got != expectedStationContainerName {
		t.Fail()
	}
}


var testShellScript = []string{
	"ls -al",
	"echo \"Hello World!\"",
}

var expectedShellScript = "ls -al && echo \"Hello World!\""

func TestGenShellScript(t *testing.T) {
	got := agent.GenShellScript(testShellScript)
	if got != expectedShellScript {
		t.Fail()
	}
}
