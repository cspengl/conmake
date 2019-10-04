package conmaker

import(
  "github.com/cspengl/conmake/pkg/utils/yaml"
)

const(
  keyVersion  = "version"
  keyProject  = "project"
  keyAgent    = "agent"
  keySteps    = "steps"
  keyStations = "stations"
)

type Conmakefile struct {
  Version string
  Project string
}

func NewConmakefile(data []byte) Conmakefile {

  yamlFile := yaml.Load(data)

  return Conmakefile {
    Version: yamlFile.Get(keyVersion).(string),
    Project: yamlFile.Get(keyProject).(string),
  }
}
