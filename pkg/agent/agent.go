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

//Package agent contains the generic agent definition
package agent

const (
	//ConmakeTag defines the tag for tagging conmake images
	ConmakeTag = "conmake"

	//Workspace defines the mounting target for the projectpath inside
	//the workstation container
	Workspace = "/workspace/"
)

//StationConfig models the configuration of a station to be spinned up and used
type StationConfig struct {
	//Name of the project
	ProjectName string
	//Name of the station
	StationName string
	//Used image
	Image string
	//Script to be executed on the station
	Script []string
	//Mounting source
	Workspace string
}

//Agent models a generic agent performing steps based on
//a given StationConfig
type Agent interface {
	PerformStep(*StationConfig) error
	InitStation(*StationConfig, bool) (string, error)
	DeleteStation(*StationConfig) error
	StationExists(*StationConfig) (bool, error)
	Info()
}

//GenShellScript generates a shell script from a list of commands
func GenShellScript(script []string) string {
	res := ""

	if len(script) == 0 {
		return ""
	}

	for _, cmd := range script {
		res = res + cmd + " && "
	}

	return res[:len(res)-4]
}

// func PerformOnHost(script []string) error {
// 	//TODO: Has to be implemented
// 	return nil
// }

//ConstructStationImageName constructs the name of an image based on a given station configuration
func ConstructStationImageName(config *StationConfig) string {
	return ConstructStationImageNameFromRaw(config.ProjectName, config.StationName)
}

//ConstructStationImageNameFromRaw constructs the image name based on the project name and the stationname
func ConstructStationImageNameFromRaw(projectname, stationname string) string {
	return projectname + "/" + stationname + ":" + ConmakeTag
}

//ConstructStationContainerName constructs the name of a container based on a given station config.
func ConstructStationContainerName(config *StationConfig) string {
	return config.ProjectName + "-" + config.StationName
}
