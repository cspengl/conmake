package conmakefile

import(
  "github.com/cspengl/conmake/pkg/utils"
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

  yamlFile := utils.Parse(data)

  return Conmakefile {
    Version: yamlFile.Get(keyVersion).Data().(string),
    Project: yamlFile.Get(keyProject).Data().(string),
  }
}
