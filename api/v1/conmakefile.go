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

// Package v1 is the package containing version 1 of the conmake API
package v1

//Conmakefile models the YAML Object of a Conmakefile
type Conmakefile struct {
	//Version is the API version of the used conmake file
	Version      string                 `yaml:"version"`
	//Project is the name of the project
	Project      string                 `yaml:"project"`
	//Steps is a list of 'targets'
	Steps        map[string]Step        `yaml:"steps,omitempty"`
	//Workstations is a list of 'stations' to run steps on
	Workstations map[string]Workstation `yaml:"workstations,omitempty"`
}

//Step models the YAML Object of a step inside a Conmakefile
type Step struct {
	Workstation string   `yaml:"workstation"`
	Script      []string `yaml:"script,omitempty"`
}

//Workstation models the YAML Object of a workstation inside a Conmakefile
type Workstation struct {
	//Base is the base image identifier for the workstation
	Base     string   `yaml:"base"`
	//Script is the initialization script for the workstation (optional)
	Script   []string `yaml:"preparation,omitempty"`
}
