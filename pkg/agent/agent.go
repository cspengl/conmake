package agent

import(
  "github.com/cspengl/conmake/pkg/models"
)

type PerformConfig struct{
  ProjectName string
  ProjectPath string
  StepName    string
  Step        models.Step
}

type StationConfig struct {
  ProjectName string
  StationName string
  Workstation models.Workstation
}

type Agent interface {
  PerformStep(PerformConfig) error
  InitStation(*StationConfig) error
  Info()
}
