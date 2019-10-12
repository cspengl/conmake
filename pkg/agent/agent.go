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

package agent

const(
  ConmakeTag = "conmake"
  Workspace = "/workspace/"
)

type StationConfig struct {
  ProjectName string
  StationName string
  Image       string
  Script      []string
  Workspace   string
}

type Agent interface {
  PerformStep(*StationConfig) error
  InitStation(*StationConfig, bool) (string, error)
  DeleteStation(*StationConfig) error
  StationExists(*StationConfig) (bool, error)
  Info()
}

func genShellScript(script []string) string {
  res := ""

  if len(script) == 0 {
    return ""
  }

  for _, cmd := range script{
    res = res + cmd + " && "
  }

  return res[:len(res)-4]
}

func PerformOnHost(script []string) error {
  //TODO: Has to be implemented
  return nil
}

func ConstructStationImageName(config *StationConfig) string{
  return ConstructStationImageNameFromRaw(config.ProjectName, config.StationName)
}

func ConstructStationImageNameFromRaw(projectname, stationname string) string {
  return projectname+"/"+stationname+":"+ConmakeTag
}

func ConstructStationContainerName(config *StationConfig) string{
  return config.ProjectName+"-"+config.StationName
}
